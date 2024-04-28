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
