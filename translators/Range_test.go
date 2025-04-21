package translators

import (
	"testing"

	sbtest "github.com/barbell-math/smoothbrain-test"
)

func TestRangeTranslator(t *testing.T) {
	var _t Translator[int]

	_t = NewRange[BuiltinInt](
		3, 5, BuiltinInt{Base: 10},
	)

	_, err := _t.Translate("2")
	sbtest.ContainsError(t, ValOutsideRange, err)

	v, err := _t.Translate("3")
	sbtest.Nil(t, err)
	sbtest.Eq(t, v, 3)
	v, err = _t.Translate("4")
	sbtest.Nil(t, err)
	sbtest.Eq(t, v, 4)

	_, err = _t.Translate("5")
	sbtest.ContainsError(t, ValOutsideRange, err)
	_, err = _t.Translate("6")
	sbtest.ContainsError(t, ValOutsideRange, err)
}
