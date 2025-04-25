package computers

type (
	// Used to represent a comptuer that when encountered should stop parsing of
	// the computation of all further computer arguments. The error that the
	// stopper is created with will be returned when [Stopper.ComputeVals] is
	// called.
	stopper[T any] struct{ err error }
)

func NewStopper[T any](err error) stopper[T] {
	return stopper[T]{err: err}
}

func (s stopper[T]) ComputeVals() (T, error) {
	var tmp T
	return tmp, s.err
}

func (s stopper[T]) Reset() {}
