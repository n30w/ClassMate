package domain

import (
	"github.com/n30w/Darkspace/internal/models"
)

// announcement and discussion services
type MediaStore interface {
	GetMediaById(id string) (*models.Media, error)
	InsertMedia(media *models.Media) (*models.Media, error)
	InsertMediaIntoCourse(m *models.Media) error
	InsertMediaIntoAssignment(m *models.Media) error
	InsertMediaIntoSubmission(m *models.Media) error
	InsertMediaIntoCourseBanner(m *models.Media) error
}

type MediaService struct {
	store MediaStore
}

func NewMediaService(m MediaStore) *MediaService { return &MediaService{store: m} }

func (ms *MediaService) AddBanner(
	media *models.Media,
) (*models.Media, error) {
	media, err := ms.store.InsertMedia(media)
	if err != nil {
		return nil, err
	}

	err = ms.store.InsertMediaIntoCourse(media)
	if err != nil {
		return nil, err
	}
	err = ms.store.InsertMediaIntoCourseBanner(media)
	if err != nil {
		return nil, err
	}

	return media, nil
}

func (ms *MediaService) AddAssignmentMedia(
	media *models.Media,
) (*models.Media, error) {
	media, err := ms.store.InsertMedia(media)
	if err != nil {
		return nil, err
	}
	err = ms.store.InsertMediaIntoAssignment(media)
	if err != nil {
		return nil, err
	}
	return media, nil
}
func (ms *MediaService) AddSubmissionMedia(
	media *models.Media,
) (*models.Media, error) {
	media, err := ms.store.InsertMedia(media)
	if err != nil {
		return nil, err
	}
	err = ms.store.InsertMediaIntoSubmission(media)
	if err != nil {
		return nil, err
	}
	return media, nil
}

// GetMedia retrieves a piece of media from a file system given a path.
// It does two things: finds a piece of media in the database by its
// path and, if it does find it, returns it as a sequence of bytes.
func (ms *MediaService) GetMedia(id string) (*models.Media, error) {
	media, err := ms.store.GetMediaById(id)
	if err != nil {
		return nil, err
	}
	return media, nil
}
