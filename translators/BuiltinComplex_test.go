package translators

import (
	"strconv"
	"testing"

	sbtest "github.com/barbell-math/smoothbrain-test"
)

func TestBuiltinComplex64Translator(t *testing.T) {
	var _t Translator[complex64]

	_t = BuiltinComplex64{}
	val, err := _t.Translate("123")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, complex64(123))

	val, err = _t.Translate("123i")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, complex64(123i))

	val, err = _t.Translate("123+456i")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, complex64(123+456i))

	_, err = _t.Translate("asdf")
	sbtest.ContainsError(t, strconv.ErrSyntax, err)
}

func TestBuiltinComplex128Translator(t *testing.T) {
	var _t Translator[complex128]

	_t = BuiltinComplex128{}
	val, err := _t.Translate("123")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, complex128(123))

	val, err = _t.Translate("123i")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, complex128(123i))

	val, err = _t.Translate("123+456i")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, complex128(123+456i))

	_, err = _t.Translate("asdf")
	sbtest.ContainsError(t, strconv.ErrSyntax, err)
}
