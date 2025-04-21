package translators

import (
	"strconv"
	"testing"

	sbmap "github.com/barbell-math/smoothbrain-hashmap"
	sbtest "github.com/barbell-math/smoothbrain-test"
)

func TestListTranslator(t *testing.T) {
	var _t Translator[[]int]

	_t = NewList[BuiltinInt](BuiltinInt{Base: 10})
	val, err := _t.Translate("1")
	sbtest.Nil(t, err)
	sbtest.SlicesMatch(t, []int{1}, val)

	val, err = _t.Translate("2")
	sbtest.Nil(t, err)
	sbtest.SlicesMatch(t, []int{1, 2}, val)

	val, err = _t.Translate("3")
	sbtest.Nil(t, err)
	sbtest.SlicesMatch(t, []int{1, 2, 3}, val)

	_t.Reset()
	val, err = _t.Translate("1")
	sbtest.Nil(t, err)
	sbtest.SlicesMatch(t, []int{1}, val)

	_t.Reset()
	_, err = _t.Translate("asdf")
	sbtest.ContainsError(t, strconv.ErrSyntax, err)
}

func TestListSelectorTranslator(t *testing.T) {
	var _t Translator[[]int]

	m := sbmap.New[int, struct{}]()
	m.Put(1, struct{}{})
	m.Put(2, struct{}{})
	m.Put(3, struct{}{})
	_t = NewListSelector[BuiltinInt](BuiltinInt{Base: 10}, m)
	val, err := _t.Translate("1")
	sbtest.Nil(t, err)
	sbtest.SlicesMatch(t, []int{1}, val)

	val, err = _t.Translate("2")
	sbtest.Nil(t, err)
	sbtest.SlicesMatch(t, []int{1, 2}, val)

	val, err = _t.Translate("3")
	sbtest.Nil(t, err)
	sbtest.SlicesMatch(t, []int{1, 2, 3}, val)

	val, err = _t.Translate("4")
	sbtest.ContainsError(t, ValNotInAllowedListErr, err)
	sbtest.SlicesMatch(t, []int{1, 2, 3}, val)

	_t.Reset()
	val, err = _t.Translate("1")
	sbtest.Nil(t, err)
	sbtest.SlicesMatch(t, []int{1}, val)

	_t.Reset()
	_, err = _t.Translate("asdf")
	sbtest.ContainsError(t, strconv.ErrSyntax, err)
}
