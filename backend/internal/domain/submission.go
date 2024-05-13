package domain

import (
	"fmt"

	"github.com/n30w/Darkspace/internal/models"
)

type SubmissionStore interface {
	GetSubmissions(assignmentId string) ([]*models.Submission, error)
	GetSubmissionById(submissionId string) (*models.Submission, error)
	GetSubmissionMedia(submission *models.Submission) (*models.Submission, error)
	GetSubmissionIdByUserAndAssignment(netId string, assignmentId string) (string, error)
	InsertSubmission(sub *models.Submission) (
		*models.Submission,
		error,
	)
	InsertSubmissionIntoAssignment(sub *models.Submission) (*models.Submission, error)
	InsertSubmissionIntoUser(sub *models.Submission) (*models.Submission, error)
	UpdateSubmission(submission *models.Submission) error
	DeleteSubmissionByID(id string) error
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
	fmt.Printf("Inserting into submissions table...\n")
	// Insert submission into submission table
	s, err := ss.store.InsertSubmission(s)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Inserting into assignment_submissions table...\n")

	// Insert submission into assignment_submissions table
	s, err = ss.store.InsertSubmissionIntoAssignment(s)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Inserting into user_submissions table...\n")

	// Insert submission into user_submissions table
	s, err = ss.store.InsertSubmissionIntoUser(s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (ss *SubmissionService) GradeSubmission(grade int, feedback string, submissionid string) (*models.Submission, error) {
	submission, err := ss.store.GetSubmissionById(submissionid)
	if err != nil {
		return nil, err
	}
	submission.Grade = float64(grade)
	submission.Feedback = feedback

	err = ss.store.UpdateSubmission(submission)
	if err != nil {
		return nil, err
	}
	return submission, nil
}

func (ss *SubmissionService) DeleteSubmission(id string) error {
	err := ss.store.DeleteSubmissionByID(id)
	if err != nil {
		return err
	}
	return nil
}

func (ss *SubmissionService) GetSubmission(id string) (
	*models.Submission,
	error,
) {
	submission, err := ss.store.GetSubmissionById(id)
	if err != nil {
		return nil, err
	}
	return submission, nil
}

func (ss *SubmissionService) UpdateSubmission(id string) (
	*models.Submission,
	error,
) {

	return nil, nil
}

// GetUserSubmission retrieves the submission by a user for an assignment given
// a netId and assignmentId
func (ss *SubmissionService) GetUserSubmission(userId string, assignmentId string) (
	*models.Submission,
	error,
) {
	submissionId, err := ss.store.GetSubmissionIdByUserAndAssignment(userId, assignmentId)
	if err != nil {
		return nil, err
	}
	submission, err := ss.store.GetSubmissionById(submissionId)
	if err != nil {
		return nil, err
	}
	submission, err = ss.store.GetSubmissionMedia(submission)
	if err != nil {
		return nil, err
	}
	return submission, nil
}

// GetSubmissions retrieves the submissions for a specific course given
// a Course ID and Assignment ID. It returns a slice of submissions
// for the given assignment.
func (ss *SubmissionService) GetSubmissions(assignmentId string) (
	[]*models.Submission,
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
