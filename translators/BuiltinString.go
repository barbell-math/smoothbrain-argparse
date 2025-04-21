package translators

type (
	// Represents a cmd line argument that will be translated to a string type.
	BuiltinString struct{}
)

func (_ BuiltinString) Translate(arg string) (string, error) {
	return arg, nil
}
func (_ BuiltinString) Reset()               {}
func (_ BuiltinString) HelpAddendum() string { return "" }
