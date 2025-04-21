package translators

import (
	"os"
	"testing"

	sbtest "github.com/barbell-math/smoothbrain-test"
)

func TestIsDirTranslator(t *testing.T) {
	var _t Translator[string]

	_t = IsDir{}
	_, err := _t.Translate("/non-existant-dir")
	sbtest.ContainsError(t, os.ErrNotExist, err)

	p, err := _t.Translate(".")
	sbtest.Nil(t, err)
	sbtest.Eq(t, p, ".")
}

func TestIsFileTranslator(t *testing.T) {
	var _t Translator[string]

	_t = IsFile{}
	_, err := _t.Translate("/non-existant-file")
	sbtest.ContainsError(t, os.ErrNotExist, err)

	p, err := _t.Translate("./FileSys_test.go")
	sbtest.Nil(t, err)
	sbtest.Eq(t, p, "./FileSys_test.go")
}

func TestOpenFileTranslator(t *testing.T) {
	var _t Translator[*os.File]

	_t = NewOpenFile().SetFlags(os.O_RDONLY)
	fHandle, err := _t.Translate("/non-existant-file")
	sbtest.ContainsError(t, os.ErrNotExist, err)
	sbtest.Nil(t, fHandle)

	_t = NewOpenFile().SetFlags(os.O_RDONLY)
	fHandle, err = _t.Translate("./FileSys_test.go")
	sbtest.Nil(t, err)
	sbtest.NotNil(t, fHandle)
	fHandle.Close()
}
