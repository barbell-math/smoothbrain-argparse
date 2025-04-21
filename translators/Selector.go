package translators

import (
	"fmt"
	"slices"

	sberr "github.com/barbell-math/smoothbrain-errs"
	sbmap "github.com/barbell-math/smoothbrain-hashmap"
)

type (
	// A translator that imposes a set of specific values on a cmd line
	// argument.
	selector[T Translator[U], U any] struct {
		allowedVals     sbmap.Map[U, struct{}]
		valueTranslator Translator[U]
	}
)

func NewSelector[T Translator[U], U any](
	translator T,
	allowedValues sbmap.Map[U, struct{}],
) *selector[T, U] {
	return &selector[T, U]{
		valueTranslator: translator,
		allowedVals:     allowedValues,
	}
}

func (s selector[T, U]) Translate(arg string) (U, error) {
	rv, err := s.valueTranslator.Translate(arg)
	if err != nil {
		return rv, err
	}
	_, ok := s.allowedVals.Get(rv)
	if !ok {
		return rv, sberr.WrapValueList(
			ValNotInAllowedListErr,
			"The supplied value must be found in the list shown below",
			sberr.WrapListVal{"Supplied value", rv},
			sberr.WrapListVal{"Allowed values", &s.allowedVals},
		)
	}
	return rv, nil
}

func (s selector[T, U]) Reset() {
	s.valueTranslator.Reset()
}
func (s selector[T, U]) HelpAddendum() string {
	return fmt.Sprintf(
		"Expecting one of the following values:\n%v",
		slices.Collect(s.allowedVals.Keys()),
	)
}
