package domain

import (
	"fmt"

	"github.com/n30w/Darkspace/internal/models"
	"github.com/xuri/excelize/v2"
)

type ExcelStore interface {
	GetCourseByID(courseid string) (*models.Course, error)
	GetAssignmentById(assignmentId string) (*models.Assignment, error)
	GetSubmissionById(submissionId string) (*models.Submission, error)
	GradeSubmission(grade float64, submission *models.Submission) error
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

	for _, id := range course.Assignments {
		f.NewSheet(id)
		assignment, err := es.store.GetAssignmentById(id)
		if err != nil {
			return nil, err
		}
		headers := []string{"NetID", "#SID", "Numeric Grade", "Feedback"}
		for i, header := range headers {
			f.SetCellValue(
				id,
				fmt.Sprintf("%s%d", string(rune(65+i)), 1),
				header,
			) // Set headers
		}
		for index, userid := range course.Roster {
			err = f.SetCellValue(
				id,
				fmt.Sprintf("%s%d", string(rune(65)), index),
				userid,
			) // Add user in column A
			err = f.SetCellValue(
				id,
				fmt.Sprintf("%s%d", string(rune(66)), index),
				assignment.Submission[index],
			) // Add submission id in column B
			submission, err := es.store.GetSubmissionById(assignment.Submission[index])
			if err != nil {
				return nil, err
			}
			err = f.SetCellValue(
				id,
				fmt.Sprintf("%s%d", string(rune(67)), index),
				submission.Grade,
			) // Add submission grade in column C
			err = f.SetCellValue(
				id,
				fmt.Sprintf("%s%d", string(rune(68)), index),
				submission.Feedback,
			) // Add submission feedback in column D
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
		assignmentid := excel.GetSheetName(id)
		for _, row := range excel.GetRows(assignmentid) {
			sid := row[1]
			submission, err := cs.store.GetSubmissionById(sid)
			if err != nil {
				return err
			}
			err = cs.store.GradeSubmission(row[2], row[3], submission)
		}
	}
	return nil
}
