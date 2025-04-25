package translators

import (
	"errors"
	"fmt"
	"reflect"

	sberr "github.com/barbell-math/smoothbrain-errs"
	"golang.org/x/exp/constraints"
)

type (
	// A translator that imposes a range on the supplied cmd line argument.
	_range[
		T Translator[U],
		U constraints.Integer | constraints.Float,
	] struct {
		min           U // Inclusive min
		max           U // Exclusive max
		numTranslator T
	}
)

var (
	ValOutsideRange = errors.New(
		"The supplied value is outside the allowed range.",
	)
)

func NewRange[
	T Translator[U],
	U constraints.Signed | constraints.Unsigned | constraints.Float,
](min U, max U, translator T) _range[T, U] {
	return _range[T, U]{
		min:           min,
		max:           max,
		numTranslator: translator,
	}
}

func (r _range[T, U]) Translate(arg string) (U, error) {
	rv, err := r.numTranslator.Translate(arg)
	if err != nil {
		return rv, err
	}

	if rv < r.min || rv >= r.max {
		return rv, sberr.WrapValueList(
			ValOutsideRange,
			"",
			sberr.WrapListVal{ItemName: "Min (inclusive)", Item: r.min},
			sberr.WrapListVal{ItemName: "Max (exclusive)", Item: r.max},
		)
	}

	return rv, nil
}

func (r _range[T, U]) Reset() {
	r.numTranslator.Reset()
}

func (r _range[T, U]) HelpAddendum() string {
	return fmt.Sprintf(
		"Expecting %s within range [%v, %v]",
		reflect.TypeFor[T]().String(), r.min, r.max,
	)
}
