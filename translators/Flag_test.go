package translators

import (
	"testing"

	sbtest "github.com/barbell-math/smoothbrain-test"
)

func TestFlagTranslator(t *testing.T) {
	var _t Translator[bool]

	_t = Flag{}

	val, err := _t.Translate("asf")
	sbtest.Nil(t, err)
	sbtest.True(t, val)
}
