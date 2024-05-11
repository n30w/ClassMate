package domain

import (
	"io"
	"mime/multipart"
	"os"
)

type FileStore interface {
	CreateFile(path string) (*os.File, string, error)
	CopyFile(f1 io.Writer, f2 io.Reader) error
}

type FileService struct {
	store FileStore
}

func NewFileService(store FileStore) *FileService { return &FileService{store: store} }

// Save saves a file to disk. This is used for incoming
// files from the handlers. It returns a path to where the
// file was saved and an error.
func (fs *FileService) Save(name string, in multipart.File) (string, error) {

	f, p, err := fs.store.CreateFile(name)
	if err != nil {
		return "", err
	}

	defer f.Close()

	err = fs.store.CopyFile(f, in)
	if err != nil {
		return "", err
	}

	return p, nil
}

// GetFile opens a file at the specified path and returns it.
func (fs *FileService) GetFile(path string) (*os.File, error) {
	// Open the file at the specified path
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return file, nil
}
