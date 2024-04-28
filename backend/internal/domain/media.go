package domain

import "github.com/n30w/Darkspace/internal/models"

// announcement and discussion services
type MediaStore interface {
	InsertMediaReference()
	GetMediaReferenceById()
}

type MediaService struct {
	store MediaStore
}

func NewMediaService(m MessageStore) *MediaService { return &MediaService{store: med} }

func (ss *SubmissionService) UpdateSubmission(id string) (*models.Submission, error) { // check if there already exists a submission from the user
	return nil, nil
}
