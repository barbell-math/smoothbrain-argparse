package translators

import (
	"strconv"
	"testing"

	sbtest "github.com/barbell-math/smoothbrain-test"
)

func TestBuiltinFloat32Translator(t *testing.T) {
	var _t Translator[float32]

	_t = BuiltinFloat32{}
	val, err := _t.Translate("1234.5678")
	sbtest.Nil(t, err)
	sbtest.EqFloat(t, 1234.5678, val, 1e-6)

	_, err = _t.Translate("asdf")
	sbtest.ContainsError(t, strconv.ErrSyntax, err)
}

func TestBuiltinFloat64Translator(t *testing.T) {
	var _t Translator[float64]

	_t = BuiltinFloat64{}
	val, err := _t.Translate("1234.5678")
	sbtest.Nil(t, err)
	sbtest.EqFloat(t, 1234.5678, val, 1e-6)

	_, err = _t.Translate("asdf")
	sbtest.ContainsError(t, strconv.ErrSyntax, err)
}
