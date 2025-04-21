package translators

import (
	"testing"

	sbtest "github.com/barbell-math/smoothbrain-test"
)

func TestBuiltinIntTranslator(t *testing.T) {
	var _t Translator[int]

	_t = BuiltinInt{Base: 2}
	val, err := _t.Translate("0b10")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, 2)

	_t = BuiltinInt{Base: 8}
	val, err = _t.Translate("13")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, 11)

	_t = BuiltinInt{Base: 10}
	val, err = _t.Translate("13")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, 13)

	_t = BuiltinInt{Base: 16}
	val, err = _t.Translate("0x13")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, 19)
}

func TestBuiltinInt8Translator(t *testing.T) {
	var _t Translator[int8]

	_t = BuiltinInt8{Base: 2}
	val, err := _t.Translate("0b10")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, 2)

	_t = BuiltinInt8{Base: 8}
	val, err = _t.Translate("13")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, 11)

	_t = BuiltinInt8{Base: 10}
	val, err = _t.Translate("13")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, 13)

	_t = BuiltinInt8{Base: 16}
	val, err = _t.Translate("0x13")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, 19)
}

func TestBuiltinInt16Translator(t *testing.T) {
	var _t Translator[int16]

	_t = BuiltinInt16{Base: 2}
	val, err := _t.Translate("0b10")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, 2)

	_t = BuiltinInt16{Base: 8}
	val, err = _t.Translate("13")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, 11)

	_t = BuiltinInt16{Base: 10}
	val, err = _t.Translate("13")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, 13)

	_t = BuiltinInt16{Base: 16}
	val, err = _t.Translate("0x13")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, 19)
}

func TestBuiltinInt32Translator(t *testing.T) {
	var _t Translator[int32]

	_t = BuiltinInt32{Base: 2}
	val, err := _t.Translate("0b10")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, 2)

	_t = BuiltinInt32{Base: 8}
	val, err = _t.Translate("13")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, 11)

	_t = BuiltinInt32{Base: 10}
	val, err = _t.Translate("13")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, 13)

	_t = BuiltinInt32{Base: 16}
	val, err = _t.Translate("0x13")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, 19)
}

func TestBuiltinInt64Translator(t *testing.T) {
	var _t Translator[int64]

	_t = BuiltinInt64{Base: 2}
	val, err := _t.Translate("0b10")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, 2)

	_t = BuiltinInt64{Base: 8}
	val, err = _t.Translate("13")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, 11)

	_t = BuiltinInt64{Base: 10}
	val, err = _t.Translate("13")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, 13)

	_t = BuiltinInt64{Base: 16}
	val, err = _t.Translate("0x13")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, 19)
}
