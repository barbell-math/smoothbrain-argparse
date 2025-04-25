// A very simple CMD line argument parser that allows for arguments to be
// specified from both the CLI and a TOML config file.
package sbargp

import (
	"errors"
	"flag"
	"maps"
	"os"
	"slices"
	"strings"

	"github.com/BurntSushi/toml"
	sberrs "github.com/barbell-math/smoothbrain-errs"
)

type (
	ParserOpts[T any] struct {
		// The program name that should be printed out on the help menu in the
		// following format: `Usage of <program name>`. If left empty
		// <program name> will be replaced with `this program`.
		ProgName string
		// All arguments provided in this list must be provided. Arguments can
		// be provided in either the CMD line, in a TOML conf file, or both. If
		// an argument is provided in either place it will be marked as present.
		RequiredArgs []string
		// The function that allows the user (you) to define the CMD line
		// arguments. The CMD line argument parameters can be named the same as
		// the keys in the TOML file. Naming the parameters the same as their
		// associated key in the TOML file is generally recommended, and may be
		// necessary to make it so arguments in the [RequiredArgs] list are seen
		// as present from both the CMD line and TOML conf file. Hierarchical
		// keys of the TOML file are separated by a '.'.
		//
		// This function can also be used to set default values. Assuming that
		// you don't have a FlagSet argument for a given value in conf, simply
		// set that value to the default value. If a FlagSet argument
		// corresponds to a value in conf then the default value will be
		// whatever is defined by the FlagSet method.
		ArgDefsSetter func(conf *T, fs *flag.FlagSet) error
	}
)

var (
	InvalidArgumentErr         = errors.New("Invalid argument")
	InvalidConfFileErr         = errors.New("Invalid conf file")
	MissingRequiredArgsErr     = errors.New("Missing required args")
	InvalidValuesInConfFileErr = errors.New("Invalid values in conf file")
)

// Takes a sequence of CMD line arguments and parses them. The supplied `conf`
// value will be populated with any values.
//
// A `conf` argument will be added that
// will accept a path to a TOML config file. This TOML config file will be
// treated as another source of arguments, and the `conf` value will be
// populated with the contents of that file.
//
// The arguments that are present in the TOML config file will take precedence
// over all CMD line arguments.
func Parse[T any](
	conf *T,
	cliArgs []string,
	opts ParserOpts[T],
) error {
	var err error
	var confFile string
	var confMeta toml.MetaData
	var buf strings.Builder

	if opts.ProgName == "" {
		opts.ProgName = "this program"
	}

	fs := flag.NewFlagSet(opts.ProgName, flag.ContinueOnError)
	fs.SetOutput(&buf)
	fs.StringVar(
		&confFile,
		"conf",
		"",
		"Path to toml config file",
	)
	if opts.ArgDefsSetter != nil {
		if err = opts.ArgDefsSetter(conf, fs); err != nil {
			// Don't goto errReturn because the FlagSet was not properly initialized
			return err
		}
	}

	if err = parseCLIArgs(fs, cliArgs); err != nil {
		goto errReturn
	}
	confMeta, err = parseConfFile(conf, confFile)
	if err != nil {
		goto errReturn
	}
	if err = checkRequiredArgs(fs, &confMeta, &opts); err != nil {
		goto errReturn
	}

	return nil
errReturn:
	buf.Reset()
	fs.Usage()
	return sberrs.Wrap(err, buf.String())
}

func parseCLIArgs(fs *flag.FlagSet, args []string) error {
	if err := fs.Parse(args); err != nil {
		return err
	}
	if len(fs.Args()) > 0 {
		return sberrs.Wrap(InvalidArgumentErr, "%s", fs.Args()[0])
	}

	return nil
}

func parseConfFile[T any](
	c *T,
	confFile string,
) (toml.MetaData, error) {
	var err error
	var confMeta toml.MetaData

	if confFile != "" {
		if _, err := os.Stat(confFile); err != nil {
			return confMeta, sberrs.Wrap(InvalidConfFileErr, "%w", err)
		}

		confMeta, err = toml.DecodeFile(confFile, &c)
		if err != nil {
			return confMeta, sberrs.Wrap(InvalidConfFileErr, "%w", err)
		}

		undecoded := confMeta.Undecoded()
		if len(undecoded) > 0 {
			return confMeta, sberrs.AppendError(
				InvalidConfFileErr,
				sberrs.Wrap(
					InvalidValuesInConfFileErr,
					"Invalid Keys: %v",
					undecoded,
				),
			)
		}
	}

	return confMeta, nil
}

func checkRequiredArgs[T any](
	fs *flag.FlagSet,
	confMeta *toml.MetaData,
	opts *ParserOpts[T],
) error {
	requiredArgs := make(map[string]struct{}, len(opts.RequiredArgs))
	for _, arg := range opts.RequiredArgs {
		requiredArgs[arg] = struct{}{}
	}

	fs.Visit(func(f *flag.Flag) {
		if _, ok := requiredArgs[f.Name]; ok {
			delete(requiredArgs, f.Name)
		}
	})
	for _, k := range confMeta.Keys() {
		kStr := k.String()
		if _, ok := requiredArgs[kStr]; ok {
			delete(requiredArgs, kStr)
		}
	}
	if len(requiredArgs) > 0 {
		return sberrs.Wrap(
			MissingRequiredArgsErr,
			"Still need: %v",
			slices.Collect(maps.Keys(requiredArgs)),
		)
	}

	return nil
}
