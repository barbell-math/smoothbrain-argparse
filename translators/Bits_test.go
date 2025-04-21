package translators

import (
	"testing"

	sbtest "github.com/barbell-math/smoothbrain-test"
)

func TestBitsTranslator(t *testing.T) {
	var _t Translator[[]byte]
	_t = Bits{}

	_, err := _t.Translate("00asdf")
	sbtest.ContainsError(t, BitsTranslationErr, err)

	res, err := _t.Translate("00000000")
	sbtest.Nil(t, err)
	sbtest.SlicesMatch[byte](t, []byte{0, 0}, res)

	res, err = _t.Translate("10101010")
	sbtest.Nil(t, err)
	sbtest.SlicesMatch[byte](t, []byte{0b1010, 0b1010}, res)

	res, err = _t.Translate("10101010")
	sbtest.Nil(t, err)
	sbtest.SlicesMatch[byte](t, []byte{0b1010, 0b1010}, res)

	res, err = _t.Translate("10100101")
	sbtest.Nil(t, err)
	sbtest.SlicesMatch[byte](t, []byte{0b1010, 0b0101}, res)

	res, err = _t.Translate("101001011010")
	sbtest.Nil(t, err)
	sbtest.SlicesMatch[byte](t, []byte{0b1010, 0b0101, 0b1010}, res)

	res, err = _t.Translate("1 0100 1011 0101")
	sbtest.Nil(t, err)
	sbtest.SlicesMatch[byte](t, []byte{0b0001, 0b0100, 0b1011, 0b0101}, res)

	res, err = _t.Translate("10 0100 1011 0101")
	sbtest.Nil(t, err)
	sbtest.SlicesMatch[byte](t, []byte{0b0010, 0b0100, 0b1011, 0b0101}, res)

	res, err = _t.Translate("100 0100 1011 0101")
	sbtest.Nil(t, err)
	sbtest.SlicesMatch[byte](t, []byte{0b0100, 0b0100, 0b1011, 0b0101}, res)

	res, err = _t.Translate("")
	sbtest.Nil(t, err)
	sbtest.SlicesMatch[byte](t, []byte{}, res)

	res, err = _t.Translate("1")
	sbtest.Nil(t, err)
	sbtest.SlicesMatch[byte](t, []byte{0b0001}, res)

	res, err = _t.Translate("10")
	sbtest.Nil(t, err)
	sbtest.SlicesMatch[byte](t, []byte{0b0010}, res)

	res, err = _t.Translate("100")
	sbtest.Nil(t, err)
	sbtest.SlicesMatch[byte](t, []byte{0b0100}, res)
}
