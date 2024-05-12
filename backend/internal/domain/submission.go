package domain

import (
	"github.com/n30w/Darkspace/internal/models"
)

type SubmissionStore interface {
	GetSubmissions(assignmentId string) ([]models.Submission, error)
	InsertSubmission(sub *models.Submission, file string) (
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
) {

	return nil, nil
}

// GetSubmissions retrieves the submissions for a specific course given
// a Course ID and Assignment ID. It returns a slice of submissions
// for the given assignment.
func (ss *SubmissionService) GetSubmissions(assignmentId string) (
	[]models.Submission,
	error,
) {
	// Get all submissions using assignmentId.
	submissions, err := ss.store.GetSubmissions(assignmentId)
	if err != nil {
		return nil, err
	}

	return submissions, nil
}

// UpdateSubmissions updates submissions from a slice of
// submissions. This is used for updating submission entries
// in the database from an Excel file.
func (ss *SubmissionService) UpdateSubmissions(
	submissions []models.Submission,
) error {
	// You can technically do this in one go, but not sure
	// how to write that query...
	for _, submission := range submissions {
		err := ss.store.UpdateSubmission(&submission)
		if err != nil {
			return err
		}
	}

	return nil
}
