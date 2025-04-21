package translators

import (
	"math/big"
	"testing"

	sbtest "github.com/barbell-math/smoothbrain-test"
)

func TestBigIntTranslator(t *testing.T) {
	var _t Translator[big.Int]

	_t = BigInt{}
	val, err := _t.Translate("1234")
	sbtest.Nil(t, err)
	sbtest.Eq(t, 0, val.Cmp(big.NewInt(1234)))

	_, err = _t.Translate("asdf")
	sbtest.ContainsError(t, CouldNotParseBigIntErr, err)
}

func TestBigRatTranslator(t *testing.T) {
	var _t Translator[big.Rat]

	_t = BigRat{}
	val, err := _t.Translate("1234")
	sbtest.Nil(t, err)
	sbtest.Eq(t, 0, val.Cmp(big.NewRat(1234, 1)))

	_, err = _t.Translate("asdf")
	sbtest.ContainsError(t, CouldNotParseBigRatErr, err)
}

func TestBigFloatTranslator(t *testing.T) {
	var _t Translator[big.Float]

	_t = BigFloat{}
	val, err := _t.Translate("1234")
	sbtest.Nil(t, err)
	sbtest.Eq(t, 0, val.Cmp(big.NewFloat(1234)))

	_, err = _t.Translate("asdf")
	sbtest.ContainsError(t, CouldNotParseBigFloatErr, err)
}
