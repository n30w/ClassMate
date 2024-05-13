// storage.go contains any data access layer representations that access
// any type of storage, such as a directory, a volume, or a remote object
// storage service such as Amazon S3.

package dal

import (
	"io"
	"os"
	"path"
)

const darkspaceDirectory = "darkspace_volume"
const defaultsDirectory = "defaults"
const templateDirectory = "templates"

type volume struct {
	path      string // path is the general URI path to the volume.
	defaults  string // defaults is where default resources are stored.
	templates string // templates is where templates are stored.
}

// String prints out the volume's set path.
func (v volume) String() string {
	return v.path
}

// Template returns the path of the templates.
func (v volume) Template() string {
	return path.Join(v.templates)
}

type LocalVolume struct {
	volume
}

func NewLocalVolume(p string) *LocalVolume {
	v := &LocalVolume{
		volume: volume{
			path: path.Join(p, darkspaceDirectory),
		},
	}

	v.defaults = path.Join(v.path, defaultsDirectory)
	v.templates = path.Join(v.path, templateDirectory)

	return v
}

// CreateFile makes a new file and returns it. Does not automatically close!
func (lv *LocalVolume) CreateFile(name string) (*os.File, string, error) {
	p := path.Join(lv.path, name)

	f, err := os.Create(p)
	if err != nil {
		return nil, "", err
	}

	return f, p, nil
}

func (lv *LocalVolume) CopyFile(f1 io.Writer, f2 io.Reader) error {
	_, err := io.Copy(f1, f2)
	if err != nil {
		return err
	}

	return nil
}
