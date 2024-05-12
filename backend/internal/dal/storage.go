// storage.go contains any data access layer representations that access
// any type of storage, such as a directory, a volume, or a remote object
// storage service such as Amazon S3.

package dal

import (
	"io"
	"os"
	"path"
)

type volume struct {
	path     string // path is the general URI path to the volume.
	defaults string // defaults is where default resources are stored.
}

// String prints out the volume's set path.
func (v volume) String() string {
	return v.path
}

type LocalVolume struct {
	volume
}

func NewLocalVolume(p string) *LocalVolume {
	return &LocalVolume{
		volume: volume{
			path:     p,
			defaults: p + "/default",
		},
	}
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
