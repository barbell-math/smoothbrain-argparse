package translators

import "strconv"

type (
	// Represents a cmd line argument that will be translated to a bool type.
	BuiltinBool struct{}
)

func (_ BuiltinBool) Translate(arg string) (bool, error) {
	b, err := strconv.ParseBool(arg)
	return b, err
}
func (_ BuiltinBool) Reset() {}
func (_ BuiltinBool) HelpAddendum() string {
	return "One of 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False"
}
