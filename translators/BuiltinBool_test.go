package translators

import (
	"strconv"
	"testing"

	sbtest "github.com/barbell-math/smoothbrain-test"
)

func TestBuilitinBoolTranslator(t *testing.T) {
	var _t Translator[bool]

	_t = BuiltinBool{}
	val, err := _t.Translate("f")
	sbtest.Nil(t, err)
	sbtest.False(t, val)

	val, err = _t.Translate("t")
	sbtest.Nil(t, err)
	sbtest.True(t, val)

	_, err = _t.Translate("asdf")
	sbtest.ContainsError(t, strconv.ErrSyntax, err)
}
