package translators

import (
	"net/mail"
	"testing"

	sbtest "github.com/barbell-math/smoothbrain-test"
)

func TestEmailTranslator(t *testing.T) {
	var _t Translator[mail.Address]

	_t = Email{}
	val, err := _t.Translate("test@test.com")
	sbtest.Nil(t, err)
	sbtest.Eq(t, "test@test.com", val.Address)
	sbtest.Eq(t, "", val.Name)

	_, err = _t.Translate("asdf")
	sbtest.NotNil(t, err)
}

func TestStrEmailTranslator(t *testing.T) {
	var _t Translator[string]

	_t = StrEmail{}
	val, err := _t.Translate("test@test.com")
	sbtest.Nil(t, err)
	sbtest.Eq(t, "test@test.com", val)

	_, err = _t.Translate("asdf")
	sbtest.NotNil(t, err)
}
