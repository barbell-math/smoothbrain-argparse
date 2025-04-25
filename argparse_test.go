package sbargp

import (
	"flag"
	"strings"
	"testing"

	sbtest "github.com/barbell-math/smoothbrain-test"
)

type conf struct {
	One   int
	Two   float64
	Three string
}

func testArgDefs(c *conf, fs *flag.FlagSet) error {
	fs.IntVar(&c.One, "one", 1, "an int")
	fs.Float64Var(&c.Two, "two", 2.0, "a float")
	fs.StringVar(&c.Three, "three", "3", "a string")
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
