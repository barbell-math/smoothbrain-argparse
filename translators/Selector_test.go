package translators

import (
	"testing"

	sbmap "github.com/barbell-math/smoothbrain-hashmap"
	sbtest "github.com/barbell-math/smoothbrain-test"
)

func TestSelectorPassing(t *testing.T) {
	m := sbmap.New[int, struct{}]()
	m.Put(1, struct{}{})
	s := NewSelector(BuiltinInt{Base: 10}, m)

	res, err := s.Translate("1")
	sbtest.Nil(t, err)
	sbtest.Eq(t, res, 1)
}

func TestSelectorFailing(t *testing.T) {
	m := sbmap.New[int, struct{}]()
	m.Put(1, struct{}{})
	m.Put(2, struct{}{})
	m.Put(3, struct{}{})
	s := NewSelector(BuiltinInt{Base: 10}, m)

	_, err := s.Translate("4")
	sbtest.ContainsError(t, ValNotInAllowedListErr, err)
}
