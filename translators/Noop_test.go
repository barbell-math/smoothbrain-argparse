package translators

import (
	"testing"

	sbtest "github.com/barbell-math/smoothbrain-test"
)

func TestNoopTranslator(t *testing.T) {
	var _t Translator[float32]

	_t = Noop[float32]{}

	val, err := _t.Translate("asf")
	sbtest.Nil(t, err)
	sbtest.Eq(t, float32(0), val)
}
