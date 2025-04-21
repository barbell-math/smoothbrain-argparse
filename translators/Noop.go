package translators

type (
	// Used to represent a cmd line argument that will always be the default
	// zero-value initialized T type.
	Noop[T any] struct{}
)

func (_ Noop[T]) Translate(arg string) (T, error) {
	var tmp T
	return tmp, nil
}
func (_ Noop[T]) Reset()               {}
func (_ Noop[T]) HelpAddendum() string { return "" }
