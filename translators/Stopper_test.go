package translators

import (
	"errors"
	"testing"

	sbtest "github.com/barbell-math/smoothbrain-test"
)

func TestStopperTranslator(t *testing.T) {
	var testErr1 = errors.New("err1")

	var _t Translator[float32]

	_t = NewStopper[float32](testErr1)

	val, err := _t.Translate("asf")
	sbtest.ContainsError(t, testErr1, err)
	sbtest.Eq(t, float32(0), val)
}
