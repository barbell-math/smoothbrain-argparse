package translators

import (
	"fmt"
	"math"
	"strconv"
)

type (
	// Represents a cmd line argument that will be translated to a int type.
	BuiltinInt struct {
		Base int
	}

	// Represents a cmd line argument that will be translated to a int8 type.
	BuiltinInt8 struct {
		Base int
	}

	// Represents a cmd line argument that will be translated to a int16 type.
	BuiltinInt16 struct {
		Base int
	}

	// Represents a cmd line argument that will be translated to a int32 type.
	BuiltinInt32 struct {
		Base int
	}

	// Represents a cmd line argument that will be translated to a int64 type.
	BuiltinInt64 struct {
		Base int
	}
)

func (i BuiltinInt) Translate(arg string) (int, error) {
	// bit size of 0 corresponds with int
	var i64 int64
	var err error
	if i.Base == 2 && len(arg) > 2 && arg[0:2] == "0b" {
		i64, err = strconv.ParseInt(arg[2:], i.Base, 0)
	} else if i.Base == 8 && len(arg) > 2 && arg[0:2] == "0o" {
		i64, err = strconv.ParseInt(arg[2:], i.Base, 0)
	} else if i.Base == 16 && len(arg) > 2 && arg[0:2] == "0x" {
		i64, err = strconv.ParseInt(arg[2:], i.Base, 0)
	} else {
		i64, err = strconv.ParseInt(arg, i.Base, 0)
	}
	return int(i64), err
}
func (i BuiltinInt) Reset() {}
func (i BuiltinInt) HelpAddendum() string {
	return fmt.Sprintf(
		"Expecting int with base %d in range [%d, %d]",
		i.Base, math.MinInt, math.MaxInt,
	)
}

func (i BuiltinInt8) Translate(arg string) (int8, error) {
	var i64 int64
	var err error
	if i.Base == 2 && len(arg) > 2 && arg[0:2] == "0b" {
		i64, err = strconv.ParseInt(arg[2:], i.Base, 8)
	} else if i.Base == 8 && len(arg) > 2 && arg[0:2] == "0o" {
		i64, err = strconv.ParseInt(arg[2:], i.Base, 8)
	} else if i.Base == 16 && len(arg) > 2 && arg[0:2] == "0x" {
		i64, err = strconv.ParseInt(arg[2:], i.Base, 8)
	} else {
		i64, err = strconv.ParseInt(arg, i.Base, 8)
	}
	return int8(i64), err
}
func (i BuiltinInt8) Reset() {}
func (i BuiltinInt8) HelpAddendum() string {
	return fmt.Sprintf(
		"Expecting int with base %d in range [%d, %d]",
		i.Base, math.MinInt8, math.MaxInt8,
	)
}

func (i BuiltinInt16) Translate(arg string) (int16, error) {
	var i64 int64
	var err error
	if i.Base == 2 && len(arg) > 2 && arg[0:2] == "0b" {
		i64, err = strconv.ParseInt(arg[2:], i.Base, 16)
	} else if i.Base == 8 && len(arg) > 2 && arg[0:2] == "0o" {
		i64, err = strconv.ParseInt(arg[2:], i.Base, 16)
	} else if i.Base == 16 && len(arg) > 2 && arg[0:2] == "0x" {
		i64, err = strconv.ParseInt(arg[2:], i.Base, 16)
	} else {
		i64, err = strconv.ParseInt(arg, i.Base, 16)
	}
	return int16(i64), err
}
func (i BuiltinInt16) Reset() {}
func (i BuiltinInt16) HelpAddendum() string {
	return fmt.Sprintf(
		"Expecting int with base %d in range [%d, %d]",
		i.Base, math.MinInt16, math.MaxInt16,
	)
}

func (i BuiltinInt32) Translate(arg string) (int32, error) {
	var i64 int64
	var err error
	if i.Base == 2 && len(arg) > 2 && arg[0:2] == "0b" {
		i64, err = strconv.ParseInt(arg[2:], i.Base, 32)
	} else if i.Base == 8 && len(arg) > 2 && arg[0:2] == "0o" {
		i64, err = strconv.ParseInt(arg[2:], i.Base, 32)
	} else if i.Base == 16 && len(arg) > 2 && arg[0:2] == "0x" {
		i64, err = strconv.ParseInt(arg[2:], i.Base, 32)
	} else {
		i64, err = strconv.ParseInt(arg, i.Base, 32)
	}
	return int32(i64), err
}
func (i BuiltinInt32) Reset() {}
func (i BuiltinInt32) HelpAddendum() string {
	return fmt.Sprintf(
		"Expecting int with base %d in range [%d, %d]",
		i.Base, math.MinInt32, math.MaxInt32,
	)
}

func (i BuiltinInt64) Translate(arg string) (int64, error) {
	var i64 int64
	var err error
	if i.Base == 2 && len(arg) > 2 && arg[0:2] == "0b" {
		i64, err = strconv.ParseInt(arg[2:], i.Base, 64)
	} else if i.Base == 8 && len(arg) > 2 && arg[0:2] == "0o" {
		i64, err = strconv.ParseInt(arg[2:], i.Base, 64)
	} else if i.Base == 16 && len(arg) > 2 && arg[0:2] == "0x" {
		i64, err = strconv.ParseInt(arg[2:], i.Base, 64)
	} else {
		i64, err = strconv.ParseInt(arg, i.Base, 64)
	}
	return i64, err
}
func (i BuiltinInt64) Reset() {}
func (i BuiltinInt64) HelpAddendum() string {
	return fmt.Sprintf(
		"Expecting int with base %d in range [%d, %d]",
		i.Base, math.MinInt64, math.MaxInt64,
	)
}
