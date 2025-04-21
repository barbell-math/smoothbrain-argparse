package translators

import (
	"testing"

	sbtest "github.com/barbell-math/smoothbrain-test"
)

func TestFlagCntrTranslator(t *testing.T) {
	var _t Translator[int]

	_t = &FlagCntr[int]{}
	c, err := _t.Translate("asdf")
	sbtest.Nil(t, err)
	sbtest.Eq(t, c, 1)

	c, err = _t.Translate("asdf")
	sbtest.Nil(t, err)
	sbtest.Eq(t, c, 2)

	c, err = _t.Translate("asdf")
	sbtest.Nil(t, err)
	sbtest.Eq(t, c, 3)

	c, err = _t.Translate("asdf")
	sbtest.Nil(t, err)
	sbtest.Eq(t, c, 4)

	_t.Reset()

	c, err = _t.Translate("asdf")
	sbtest.Nil(t, err)
	sbtest.Eq(t, c, 1)
}

func TestLimitedFlagCntrTranslator(t *testing.T) {
	var _t Translator[int]

	_t = NewLimitedFlagCntr(3)
	c, err := _t.Translate("asdf")
	sbtest.Nil(t, err)
	sbtest.Eq(t, c, 1)

	c, err = _t.Translate("asdf")
	sbtest.Nil(t, err)
	sbtest.Eq(t, c, 2)

	c, err = _t.Translate("asdf")
	sbtest.Nil(t, err)
	sbtest.Eq(t, c, 3)

	c, err = _t.Translate("asdf")
	sbtest.ContainsError(t, FlagProvidedToManyTimesErr, err)
	sbtest.Eq(t, c, 3)
}
