package translators

import (
	"errors"
	"fmt"
	"slices"

	sberr "github.com/barbell-math/smoothbrain-errs"
	sbmap "github.com/barbell-math/smoothbrain-hashmap"
)

type (
	// A translator that collects all supplied values into a slice, optionally
	// restricting the set of allowed values.
	listValues[T Translator[U], U any] struct {
		valueTranslator T
		allowedVals     sbmap.Map[U, struct{}]
		vals            []U
	}
)

var (
	ValNotInAllowedListErr = errors.New(
		"Value was not found in the allowed list",
	)
)

// Creates a new list translator with no restrictions on the set of allowed
// values.
func NewList[T Translator[U], U any](
	translator T,
) *listValues[T, U] {
	return &listValues[T, U]{
		valueTranslator: translator,
	}
}

// Creates a new list translator with potential restrictions on the set of
// allowed values. If the `allowedValues` map is empty there will be no
// restrictions on the set of allowed values. If the `allowedValues` map has
// values the set of allowed values will be restricted to what is in the map.
func NewListSelector[T Translator[U], U any](
	translator T,
	allowedValues sbmap.Map[U, struct{}],
) *listValues[T, U] {
	return &listValues[T, U]{
		valueTranslator: translator,
		allowedVals:     allowedValues,
	}
}

func (l *listValues[T, U]) Translate(arg string) ([]U, error) {
	v, err := l.valueTranslator.Translate(arg)
	if err != nil {
		return l.vals, err
	}
	if l.allowedVals.Len() > 0 {
		_, ok := l.allowedVals.Get(v)
		if !ok {
			return l.vals, sberr.WrapValueList(
				ValNotInAllowedListErr,
				"The supplied value must be found in the list shown below",
				sberr.WrapListVal{"Supplied value", v},
				sberr.WrapListVal{"Allowed values", &l.allowedVals},
			)
		}
	}

	l.vals = append(l.vals, v)
	return l.vals, nil
}

func (l *listValues[T, U]) Reset() {
	l.valueTranslator.Reset()
	l.vals = []U{}
}

func (l *listValues[T, U]) HelpAddendum() string {
	if l.allowedVals.Len() == 0 {
		return fmt.Sprintf(
			"Expecting a list of the following:\n%s",
			l.valueTranslator.HelpAddendum(),
		)
	} else {
		return fmt.Sprintf(
			"Expecting a list of the following values:\n%v",
			slices.Collect(l.allowedVals.Keys()),
		)
	}
}
