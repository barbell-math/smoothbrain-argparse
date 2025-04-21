package translators

import (
	"testing"

	sbtest "github.com/barbell-math/smoothbrain-test"
)

func TestBuiltinUintTranslator(t *testing.T) {
	var _t Translator[uint]

	_t = BuiltinUint{Base: 2}
	val, err := _t.Translate("0b10")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, 2)

	_t = BuiltinUint{Base: 8}
	val, err = _t.Translate("13")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, 11)

	_t = BuiltinUint{Base: 10}
	val, err = _t.Translate("13")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, 13)

	_t = BuiltinUint{Base: 16}
	val, err = _t.Translate("0x13")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, 19)
}

func TestBuiltinUint8Translator(t *testing.T) {
	var _t Translator[uint8]

	_t = BuiltinUint8{Base: 2}
	val, err := _t.Translate("0b10")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, 2)

	_t = BuiltinUint8{Base: 8}
	val, err = _t.Translate("13")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, 11)

	_t = BuiltinUint8{Base: 10}
	val, err = _t.Translate("13")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, 13)

	_t = BuiltinUint8{Base: 16}
	val, err = _t.Translate("0x13")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, 19)
}

func TestBuiltinUint16Translator(t *testing.T) {
	var _t Translator[uint16]

	_t = BuiltinUint16{Base: 2}
	val, err := _t.Translate("0b10")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, 2)

	_t = BuiltinUint16{Base: 8}
	val, err = _t.Translate("13")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, 11)

	_t = BuiltinUint16{Base: 10}
	val, err = _t.Translate("13")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, 13)

	_t = BuiltinUint16{Base: 16}
	val, err = _t.Translate("0x13")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, 19)
}

func TestBuiltinUint32Translator(t *testing.T) {
	var _t Translator[uint32]

	_t = BuiltinUint32{Base: 2}
	val, err := _t.Translate("0b10")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, 2)

	_t = BuiltinUint32{Base: 8}
	val, err = _t.Translate("13")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, 11)

	_t = BuiltinUint32{Base: 10}
	val, err = _t.Translate("13")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, 13)

	_t = BuiltinUint32{Base: 16}
	val, err = _t.Translate("0x13")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, 19)
}

func TestBuiltinUint64Translator(t *testing.T) {
	var _t Translator[uint64]

	_t = BuiltinUint64{Base: 2}
	val, err := _t.Translate("0b10")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, 2)

	_t = BuiltinUint64{Base: 8}
	val, err = _t.Translate("13")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, 11)

	_t = BuiltinUint64{Base: 10}
	val, err = _t.Translate("13")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, 13)

	_t = BuiltinUint64{Base: 16}
	val, err = _t.Translate("0x13")
	sbtest.Nil(t, err)
	sbtest.Eq(t, val, 19)
}
