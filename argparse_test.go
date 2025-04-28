package sbargp

import (
	"flag"
	"strings"
	"testing"
	"time"

	sbtest "github.com/barbell-math/smoothbrain-test"
)

type (
	conf struct {
		One   int
		Two   float64
		Three string
	}

	allVals struct {
		I8  int8
		I16 int16
		I32 int32
		I64 int64

		Ui8  uint8
		Ui16 uint16
		Ui32 uint32
		Ui64 uint64

		F32 float32
		F64 float64

		OffsetDT1 time.Time
		OffsetDT2 time.Time
		OffsetDT3 time.Time
		OffsetDT4 time.Time

		LocalDT1 time.Time
		LocalDT2 time.Time

		LocalDate  time.Time
		LocalTime1 time.Time
		LocalTime2 time.Time
	}
)

func testArgDefs(c *conf, fs *flag.FlagSet) error {
	fs.IntVar(&c.One, "one", 1, "an int")
	fs.Float64Var(&c.Two, "two", 2.0, "a float")
	fs.StringVar(&c.Three, "three", "3", "a string")
	return nil
}

func testAllValueArgDefs(conf *allVals, fs *flag.FlagSet) error {
	fs.Func("I8", "int8", Int(&conf.I8, 0, 10))
	fs.Func("I16", "int16", Int(&conf.I16, 0, 10))
	fs.Func("I32", "int32", Int(&conf.I32, 0, 10))
	fs.Func("I64", "int64", Int(&conf.I64, 0, 10))

	fs.Func("Ui8", "uint8", Uint(&conf.Ui8, 0, 10))
	fs.Func("Ui16", "uint16", Uint(&conf.Ui16, 0, 10))
	fs.Func("Ui32", "uint32", Uint(&conf.Ui32, 0, 10))
	fs.Func("Ui64", "uint64", Uint(&conf.Ui64, 0, 10))

	fs.Func("F32", "float32", Float(&conf.F32, 0))
	fs.Func("F64", "float64", Float(&conf.F64, 0))

	fs.Func("OffsetDT1", "time.Time", Time(&conf.OffsetDT1, time.Now()))
	fs.Func("OffsetDT2", "time.Time", Time(&conf.OffsetDT2, time.Now()))
	fs.Func("OffsetDT3", "time.Time", Time(&conf.OffsetDT3, time.Now()))
	fs.Func("OffsetDT4", "time.Time", Time(&conf.OffsetDT4, time.Now()))

	fs.Func("LocalDT1", "time.Time", Time(&conf.LocalDT1, time.Now()))
	fs.Func("LocalDT2", "time.Time", Time(&conf.LocalDT2, time.Now()))

	fs.Func("LocalDate", "time.Time", Time(&conf.LocalDate, time.Now()))
	fs.Func("LocalTime1", "time.Time", Time(&conf.LocalTime1, time.Now()))
	fs.Func("LocalTime2", "time.Time", Time(&conf.LocalTime2, time.Now()))

	return nil
}

func TestHelp(t *testing.T) {
	var c conf
	err := Parse(&c, []string{"-h"}, ParserOpts[conf]{
		ProgName: "testing",
		ArgDefsSetter: func(c *conf, fs *flag.FlagSet) error {
			fs.IntVar(&c.One, "one", 1, "an int")
			fs.Float64Var(&c.Two, "two", 2.0, "a float")
			fs.StringVar(&c.Three, "three", "3", "a string")
			return nil
		},
	})
	sbtest.ContainsError(t, flag.ErrHelp, err)

	errStr := strings.ReplaceAll(err.Error(), "\t", "    ")
	expStr := `flag: help requested
  |- Usage of testing:
  -conf string
        Path to toml config file
  -one int
        an int (default 1)
  -three string
        a string (default "3")
  -two float
        a float (default 2)
`
	sbtest.Eq(t, errStr, expStr)
}

func TestConfDuplicated(t *testing.T) {
	var c conf
	sbtest.Panics(t, func() {
		Parse(&c, []string{"-h"}, ParserOpts[conf]{
			ProgName: "testing",
			ArgDefsSetter: func(c *conf, fs *flag.FlagSet) error {
				fs.IntVar(&c.One, "one", 1, "an int")
				fs.Float64Var(&c.Two, "two", 2.0, "a float")
				fs.StringVar(&c.Three, "conf", "3", "a string")
				return nil
			},
		})
	})
}

func TestInvalidArgument(t *testing.T) {
	var c conf
	err := Parse(&c, []string{"-asdf"}, ParserOpts[conf]{
		ProgName:      "testing",
		ArgDefsSetter: testArgDefs,
	})
	sbtest.NotNil(t, err)

	errStr := strings.ReplaceAll(err.Error(), "\t", "    ")
	expStr := `flag provided but not defined: -asdf
  |- Usage of testing:
  -conf string
        Path to toml config file
  -one int
        an int (default 1)
  -three string
        a string (default "3")
  -two float
        a float (default 2)
`
	sbtest.Eq(t, errStr, expStr)
}

func TestExtraArgsSuppliedAccordingToFlag(t *testing.T) {
	var c conf
	err := Parse(&c, []string{"asdf"}, ParserOpts[conf]{
		ProgName:      "testing",
		ArgDefsSetter: testArgDefs,
	})
	sbtest.ContainsError(t, InvalidArgumentErr, err)

	errStr := strings.ReplaceAll(err.Error(), "\t", "    ")
	expStr := `Invalid argument
  |- asdf
  |- Usage of testing:
  -conf string
        Path to toml config file
  -one int
        an int (default 1)
  -three string
        a string (default "3")
  -two float
        a float (default 2)
`
	sbtest.Eq(t, errStr, expStr)
}

func TestConfFileDoesNotExist(t *testing.T) {
	var c conf
	err := Parse(&c, []string{"-conf", "asdf"}, ParserOpts[conf]{
		ProgName:      "testing",
		ArgDefsSetter: testArgDefs,
	})
	sbtest.ContainsError(t, InvalidConfFileErr, err)

	errStr := strings.ReplaceAll(err.Error(), "\t", "    ")
	expStr := `Invalid conf file
  |- stat asdf: no such file or directory
  |- Usage of testing:
  -conf string
        Path to toml config file
  -one int
        an int (default 1)
  -three string
        a string (default "3")
  -two float
        a float (default 2)
`
	sbtest.Eq(t, errStr, expStr)
}

func TestConfFileDecodeError(t *testing.T) {
	var c conf
	err := Parse(
		&c,
		[]string{"-conf", "./bs/invalid.toml"},
		ParserOpts[conf]{ProgName: "testing", ArgDefsSetter: testArgDefs},
	)
	sbtest.ContainsError(t, InvalidConfFileErr, err)

	errStr := strings.ReplaceAll(err.Error(), "\t", "    ")
	expStr := `Invalid conf file
  |- toml: line 1: expected '.' or '=', but got '\n' instead
  |- Usage of testing:
  -conf string
        Path to toml config file
  -one int
        an int (default 1)
  -three string
        a string (default "3")
  -two float
        a float (default 2)
`
	sbtest.Eq(t, errStr, expStr)
}

func TestExtraArgsSuppliedAccordingToTomlKeys(t *testing.T) {
	var c conf
	err := Parse(
		&c,
		[]string{"-conf", "./bs/extraFlags.toml"},
		ParserOpts[conf]{ProgName: "testing", ArgDefsSetter: testArgDefs},
	)
	sbtest.ContainsError(t, InvalidConfFileErr, err)

	errStr := strings.ReplaceAll(err.Error(), "\t", "    ")
	expStr := `Invalid conf file
Invalid values in conf file
  |- Invalid Keys: [asdf]
  |- Usage of testing:
  -conf string
        Path to toml config file
  -one int
        an int (default 1)
  -three string
        a string (default "3")
  -two float
        a float (default 2)
`
	sbtest.Eq(t, errStr, expStr)
}

func TestMissingRequiredArgCMDLineOnly(t *testing.T) {
	var c conf
	err := Parse(
		&c,
		[]string{"-one", "1", "-two", "2"},
		ParserOpts[conf]{
			ProgName:      "testing",
			ArgDefsSetter: testArgDefs,
			RequiredArgs:  []string{"one", "two", "three"},
		},
	)
	sbtest.ContainsError(t, MissingRequiredArgsErr, err)

	errStr := strings.ReplaceAll(err.Error(), "\t", "    ")
	expStr := `Missing required args
  |- Still need: [three]
  |- Usage of testing:
  -conf string
        Path to toml config file
  -one int
        an int (default 1)
  -three string
        a string (default "3")
  -two float
        a float (default 2)
`
	sbtest.Eq(t, errStr, expStr)
}

func TestMissingRequiredArgTOMLOnly(t *testing.T) {
	var c conf
	err := Parse(
		&c,
		[]string{"-conf", "./bs/oneTwo.toml"},
		ParserOpts[conf]{
			ProgName:      "testing",
			ArgDefsSetter: testArgDefs,
			RequiredArgs:  []string{"one", "two", "three"},
		},
	)
	sbtest.ContainsError(t, MissingRequiredArgsErr, err)

	errStr := strings.ReplaceAll(err.Error(), "\t", "    ")
	expStr := `Missing required args
  |- Still need: [three]
  |- Usage of testing:
  -conf string
        Path to toml config file
  -one int
        an int (default 1)
  -three string
        a string (default "3")
  -two float
        a float (default 2)
`
	sbtest.Eq(t, errStr, expStr)
}

func TestMissingRequiredArgMixedTOMLAndCMDLine(t *testing.T) {
	var c conf
	err := Parse(
		&c,
		[]string{"-conf", "./bs/one.toml", "-two", "2"},
		ParserOpts[conf]{
			ProgName:      "testing",
			ArgDefsSetter: testArgDefs,
			RequiredArgs:  []string{"one", "two", "three"},
		},
	)
	sbtest.ContainsError(t, MissingRequiredArgsErr, err)

	errStr := strings.ReplaceAll(err.Error(), "\t", "    ")
	expStr := `Missing required args
  |- Still need: [three]
  |- Usage of testing:
  -conf string
        Path to toml config file
  -one int
        an int (default 1)
  -three string
        a string (default "3")
  -two float
        a float (default 2)
`
	sbtest.Eq(t, errStr, expStr)
}

func TestTOMLArgPrecedence(t *testing.T) {
	var c conf
	err := Parse(
		&c,
		[]string{"-conf", "./bs/oneTwo.toml", "-two", "3", "-three", "asdf"},
		ParserOpts[conf]{
			ProgName:      "testing",
			ArgDefsSetter: testArgDefs,
		},
	)
	sbtest.Nil(t, err)
	sbtest.Eq(t, c.One, 1)
	sbtest.Eq(t, c.Two, 2)
	sbtest.Eq(t, c.Three, "asdf")
}

// This test attempts to show consistency between the CMD line arguments and the
// values in a TOML file. Not every value will be 100% consistent however.
//   - Ints in TOML can be specified with underscores: 1_000 is valid. That is
//     not valid for the default flag int parser.
//   - time.Duration values in TOML will be interpreted as ns when no unit is
//     provided but the default flag duration parser requires units to be
//     specified.
//
// There are some other more distinct differences though:
//   - The CMD line can have abbreviations for certain things (v for verbose).
//     TOML has no concept of this whatsoever.
//   - Arguments on the CMD line can be specified multiple times and that is ok.
//     In TOML the same key cannot be duplicated. This problem gets even more
//     complicated when you consider that each time a CMD line argument is
//     provided it can do things like update a slice, increment a counter, etc.
//
// So, all of this to say: while the parsing of values from the CMD line
// interface and TOML file are kept as similar as possible, it so not possible
// to make them fully match.
//
// They could be made _more_ similar through a custom argument parser layer that
// outputs a string representation of a toml file that represents the CMD line
// arguments. This could then be handed to the TOML parser and all behavior
// would match. *But*, TOML cannot handle certain use cases that the CMD line
// can, as stated above. Trying to get the custom argument parser to output the
// correct TOML in these scenarios would require the parser to have some sort of
// semantic awareness. This semantic awareness is complicated and only further
// complicates the CMD line vs TOML interface for a user because they would now
// have to ask themselves "how does this CMD line argument translate to TOML?"
// Instead, I believe that it is simpler to deal with the minor annoyances of
// the small discrepancies in parsing behavior and allow the user to think of
// the CMD line interface, TOML file, and resulting config information all
// separately.
func TestAllValsConsistency(t *testing.T) {
	var cliConf allVals
	var tomlConf allVals

	err := Parse(
		&cliConf,
		[]string{
			"-I8", "8", "-I16", "16", "-I32", "32", "-I64", "64",
			"-Ui8", "8", "-Ui16", "16", "-Ui32", "32", "-Ui64", "64",
			"-F32", "32", "-F64", "64",
			"-OffsetDT1", "1979-05-27T07:32:00Z",
			"-OffsetDT2", "1979-05-27T00:32:00-07:00",
			"-OffsetDT3", "1979-05-27T00:32:00.999999-07:00",
			"-OffsetDT4", "1979-05-27 07:32:00Z",
			"-LocalDT1", "1979-05-27T07:32:00",
			"-LocalDT2", "1979-05-27T00:32:00.999999",
			"-LocalDate", "1979-05-27",
			"-LocalTime1", "07:32:00",
			"-LocalTime2", "00:32:00.999999",
		},
		ParserOpts[allVals]{
			ProgName:      "allVals",
			RequiredArgs:  []string{},
			ArgDefsSetter: testAllValueArgDefs,
		},
	)
	sbtest.Nil(t, err)
	err = Parse(
		&tomlConf,
		[]string{"-conf", "./bs/allVals.toml"},
		ParserOpts[allVals]{
			ProgName:      "allVals",
			RequiredArgs:  []string{},
			ArgDefsSetter: testAllValueArgDefs,
		},
	)
	sbtest.Nil(t, err)
	sbtest.Eq(t, tomlConf, cliConf)
}
