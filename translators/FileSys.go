package translators

import (
	"io/fs"
	"os"

	sberr "github.com/barbell-math/smoothbrain-errs"
)

type (
	// A translator that checks that the supplied directory exists.
	IsDir struct{}
	// A translator that checks that the supplied file exists.
	IsFile struct{}

	// A translator that makes the supplied directory along with all necessary
	// parent directories.
	mkdir struct {
		// The permissions used to create all dirs and sub-dirs. See
		// [os.MkdirAll] for reference.
		permissions fs.FileMode
	}
	// A translator that makes the supplied file.
	openFile struct {
		// The flags used to determine the file mode. See [os.RDONLY] and
		// friends.
		flags int
		// The permissions used to open the file with. See [os.OpenFile] for
		// reference.
		permissions fs.FileMode
	}
)

func NewOpenFile() *openFile {
	return &openFile{
		flags:       os.O_RDONLY,
		permissions: 0644,
	}
}

// The flags used to determine the file mode. See [os.RDONLY] and
// friends.
func (o *openFile) SetFlags(v int) *openFile {
	o.flags = v
	return o
}

// The permissions used to open the file with. See [os.OpenFile] for
// reference.
func (o *openFile) SetPermissions(v fs.FileMode) *openFile {
	o.permissions = v
	return o
}

// The flags used to determine the file mode. See [os.RDONLY] and
// friends.
func (o *openFile) GetFlags() int {
	return o.flags
}

// The permissions used to open the file with. See [os.OpenFile] for
// reference.
func (o *openFile) GetPermissions() fs.FileMode {
	return o.permissions
}

// Returns a new mkdir struct initialized with the default values.
func NewMkdir() *mkdir {
	return &mkdir{
		permissions: 0644,
	}
}

// The permissions used to create all dirs and sub-dirs. See
// [os.MkdirAll] for reference.
func (m *mkdir) SetPermissions(v fs.FileMode) *mkdir {
	m.permissions = v
	return m
}

// The permissions used to create all dirs and sub-dirs. See
// [os.MkdirAll] for reference.
func (m *mkdir) GetPermissions() fs.FileMode {
	return m.permissions
}

func (_ IsDir) Translate(arg string) (string, error) {
	info, err := os.Stat(arg)
	if err != nil {
		return "", err
	}
	if !info.IsDir() {
		return "", sberr.Wrap(os.ErrNotExist, arg)
	}
	return arg, nil
}
func (_ IsDir) Reset()               {}
func (_ IsDir) HelpAddendum() string { return "" }

func (m mkdir) Translate(arg string) (string, error) {
	return arg, os.MkdirAll(arg, m.permissions)
}
func (m mkdir) Reset()               {}
func (m mkdir) HelpAddendum() string { return "" }

func (_ IsFile) Translate(arg string) (string, error) {
	info, err := os.Stat(arg)
	if err != nil {
		return "", err
	}
	if !info.Mode().IsRegular() {
		return "", sberr.Wrap(os.ErrNotExist, arg)
	}
	return arg, nil
}
func (_ IsFile) Reset()               {}
func (_ IsFile) HelpAddendum() string { return "" }

func (o openFile) Translate(arg string) (*os.File, error) {
	// If the file is not being created check that it exists.
	if o.flags|os.O_CREATE == 0 {
		info, err := os.Stat(arg)
		if err != nil {
			return nil, err
		}
		if info.Mode().IsRegular() {
			return nil, sberr.Wrap(os.ErrNotExist, arg)
		}
	}
	return os.OpenFile(arg, o.flags, o.permissions)
}
func (o openFile) Reset()               {}
func (o openFile) HelpAddendum() string { return "" }
