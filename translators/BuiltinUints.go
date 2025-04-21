package translators

import (
	"fmt"
	"math"
	"strconv"
)

type (
	// Represents a cmd line argument that will be translated to a uint type.
	BuiltinUint struct {
		Base int
	}

	// Represents a cmd line argument that will be translated to a uint8 type.
	BuiltinUint8 struct {
		Base int
	}

	// Represents a cmd line argument that will be translated to a uint16 type.
	BuiltinUint16 struct {
		Base int
	}

	// Represents a cmd line argument that will be translated to a uint32 type.
	BuiltinUint32 struct {
		Base int
	}

	// Represents a cmd line argument that will be translated to a uint64 type.
	BuiltinUint64 struct {
		Base int
	}
)

func (u BuiltinUint) Translate(arg string) (uint, error) {
	// bit size of 0 corresponds with uint
	var u64 uint64
	var err error
	if u.Base == 2 && len(arg) > 2 && arg[0:2] == "0b" {
		u64, err = strconv.ParseUint(arg[2:], u.Base, 0)
	} else if u.Base == 8 && len(arg) > 2 && arg[0:2] == "0o" {
		u64, err = strconv.ParseUint(arg[2:], u.Base, 0)
	} else if u.Base == 16 && len(arg) > 2 && arg[0:2] == "0x" {
		u64, err = strconv.ParseUint(arg[2:], u.Base, 0)
	} else {
		u64, err = strconv.ParseUint(arg, u.Base, 0)
	}
	return uint(u64), err
}
func (u BuiltinUint) Reset() {}
func (u BuiltinUint) HelpAddendum() string {
	return fmt.Sprintf(
		"Expecting int with base %d in range [%d, %d]",
		u.Base, 0, uint64(math.MaxUint),
	)
}

func (u BuiltinUint8) Translate(arg string) (uint8, error) {
	var u64 uint64
	var err error
	if u.Base == 2 && len(arg) > 2 && arg[0:2] == "0b" {
		u64, err = strconv.ParseUint(arg[2:], u.Base, 8)
	} else if u.Base == 8 && len(arg) > 2 && arg[0:2] == "0o" {
		u64, err = strconv.ParseUint(arg[2:], u.Base, 8)
	} else if u.Base == 16 && len(arg) > 2 && arg[0:2] == "0x" {
		u64, err = strconv.ParseUint(arg[2:], u.Base, 8)
	} else {
		u64, err = strconv.ParseUint(arg, u.Base, 8)
	}
	return uint8(u64), err
}
func (u BuiltinUint8) Reset() {}
func (u BuiltinUint8) HelpAddendum() string {
	return fmt.Sprintf(
		"Expecting int with base %d in range [%d, %d]",
		u.Base, 0, math.MaxUint8,
	)
}

func (u BuiltinUint16) Translate(arg string) (uint16, error) {
	var u64 uint64
	var err error
	if u.Base == 2 && len(arg) > 2 && arg[0:2] == "0b" {
		u64, err = strconv.ParseUint(arg[2:], u.Base, 16)
	} else if u.Base == 8 && len(arg) > 2 && arg[0:2] == "0o" {
		u64, err = strconv.ParseUint(arg[2:], u.Base, 16)
	} else if u.Base == 16 && len(arg) > 2 && arg[0:2] == "0x" {
		u64, err = strconv.ParseUint(arg[2:], u.Base, 16)
	} else {
		u64, err = strconv.ParseUint(arg, u.Base, 16)
	}
	return uint16(u64), err
}
func (u BuiltinUint16) Reset() {}
func (u BuiltinUint16) HelpAddendum() string {
	return fmt.Sprintf(
		"Expecting int with base %d in range [%d, %d]",
		u.Base, 0, math.MaxUint16,
	)
}

func (u BuiltinUint32) Translate(arg string) (uint32, error) {
	var u64 uint64
	var err error
	if u.Base == 2 && len(arg) > 2 && arg[0:2] == "0b" {
		u64, err = strconv.ParseUint(arg[2:], u.Base, 32)
	} else if u.Base == 8 && len(arg) > 2 && arg[0:2] == "0o" {
		u64, err = strconv.ParseUint(arg[2:], u.Base, 32)
	} else if u.Base == 16 && len(arg) > 2 && arg[0:2] == "0x" {
		u64, err = strconv.ParseUint(arg[2:], u.Base, 32)
	} else {
		u64, err = strconv.ParseUint(arg, u.Base, 32)
	}
	return uint32(u64), err
}
func (u BuiltinUint32) Reset() {}
func (u BuiltinUint32) HelpAddendum() string {
	return fmt.Sprintf(
		"Expecting int with base %d in range [%d, %d]",
		u.Base, 0, math.MaxUint32,
	)
}

func (u BuiltinUint64) Translate(arg string) (uint64, error) {
	var u64 uint64
	var err error
	if u.Base == 2 && len(arg) > 2 && arg[0:2] == "0b" {
		u64, err = strconv.ParseUint(arg[2:], u.Base, 64)
	} else if u.Base == 8 && len(arg) > 2 && arg[0:2] == "0o" {
		u64, err = strconv.ParseUint(arg[2:], u.Base, 64)
	} else if u.Base == 16 && len(arg) > 2 && arg[0:2] == "0x" {
		u64, err = strconv.ParseUint(arg[2:], u.Base, 64)
	} else {
		u64, err = strconv.ParseUint(arg, u.Base, 64)
	}
	return u64, err
}
func (u BuiltinUint64) Reset() {}
func (u BuiltinUint64) HelpAddendum() string {
	return fmt.Sprintf(
		"Expecting int with base %d in range [%d, %d]",
		u.Base, 0, uint64(math.MaxUint64),
	)
}
