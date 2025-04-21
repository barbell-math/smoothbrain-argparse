package translators

import (
	"testing"

	sbtest "github.com/barbell-math/smoothbrain-test"
)

func TestBuiltinStringTranslator(t *testing.T) {
	var _t Translator[string]

	_t = BuiltinString{}
	val, err := _t.Translate("asdf")
	sbtest.Nil(t, err)
	sbtest.Eq(t, "asdf", val)
}
