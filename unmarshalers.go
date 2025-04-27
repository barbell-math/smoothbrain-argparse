package sbargp

import (
	"errors"
	"os"
	"path/filepath"

	sberr "github.com/barbell-math/smoothbrain-errs"
)

type (
	// Can be used to represent a path to a file in a config
	File string
	// Can be used to represent an absolute path to a file in a config. When
	// unmarshaling the path will be made absolute if it is not already.
	AbsFile string
	// Can be used to represent a path to a dir in a config
	Dir string
	// Can be used to represent an absolute path to a dir in a config. When
	// unmarshaling the path will be made absolute if it is not already.
	AbsDir string
	// Can be used to represent a path to an environment variable in a config.
	// The value of the config entry will be set to the value of the environment
	// variable, if it exists. If it does not exist an error is returned.
	EnvVar string
)

var (
	EnvVarNotSetErr = errors.New(
		"The supplied environment variable was not set",
	)
)

// Checks that the supplied data is a valid path to a file and if so sets the
// underlying value to the path.
func (e *EnvVar) UnmarshalText(data []byte) error {
	envVarValue, ok := os.LookupEnv(string(data))
	if !ok {
		return sberr.Wrap(
			EnvVarNotSetErr,
			"The env var '%s' does not exist but was expected to",
			string(data),
		)
	}
	*e = EnvVar(envVarValue)
	return nil
}

// Checks that the supplied data is a valid path to a dir and if so sets the
// underlying value to the path.
func (d *Dir) UnmarshalText(data []byte) error {
	info, err := os.Stat(string(data))
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return sberr.Wrap(os.ErrNotExist, string(data))
	}
	*d = Dir(data)
	return nil
}

// Checks that the supplied data is a valid path to a dir and if so resolves the
// absolute path and sets the underlying value absolute to the path.
func (a *AbsDir) UnmarshalText(data []byte) error {
	info, err := os.Stat(string(data))
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return sberr.Wrap(os.ErrNotExist, string(data))
	}
	abs, err := filepath.Abs(string(data))
	if err != nil {
		return err
	}
	*a = AbsDir(abs)
	return nil
}

// Checks that the supplied data is a valid environment variable and if so sets
// the underlying value to the value of the environment variable.
func (f *File) UnmarshalText(data []byte) error {
	info, err := os.Stat(string(data))
	if err != nil {
		return err
	}
	if !info.Mode().IsRegular() {
		return sberr.Wrap(os.ErrNotExist, string(data))
	}
	*f = File(data)
	return nil
}

// Checks that the supplied data is a valid path to a file and if so resolves
// the absolute path and sets the underlying value absolute to the path.
func (a *AbsFile) UnmarshalText(data []byte) error {
	info, err := os.Stat(string(data))
	if err != nil {
		return err
	}
	if !info.Mode().IsRegular() {
		return sberr.Wrap(os.ErrNotExist, string(data))
	}
	abs, err := filepath.Abs(string(data))
	if err != nil {
		return err
	}
	*a = AbsFile(abs)
	return nil
}
