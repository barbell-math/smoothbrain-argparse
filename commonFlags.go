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
// has the same name as the TOML key.
func Verbosity[T constraints.Signed](
	fs *flag.FlagSet,
	val *T,
	_default T,
	longArgStart string,
) {
	fs.Func(
		fmt.Sprintf("%s.Verbose", longArgStart),
		"An integer value that controls how much information to print to the console. Higher number=more information",
		Int(val, _default, 1),
	)
	fs.Func(
		"v",
		"An integer value that controls how much information to print to the console. Higher number=more information",
		Int(val, _default, 1),
	)
}

// Sets three flags:
//   - <longArgStart>.Dir
//   - l
//   - <longArgStart>.Verbose
//   - v
//
// <longArgDir>.Dir and l will both set the directory to place any log files in.
// The flag parser will check that the dir exists.
// <longArgStart>.Verbose and v will be set by [Verbose].
//
// The longArgStart argument should be used to make sure the CMD line argument
// has the same name as the TOML key.
func Logging(
	fs *flag.FlagSet,
	lc *LoggingConf,
	longArgStart string,
) {
	Verbosity(fs, &lc.Verbosity, 0, longArgStart)
	fs.Func(
		fmt.Sprintf("%s.Dir", longArgStart),
		"The dir to place logs in",
		FromTextUnmarshaler(&lc.SaveTo, ""),
	)
	fs.Func(
		"l",
		"The dir to place logs in",
		FromTextUnmarshaler(&lc.SaveTo, ""),
	)
}

// Sets 5 flags that are intended to be used to access a database:
//   - <longArgStart>.User        (no default set)
//   - <longArgStart>.PswdEnvVar  (no default set)
//   - <longArgStart>.NetLoc      (sets default to locahost)
//   - <longArgStart>.Port        (sets default to 5432)
//   - <longArgStart>.Name        (no default set)
//
// The longArgStart argument should be used to make sure the CMD line argument
// has the same name as the TOML key.
func DB(
	fs *flag.FlagSet,
	dc *DBConf,
	longArgStart string,
) {
	fs.StringVar(
		&dc.User,
		fmt.Sprintf("%s.User", longArgStart), "",
		"The user to access the database with",
	)
	fs.StringVar(
		&dc.PswdEnvVar,
		fmt.Sprintf("%s.PswdEnvVar", longArgStart), "",
		"The environment variable to get the database password from",
	)
	fs.StringVar(
		&dc.NetLoc,
		fmt.Sprintf("%s.NetLoc", longArgStart), "localhost",
		"The network path used to access the database. Can be an address or URL",
	)
	fs.Func(
		fmt.Sprintf("%s.Port", longArgStart),
		"The port used to connect to the database",
		Uint(&dc.Port, 5432, 10),
	)
	fs.StringVar(
		&dc.Name,
		"DB.Name", "",
		"The name of the database to connect to",
	)
}
