package translators

import (
	"errors"
	"os"

	sberr "github.com/barbell-math/smoothbrain-errs"
)

type (
	// Represents a cmd line argument that will be treated as an environment
	// variable. The value of the environment variable will be returned as the
	// value for the argument.
	EnvVar struct{}
)

var (
	EnvVarNotSetErr = errors.New(
		"The supplied environment variable was not set",
	)
)

func (_ EnvVar) Translate(arg string) (string, error) {
	rv, ok := os.LookupEnv(arg)
	if !ok {
		return "", sberr.Wrap(
			EnvVarNotSetErr,
			"The env var '%s' does not exist but was expected to", arg,
		)
	}
	return rv, nil
}
func (_ EnvVar) Reset() {}
func (_ EnvVar) HelpAddendum() string {
	return "Name of an environment variable to get the string value of"
}
