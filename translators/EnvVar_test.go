package translators

import (
	"os"
	"testing"

	sbtest "github.com/barbell-math/smoothbrain-test"
)

func TestEnvVarTranslator(t *testing.T) {
	var _t Translator[string]

	_t = EnvVar{}
	_, err := _t.Translate("__TEST_ENV_VAR")
	sbtest.ContainsError(t, EnvVarNotSetErr, err)

	sbtest.Nil(t, os.Setenv("__TEST_ENV_VAR", "secret"))
	defer func() { sbtest.Nil(t, os.Unsetenv("__TEST_ENV_VAR")) }()
	val, err := _t.Translate("__TEST_ENV_VAR")
	sbtest.Nil(t, err)
	sbtest.Eq(t, "secret", val)
}
