package domain

import (
	"mime/multipart"

	"github.com/n30w/Darkspace/internal/models"
)

// announcement and discussion services
type MediaStore interface {
	GetMediaReferenceById(media *models.Media) error
	UploadMedia(multipart.File, *models.Submission)
	InsertMediaReference(media *models.Media) error
}

type MediaService struct {
	store MediaStore
}

func NewMediaService(m MediaStore) *MediaService { return &MediaService{store: m} }

func (ms *MediaService) UploadMedia(
	multipart.File,
	*models.Submission,
) (*models.Media, error) {
	return nil, nil
}

// GetMedia retrieves a piece of media from a file system given a reference.
// It does two things: finds a piece of media in the database by its
// reference and, if it does find it, returns it as a sequence of bytes.
func (ms *MediaService) GetMedia(ref string) ([]byte, error) {
	return nil, nil
}
