package sbargp

import (
	"flag"
	"fmt"

	"golang.org/x/exp/constraints"
)

type (
	DBConf struct {
		User       string
		PswdEnvVar string
		NetLoc     string
		Port       uint16
		Name       string
	}

	LoggingConf struct {
		Verbosity int
		SaveTo    Dir
	}
)

// Sets two flags:
//   - <longArgStart>.Verbose
//   - v
//
// <longArgStart>.Verbose and v will both increment the same underlying value by
// one every time they are supplied. Both -v and -verbose can be supplied
// multiple times.
//
// The longArgStart argument should be used to make sure the CMD line argument
// has the same name as the TOML key. If verosity is a top level key then set
// longArgStart to an empty string.
func Verbosity[T constraints.Signed](
	fs *flag.FlagSet,
	val *T,
	_default T,
	longArgStart string,
) {
	startStr := ""
	if len(longArgStart) > 0 {
		startStr = fmt.Sprintf("%s.", longArgStart)
	}

	fs.Func(
		fmt.Sprintf("%sVerbose", startStr),
		"An integer value that controls how much information to print to the console. Higher number=more information",
		Int(val, _default, 10),
	)
	fs.Func(
		"v",
		"An integer value that controls how much information to print to the console. Higher number=more information",
		Int(val, _default, 10),
	)
}

// Sets four flags:
//   - <longArgStart>.SaveTo
//   - l
//   - <longArgStart>.Verbose
//   - v
//
// <longArgDir>.SaveTo and l will both set the directory to place any log files
// in. The flag parser will check that the dir exists.
// <longArgStart>.Verbose and v will be set by [Verbosity].
//
// The longArgStart argument should be used to make sure the CMD line argument
// has the same name as the TOML key.
func Logging(
	fs *flag.FlagSet,
	lc *LoggingConf,
	_defaults LoggingConf,
	longArgStart string,
) {
	Verbosity(fs, &lc.Verbosity, _defaults.Verbosity, longArgStart)
	fs.Func(
		fmt.Sprintf("%s.SaveTo", longArgStart),
		"The dir to place logs in",
		FromTextUnmarshaler(&lc.SaveTo, _defaults.SaveTo),
	)
	fs.Func(
		"l",
		"The dir to place logs in",
		FromTextUnmarshaler(&lc.SaveTo, _defaults.SaveTo),
	)
}

// Sets 5 flags that are intended to be used to access a database:
//   - <longArgStart>.User
//   - <longArgStart>.PswdEnvVar
//   - <longArgStart>.NetLoc
//   - <longArgStart>.Port
//   - <longArgStart>.Name
//
// The longArgStart argument should be used to make sure the CMD line argument
// has the same name as the TOML key.
func DB(
	fs *flag.FlagSet,
	dc *DBConf,
	_defaults DBConf,
	longArgStart string,
) {
	fs.StringVar(
		&dc.User,
		fmt.Sprintf("%s.User", longArgStart), _defaults.User,
		"The user to access the database with",
	)
	fs.StringVar(
		&dc.PswdEnvVar,
		fmt.Sprintf("%s.PswdEnvVar", longArgStart), _defaults.PswdEnvVar,
		"The environment variable to get the database password from",
	)
	fs.StringVar(
		&dc.NetLoc,
		fmt.Sprintf("%s.NetLoc", longArgStart), _defaults.NetLoc,
		"The network path used to access the database. Can be an address or URL",
	)
	fs.Func(
		fmt.Sprintf("%s.Port", longArgStart),
		"The port used to connect to the database",
		Uint(&dc.Port, _defaults.Port, 10),
	)
	fs.StringVar(
		&dc.Name,
		"DB.Name", _defaults.Name,
		"The name of the database to connect to",
	)
}
