package sbargp

import (
	"encoding"
	"errors"
	"strconv"
	"time"

	sberr "github.com/barbell-math/smoothbrain-errs"
	"golang.org/x/exp/constraints"
)

type (
	// The type of function that [flag.Func] accepts
	FlagSetFunc func(arg string) error
)

var (
	MissingLeadingZerosErr = errors.New(
		"Leading zeros in date-time values cannot be left off",
	)
)

// Useful for parsing a specific kind of uint from the CMD line since flag does
// not have a generic version yet. (It only provides uint)
func Uint[T constraints.Unsigned](val *T, _default T, base int) FlagSetFunc {
	*val = _default

	var size int
	switch any(val).(type) {
	case *uint8:
		size = 8
	case *uint16:
		size = 16
	case *uint32:
		size = 32
	case *uint64:
		size = 64
	}

	return func(arg string) error {
		var u64 uint64
		var err error
		if base == 2 && len(arg) > 2 && arg[0:2] == "0b" {
			u64, err = strconv.ParseUint(arg[2:], base, size)
		} else if base == 8 && len(arg) > 2 && arg[0:2] == "0o" {
			u64, err = strconv.ParseUint(arg[2:], base, size)
		} else if base == 16 && len(arg) > 2 && arg[0:2] == "0x" {
			u64, err = strconv.ParseUint(arg[2:], base, size)
		} else {
			u64, err = strconv.ParseUint(arg, base, size)
		}
		*val = T(u64)
		return err
	}
}

// Useful for parsing a specific kind of int from the CMD line since flag does
// not have a generic version yet. (It only provides int)
func Int[T constraints.Signed](val *T, _default T, base int) FlagSetFunc {
	*val = _default

	var size int
	switch any(val).(type) {
	case *uint8:
		size = 8
	case *uint16:
		size = 16
	case *uint32:
		size = 32
	case *uint64:
		size = 64
	}

	return func(arg string) error {
		var i64 int64
		var err error
		if base == 2 && len(arg) > 2 && arg[0:2] == "0b" {
			i64, err = strconv.ParseInt(arg[2:], base, size)
		} else if base == 8 && len(arg) > 2 && arg[0:2] == "0o" {
			i64, err = strconv.ParseInt(arg[2:], base, size)
		} else if base == 16 && len(arg) > 2 && arg[0:2] == "0x" {
			i64, err = strconv.ParseInt(arg[2:], base, size)
		} else {
			i64, err = strconv.ParseInt(arg, base, size)
		}
		*val = T(i64)
		return err
	}
}

// Useful for parsing a specific kind of float from the CMD line since flag does
// not have a generic version yet. (It only provides float64)
func Float[T constraints.Float](val *T, _default T) FlagSetFunc {
	*val = _default

	var size int
	switch any(val).(type) {
	case *float32:
		size = 32
	case *float64:
		size = 64
	}

	return func(arg string) error {
		f64, err := strconv.ParseFloat(arg, size)
		*val = T(f64)
		return err
	}
}

// Useful for parsing a time value from the CMD line. The format will one of the
// [allowd date-time formats in TOML].
//
// [allowed date-time formats in TOML]: https://toml.io/en/v1.0.0#local-date-time
func Time(val *time.Time) FlagSetFunc {
	return func(arg string) error {
		arg = datetimeRepl.Replace(arg)
		var err error
		var t time.Time

		for _, dt := range dtTypes {
			t, err = time.ParseInLocation(dt.fmt, arg, dt.zone)
			if err != nil {
				continue
			}

			// If we are here then parsing the time passed, check for leading
			// zeros...
			if missing := missingLeadingZero(arg, dt.fmt); missing {
				err = sberr.Wrap(
					MissingLeadingZerosErr, "%s", val.Format(dt.fmt),
				)
			}
			break
		}

		// This is possibly the stupidest way possible to do this but this is
		// what the TOML library does and we have to match that behavior so we
		// are just going to look away
		str, _ := t.MarshalText()
		_ = val.UnmarshalText(str)

		return err
	}
}

// Useful when a type in the supplied config value is a custom type that
// implements the [encoding.TextUnmarshaler] interface which the TOML parser
// will use when parsing values.
//
// This is provided as a way to make sure the CMD line args can be parsed in the
// same manner as the values in the TOML file.
func FromTextUnmarshaler[T any, I interface {
	*T
	encoding.TextUnmarshaler
}](
	val *T,
	_default T,
) FlagSetFunc {
	*val = _default
	return func(arg string) error {
		return I(val).UnmarshalText([]byte(arg))
	}
}
