package translators

import (
	"errors"
	"math/big"
)

type (
	// Represents a cmd line argument that will be translated to a [big.Int]
	// type.
	BigInt struct{}

	// Represents a cmd line argument that will be translated to a [big.Rat]
	// type.
	BigRat struct{}

	// Represents a cmd line argument that will be translated to a [big.Float]
	// type.
	BigFloat struct{}
)

var (
	CouldNotParseBigIntErr = errors.New(
		"Could not parse the supplied string as a big int",
	)
	CouldNotParseBigRatErr = errors.New(
		"Could not parse the supplied string as a big rat",
	)
	CouldNotParseBigFloatErr = errors.New(
		"Could not parse the supplied string as a big float",
	)
)

func (_ BigInt) Translate(arg string) (big.Int, error) {
	rv := big.Int{}
	_, ok := rv.SetString(arg, 10)
	if !ok {
		return big.Int{}, CouldNotParseBigIntErr
	}
	return rv, nil
}
func (_ BigInt) Reset()               {}
func (_ BigInt) HelpAddendum() string { return "" }

func (_ BigFloat) Translate(arg string) (big.Float, error) {
	rv := big.Float{}
	_, ok := rv.SetString(arg)
	if !ok {
		return big.Float{}, CouldNotParseBigFloatErr
	}
	return rv, nil
}
func (_ BigFloat) Reset()               {}
func (_ BigFloat) HelpAddendum() string { return "" }

func (_ BigRat) Translate(arg string) (big.Rat, error) {
	rv := big.Rat{}
	_, ok := rv.SetString(arg)
	if !ok {
		return big.Rat{}, CouldNotParseBigRatErr
	}
	return rv, nil
}
func (_ BigRat) Reset()               {}
func (_ BigRat) HelpAddendum() string { return "" }
