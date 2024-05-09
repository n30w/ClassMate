package domain

import (
	"io"
	"mime/multipart"
	"os"
)

type FileStore interface {
	CreateFile(path string) (*os.File, error)
	CopyFile(f1 io.Writer, f2 io.Reader) error
}

type FileService struct {
	store FileStore
}

func NewFileService(store FileStore) *FileService { return &FileService{store: store} }

// Save saves a file to disk. This is used for incoming
// files from the handlers.
func (fs *FileService) Save(name string, in multipart.File) error {
	f, err := fs.store.CreateFile(name)
	if err != nil {
		return err
	}

	defer f.Close()

	err = fs.store.CopyFile(f, in)
	if err != nil {
		return err
	}

	return nil
}
