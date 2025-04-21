package translators

import "strconv"

type (
	// Represents a cmd line argument that will be translated to a complex64 type.
	BuiltinComplex64 struct{}

	// Represents a cmd line argument that will be translated to a complex128 type.
	BuiltinComplex128 struct{}
)

func (_ BuiltinComplex64) Translate(arg string) (complex64, error) {
	c64, err := strconv.ParseComplex(arg, 64)
	return complex64(c64), err
}
func (_ BuiltinComplex64) Reset() {}
func (_ BuiltinComplex64) HelpAddendum() string {
	return "The number represented by s must be of the form N, Ni, or N±Ni, where N stands for a floating-point number as recognized by ParseFloat, and i is the imaginary component.\nIf the second N is unsigned, a + sign is required between the two components as indicated by the ±.\nIf the second N is NaN, only a + sign is accepted."
}

func (_ BuiltinComplex128) Translate(arg string) (complex128, error) {
	c128, err := strconv.ParseComplex(arg, 32)
	return c128, err
}
func (_ BuiltinComplex128) Reset() {}
func (_ BuiltinComplex128) HelpAddendum() string {
	return "The number represented by s must be of the form N, Ni, or N±Ni, where N stands for a floating-point number as recognized by ParseFloat, and i is the imaginary component.\nIf the second N is unsigned, a + sign is required between the two components as indicated by the ±.\nIf the second N is NaN, only a + sign is accepted."
}
