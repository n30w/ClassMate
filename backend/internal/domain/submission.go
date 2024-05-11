package domain

import (
	"github.com/n30w/Darkspace/internal/models"
)

type SubmissionStore interface {
	InsertSubmission(sub *models.Submission) (
		*models.Submission,
		error,
	)
	InsertSubmissionIntoAssignment(sub *models.Submission) (*models.Submission, error)
	UpdateSubmission(submission *models.Submission) error
}

type SubmissionService struct {
	store SubmissionStore
}

func NewSubmissionService(s SubmissionStore) *SubmissionService {
	return &SubmissionService{store: s}
}

func (ss *SubmissionService) CreateSubmission(s *models.Submission) (
	*models.Submission,
	error,
) {
	// Insert submission into submission table
	s, err := ss.store.InsertSubmission(s)
	if err != nil {
		return nil, err
	}
	// Insert submission into submission_assignment table
	s, err = ss.store.InsertSubmissionIntoAssignment(s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (ss *SubmissionService) DeleteSubmission(id string) (
	*models.Submission,
	error,
) {
	return nil, nil
}

func (ss *SubmissionService) ReadSubmission(id string) (
	*models.Submission,
	error,
) {
	return nil, nil
}

func (ss *SubmissionService) UpdateSubmission(id string) (
	*models.Submission,
	error,
) { // check if there already exists a submission from the user
	return nil, nil
}
