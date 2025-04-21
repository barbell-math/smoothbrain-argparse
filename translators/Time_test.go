package translators

import (
	"testing"
	stdtime "time"

	sbtest "github.com/barbell-math/smoothbrain-test"
)

func TestTimeTranslator(t *testing.T) {
	var _t Translator[stdtime.Time]

	_t = NewTime(stdtime.DateOnly)
	val, err := _t.Translate("2000-03-05")
	sbtest.Nil(t, err)
	sbtest.True(
		t, val.Equal(stdtime.Date(2000, 3, 5, 0, 0, 0, 0, val.Location())),
	)

	_, err = _t.Translate("asdf")
	sbtest.NotNil(t, err)
}

func TestDurationTranslator(t *testing.T) {
	var _t Translator[stdtime.Duration]

	_t = Duration{}
	val, err := _t.Translate("2h3m4s")
	sbtest.Nil(t, err)
	sbtest.Eq(t, 2*3600+3*60+4, val.Seconds())

	_, err = _t.Translate("asdf")
	sbtest.NotNil(t, err)
}
