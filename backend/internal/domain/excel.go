package domain

import (
	"fmt"
	"io"
	"strconv"

	"github.com/n30w/Darkspace/internal/models"
	"github.com/xuri/excelize/v2"
)

type ExcelStore interface {
	GetCourseByID(courseid string) (*models.Course, error)
	GetAssignmentById(assignmentId string) (*models.Assignment, error)
	GetSubmissionById(submissionId string) (*models.Submission, error)
	GradeSubmission(grade float64, submission *models.Submission) error
	InsertSubmissionFeedback(
		feedback string,
		submission *models.Submission,
	) error
	GetSubmissionFeedback(path, sheet string) ([]models.Submission, error)
	OpenFile(path string) (*excelize.File, error)
}

type ExcelService struct {
	store ExcelStore
}

func NewExcelService(e ExcelStore) *ExcelService { return &ExcelService{store: e} }

// ReadSubmissions reads an Excel file from a path. This method is
// to be used when receiving an offline graded submission Excel sheet,
// which is submitted by the teacher. This method reads the
// Excel sheet and returns a slice of Submissions, which can then
// be put into the database.
func (es *ExcelService) ReadSubmissions(path string) (
	[]models.Submission,
	error,
) {
	submissions, err := es.store.GetSubmissionFeedback(path, "Sheet1")
	if err != nil {
		return nil, err
	}

	return submissions, nil
}

// WriteSubmissions writes to an Excel file which will be sent to the
// teacher for their offline grading use.
func (es *ExcelService) WriteSubmissions(
	path string,
	submissions []models.Submission,
) error {

	return nil
}

func (es *ExcelService) SendFile(path string, w io.Writer) error {
	f, err := es.store.OpenFile(path)
	if err != nil {
		return err
	}

	defer f.Close()

	err = f.Write(w)
	if err != nil {
		return nil
	}

	return nil
}

func (es *ExcelService) CreateExcel(courseId string) (*excelize.File, error) {
	f := excelize.NewFile()
	defer f.Close()

	course, err := es.store.GetCourseByID(courseId)
	if err != nil {
		return nil, err
	}
	headers := []string{"NetID", "Numeric Grade", "Feedback", "#SID"}

	for idx, id := range course.Assignments {
		sheetname := fmt.Sprintf("Assignment-%d", rune(idx))
		f.NewSheet(sheetname)
		assignment, err := es.store.GetAssignmentById(id)
		if err != nil {
			return nil, err
		}
		for i, header := range headers {
			f.SetCellValue(
				id,
				fmt.Sprintf("%s%d", string(rune(65+i)), 1),
				header,
			) // Set headers
		}
		for index, userid := range course.Roster {
			submission, err := es.store.GetSubmissionById(assignment.Submission[index])
			if err != nil {
				return nil, err
			}
			err = f.SetCellValue(
				id,
				fmt.Sprintf("%s%d", string(rune(65)), index),
				userid,
			) // Add user in column A
			if err != nil {
				return nil, err
			}
			err = f.SetCellValue(
				id,
				fmt.Sprintf("%s%d", string(rune(66)), index),
				submission.Grade,
			) // Add submission grade in column B
			if err != nil {
				return nil, err
			}
			err = f.SetCellValue(
				id,
				fmt.Sprintf("%s%d", string(rune(67)), index),
				submission.Feedback,
			) // Add submission feedback in column C
			if err != nil {
				return nil, err
			}
			err = f.SetCellValue(
				id,
				fmt.Sprintf("%s%d", string(rune(68)), index),
				assignment.Submission[index],
			) // Add submission id in column D
			if err != nil {
				return nil, err
			}
		}
	}

	if err := f.SaveAs(fmt.Sprintf("%s.xlsx", course.Title)); err != nil {
		return nil, err
	}

	return f, nil
}

func (es *ExcelService) ParseExcel(excel *excelize.File) error {
	for id := 0; id < excel.SheetCount; id++ {
		// Get the name of the sheet
		assignment := excel.GetSheetName(id) // Sheet name in the form of "Assignment-{id}"
		rows, err := excel.GetRows(assignment)
		if err != nil {
			return err
		}
		for _, row := range rows { // Loop through each row (each student)
			sid := row[0]
			submission, err := es.store.GetSubmissionById(sid) // Get submission struct
			if err != nil {
				return err
			}
			gradeStr := row[1]
			gradeFloat, err := strconv.ParseFloat(gradeStr, 64)
			if err != nil {
				return err
			}
			err = es.store.GradeSubmission(
				gradeFloat,
				submission,
			) // Grade the submission
			if err != nil {
				return err
			}
			err = es.store.InsertSubmissionFeedback(
				row[2],
				submission,
			) // Input feedback for submission
			if err != nil {
				return err
			}
		}
	}
	return nil
}
