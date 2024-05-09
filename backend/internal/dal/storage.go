// storage.go contains any data access layer representations that access
// any type of storage, such as a directory, a volume, or a remote object
// storage service such as Amazon S3.

package dal

import (
	"io"
	"os"
)

type LocalVolume struct{}

func NewLocalVolume() *LocalVolume {
	return &LocalVolume{}
}

// CreateFile makes a new file and returns it. Does not automatically close!
func (lv *LocalVolume) CreateFile(path string) (*os.File, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func (lv *LocalVolume) CopyFile(f1 io.Writer, f2 io.Reader) error {
	_, err := io.Copy(f1, f2)
	if err != nil {
		return err
	}

	return nil
}
