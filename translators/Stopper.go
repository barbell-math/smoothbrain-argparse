package translators

type (
	// Used to represent a flag that when encountered should stop parsing of the
	// cmd line arguments. The error that the stopper is created with will be
	// returned when [Stopper.Translate] is called.
	stopper[T any] struct{ err error }
)

func NewStopper[T any](e error) stopper[T] {
	return stopper[T]{err: e}
}

func (s stopper[T]) Translate(arg string) (T, error) {
	var tmp T
	return tmp, s.err
}
func (s stopper[T]) Reset() {}
func (s stopper[T]) HelpAddendum() string {
	return "Stops argument parsing when encountered"
}
