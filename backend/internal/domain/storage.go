package domain

type StorageStore interface {
}

type StorageService struct {
	store StorageStore
}

// ReadFile reads a file into memory. It returns a slice
// of bytes and an error, if there is one.
func (s *StorageService) ReadFile(path string) ([]byte, error) {
	return nil, nil
}

// WriteFile writes a file to a file path using a slice
// of data bytes[].
func (s *StorageService) WriteFile(path string, data []byte) error {
	return nil
}
