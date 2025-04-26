package sbargp

import (
	"encoding"
	"strconv"

	"golang.org/x/exp/constraints"
)

type (
	// The type of function that [flag.Func] accepts
	flagSetFunc func(arg string) error
)

// Useful for parsing a specific kind of uint from the CMD line since flag does
// not have a generic version yet. (It only provides uint)
func Uint[T constraints.Unsigned](val *T, base int) flagSetFunc {
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
func Int[T constraints.Signed](val *T, base int) flagSetFunc {
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
func Float[T constraints.Float](val *T) flagSetFunc {
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

// Useful when a type in the supplied config value is a custom type that
// implements the [encoding.FromTextUnmarshaler] interface which the TOML parser
// will use when parsing values.
//
// This is provided as a way to make sure the CMD line args can be parsed in the
// same manner as the values in the TOML file.
func FromTextUnmarshaler(val encoding.TextUnmarshaler) flagSetFunc {
	return func(arg string) error {
		return val.UnmarshalText([]byte(arg))
	}
}
