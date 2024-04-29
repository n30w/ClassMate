package domain

import (
	"fmt"
	"strconv"

	"github.com/n30w/Darkspace/internal/models"
	"github.com/xuri/excelize/v2"
)

type ExcelStore interface {
	GetCourseByID(courseid string) (*models.Course, error)
	GetAssignmentById(assignmentId string) (*models.Assignment, error)
	GetSubmissionById(submissionId string) (*models.Submission, error)
	GradeSubmission(grade float64, submission *models.Submission) error
	InsertSubmissionFeedback(feedback string, submission *models.Submission) error
}

type ExcelService struct {
	store ExcelStore
}

func NewExcelService(e ExcelStore) *ExcelService { return &ExcelService{store: e} }

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
			err = f.SetCellValue(
				id,
				fmt.Sprintf("%s%d", string(rune(66)), index),
				submission.Grade,
			) // Add submission grade in column B
			err = f.SetCellValue(
				id,
				fmt.Sprintf("%s%d", string(rune(67)), index),
				submission.Feedback,
			) // Add submission feedback in column C
			err = f.SetCellValue(
				id,
				fmt.Sprintf("%s%d", string(rune(68)), index),
				assignment.Submission[index],
			) // Add submission id in column D
		}
	}

	if err := f.SaveAs(fmt.Sprintf("%s.xlsx", course.Title)); err != nil {
		return nil, err
	}

	return f, nil
}

func (cs *ExcelService) ParseExcel(excel *excelize.File) error {
	for id := 0; id < excel.SheetCount; id++ {
		// Get the name of the sheet
		assignment := excel.GetSheetName(id) // Sheet name in the form of "Assignment-{id}"
		rows, err := excel.GetRows(assignment)
		if err != nil {
			return err
		}
		for _, row := range rows { // Loop through each row (each student)
			sid := row[0]
			submission, err := cs.store.GetSubmissionById(sid) // Get submission struct
			if err != nil {
				return err
			}
			gradeStr := row[1]
			gradeFloat, err := strconv.ParseFloat(gradeStr, 64)
			err = cs.store.GradeSubmission(gradeFloat, submission) // Grade the submission
			if err != nil {
				return err
			}
			err = cs.store.InsertSubmissionFeedback(row[2], submission) // Input feedback for submission
			if err != nil {
				return err
			}
		}
	}
	return nil
}
