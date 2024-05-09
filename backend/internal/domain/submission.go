package domain

import (
	"github.com/google/uuid"
	"github.com/n30w/Darkspace/internal/models"
)

type SubmissionStore interface {
	GetSubmissions(assignmentId string) ([]models.Submission, error)
	InsertSubmission(sub *models.Submission, file string) (
		*models.Submission,
		error,
	)
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
	s.ID = uuid.New().String()
	_, err := ss.store.InsertSubmission(s, "")
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

// GetSubmissions retrieves the submissions for a specific course given
// a Course ID and Assignment ID. It returns a slice of submissions
// for the given assignment.
func (ss *SubmissionService) GetSubmissions(courseId, assignmentId string) (
	[]models.Submission,
	error,
) {
	// Check if the course even exists.

	// Check if the assignment exists.

	// Get all submissions using query.

	return nil, nil
}

// UpdateSubmissions updates submissions from a slice of
// submissions. This is used for updating submission entries
// in the database from an Excel file.
func (ss *SubmissionService) UpdateSubmissions(
	courseId string,
	submissions []models.Submission,
) error {
	for _, submission := range submissions {
		err := ss.store.UpdateSubmission(&submission)
		if err != nil {
			return err
		}
	}

	return nil
}
