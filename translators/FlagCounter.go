package translators

import (
	"errors"
	"fmt"
	"golang.org/x/exp/constraints"

	sberr "github.com/barbell-math/smoothbrain-errs"
)

type (
	// Used to represent a flag that can be supplied many times, with a counter
	// incrementing each time the flag is encountered.
	FlagCntr[
		T constraints.Integer | constraints.Float | constraints.Complex,
	] struct {
		cntr T
	}

	// Used to represent a flag that can be supplied many times up to the
	// provided maximum number of times. A counter will be incremented each
	// time the flag is encountered.
	LimitedFlagCntr[T constraints.Integer | constraints.Float] struct {
		FlagCntr[T]
		maxTimes T
	}
)

var (
	FlagProvidedToManyTimesErr = errors.New("A flag was provided to many times.")
)

func (f *FlagCntr[T]) Translate(arg string) (T, error) {
	f.cntr++
	return f.cntr, nil
}
func (f *FlagCntr[T]) Reset() {
	f.cntr = 0
}
func (f *FlagCntr[T]) HelpAddendum() string {
	return "Can be supplied may times"
}

func NewLimitedFlagCntr[
	T constraints.Integer | constraints.Float,
](maxTimes T) *LimitedFlagCntr[T] {
	return &LimitedFlagCntr[T]{maxTimes: maxTimes}
}
func (f *LimitedFlagCntr[T]) Translate(arg string) (T, error) {
	if f.FlagCntr.cntr >= f.maxTimes {
		return f.FlagCntr.cntr, sberr.Wrap(
			FlagProvidedToManyTimesErr,
			"Maximum allowed: %v", f.maxTimes,
		)
	}
	rv, err := f.FlagCntr.Translate(arg)
	if err != nil {
		return rv, nil
	}
	return rv, nil
}
func (f *LimitedFlagCntr[T]) Reset() {
	f.FlagCntr.cntr = 0
}
func (f *LimitedFlagCntr[T]) HelpAddendum() string {
	return fmt.Sprintf(
		"Can be supplied between 0 and %v times (inclusive)", f.maxTimes,
	)
}
