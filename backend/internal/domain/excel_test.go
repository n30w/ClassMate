package domain

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/n30w/Darkspace/internal/models"
)

// ========= //
//   MOCKS   //
// ========= //

func newMockExcelStore() *mockExcelStore {
	return &mockExcelStore{
		Students:    make([]*models.User, 0),
		Assignments: make([]*models.Assignment, 0),
		Submissions: make([]*models.Submission, 0),
	}
}
func ID() string {

	// Generate a random number of length 3
	randomNumber := rand.Intn(900) + 100

	return fmt.Sprintf("%d", randomNumber)
}
func Print(mus *mockExcelStore, t *testing.T) {
	for idx, student := range mus.Students {
		t.Logf("Student %d: id(%s)", idx, student.ID)
	}
	for _, submission := range mus.Submissions {
		t.Logf(
			"Submission from User %s, Grade: %f, Feedback: %s",
			submission.User.ID,
			submission.Grade,
			submission.Feedback,
		)
	}
}

func SetDatabase(
	assignment int,
	submission int,
	students int,
	mus *mockExcelStore,
) {

	roster := make([]string, 0)
	course := &models.Course{
		Entity: models.Entity{
			ID: ID(),
		},
	}
	mus.Course = course
	for i := 0; i < students; i++ {
		student := &models.User{
			Entity: models.Entity{
				ID: ID(),
			},
		}
		roster = append(roster, student.ID)
		mus.Students = append(mus.Students, student)
	}
	course.Roster = roster

	for i := 0; i < assignment; i++ {
		assignment := &models.Assignment{
			Post: models.Post{
				Entity: models.Entity{
					ID: ID(),
				},
			}}
		mus.Assignments = append(mus.Assignments, assignment)
	}
	for j := 0; j < assignment; j++ {
		for i := 0; i < submission; i++ {
			sub := &models.Submission{
				Entity: models.Entity{
					ID: ID(),
				},
				User: *mus.Students[i],
			}
			mus.Submissions = append(mus.Submissions, sub)
			mus.Assignments[j].Submission = append(
				mus.Assignments[j].Submission,
				sub.ID,
			)
		}
	}
}

type mockExcelStore struct {
	Course      *models.Course
	Students    []*models.User
	Assignments []*models.Assignment
	Submissions []*models.Submission
}

func (mus *mockExcelStore) GetCourseByID(courseid string) (
	*models.Course,
	error,
) {
	return mus.Course, nil
}
func (mus *mockExcelStore) GetAssignmentById(assignmentId string) (
	*models.Assignment,
	error,
) {
	for _, assignment := range mus.Assignments {
		if assignment.ID == assignmentId {
			return assignment, nil
		}
	}
	return nil, fmt.Errorf("no such assignment id")
}

func (mus *mockExcelStore) GetSubmissionById(submissionId string) (
	*models.Submission,
	error,
) {
	for _, submission := range mus.Submissions {
		if submission.ID == submissionId {
			return submission, nil
		}
	}
	return nil, fmt.Errorf("no such assignment id")
}

func (mus *mockExcelStore) GradeSubmission(
	grade float64,
	submission *models.Submission,
) error {
	submission.Grade = grade
	for idx, sub := range mus.Submissions {
		if sub.ID == submission.ID {
			mus.Submissions[idx] = submission
			return nil
		}
	}

	return fmt.Errorf("No such submission")
}

func (mus *mockExcelStore) InsertSubmissionFeedback(
	feedback string,
	submission *models.Submission,
) error {
	submission.Feedback = feedback
	for idx, sub := range mus.Submissions {
		if sub.ID == submission.ID {
			mus.Submissions[idx] = submission
			return nil
		}
	}

	return fmt.Errorf("No such submission")
}
