package domain

import (
	"github.com/google/uuid"
	"github.com/n30w/Darkspace/internal/models"
)

type SubmisisonStore interface {
	InsertSubmission(submission *models.Submission) (*models.Submission, error) 
	UpdateSubmission(submission *models.Submission)
}

type SubmissionService struct {
	store SubmissionStore
}

func NewSubmissionService(s SubmissionService) *SubmissionService {
	return &SubmissionService{store: s}
}

func (ss *SubmissionService) CreateSubmission(s *models.Submission) (*models.Submission, error) {
	s.ID = uuid.New().String()
	submission, err := ss.store.InsertSubmission(s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (ss *SubmissionService) UpdateSubmission(id string, *) (*models.Submission, error){

}
