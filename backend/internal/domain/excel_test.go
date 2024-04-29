package domain

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/n30w/Darkspace/internal/models"
)

func TestCreateandParseExcel(t *testing.T) {
	mock := newMockExcelStore()
	es := NewExcelService(mock)
	SetDatabase(2, 4, 4, mock)
	file, err := es.CreateExcel(mock.Course.ID) // Create Excel
	if err != nil {
		t.Errorf("%v", err)
	}
	// Add grade and feedback
	for id := 0; id < file.SheetCount; id++ {
		// Get the name of the sheet
		assignment := file.GetSheetName(id) // Sheet name in the form of "Assignment-{id}"
		rows, err := file.GetRows(assignment)
		if err != nil {
			t.Errorf("%v", err)
		}
		for rowidx, _ := range rows { // Loop through each row (each student)
			file.SetCellValue(assignment, fmt.Sprintf("%s%d", string(rune(66)), rowidx), 0+rowidx)
			file.SetCellValue(assignment, fmt.Sprintf("%s%d", string(rune(67)), rowidx), fmt.Sprintf("Nice Work Student %d", rowidx))
		}
		err = es.ParseExcel(file) // Parse Excel
		if err != nil {
			t.Errorf("%v", err)
		}
		// Check database for grade and feedback
	}
	file, err = es.CreateExcel(mock.Course.ID) // Create Excel
	for id := 0; id < file.SheetCount; id++ {
		// Get the name of the sheet
		assignment := file.GetSheetName(id) // Sheet name in the form of "Assignment-{id}"
		rows, err := file.GetRows(assignment)
		if err != nil {
			t.Errorf("%v", err)
		}
		for rowidx, _ := range rows { // Loop through each row (each student)
			grade, err := file.GetCellValue(assignment, fmt.Sprintf("%s%d", string(rune(66)), rowidx))
			if err != nil {
				t.Errorf("%v", err)
			}
			want := 0 + rowidx
			if grade != fmt.Sprintf("%d", want) {
				t.Errorf("got %s, want %s", grade, want)
			}
			feedback, err := file.GetCellValue(assignment, fmt.Sprintf("%s%d", string(rune(67)), rowidx))
			if err != nil {
				t.Errorf("%v", err)
			}
			wantStr := fmt.Sprintf("Nice Work Student %d", rowidx)
			if grade != fmt.Sprintf("%d", wantStr) {
				t.Errorf("got %s, want %s", feedback, wantStr)
			}
		}

	}
}

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

	return fmt.Sprintf("%s", randomNumber)
}

func SetDatabase(assignment int, submission int, students int, mus *mockExcelStore) {

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
			mus.Assignments[j].Submission = append(mus.Assignments[j].Submission, sub.ID)
		}
	}
}

type mockExcelStore struct {
	Course      *models.Course
	Students    []*models.User
	Assignments []*models.Assignment
	Submissions []*models.Submission
}

func (mus *mockExcelStore) GetCourseByID(courseid string) (*models.Course, error) {
	return mus.Course, nil
}
func (mus *mockExcelStore) GetAssignmentById(assignmentId string) (*models.Assignment, error) {
	for _, assignment := range mus.Assignments {
		if assignment.ID == assignmentId {
			return assignment, nil
		}
	}
	return nil, fmt.Errorf("no such assignment id")
}

func (mus *mockExcelStore) GetSubmissionById(submissionId string) (*models.Submission, error) {
	for _, submission := range mus.Submissions {
		if submission.ID == submissionId {
			return submission, nil
		}
	}
	return nil, fmt.Errorf("no such assignment id")
}

func (mus *mockExcelStore) GradeSubmission(grade float64, submission *models.Submission) error {
	submission.Grade = grade
	for idx, sub := range mus.Submissions {
		if sub.ID == submission.ID {
			mus.Submissions[idx] = submission
			return nil
		}
	}

	return fmt.Errorf("No such submission")
}

func (mus *mockExcelStore) InsertSubmissionFeedback(feedback string, submission *models.Submission) error {
	submission.Feedback = feedback
	for idx, sub := range mus.Submissions {
		if sub.ID == submission.ID {
			mus.Submissions[idx] = submission
			return nil
		}
	}

	return fmt.Errorf("No such submission")
}
