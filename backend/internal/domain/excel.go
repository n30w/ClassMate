package domain

import (
	"fmt"

	"github.com/n30w/Darkspace/internal/models"
	"github.com/xuri/excelize/v2"
)

type ExcelStore interface {
	GetCourseByID(courseid string) (*models.Course, error)
	GetAssignmentById(assignmentid string) (*models.Assignment, error)
	GetSubmissionById(submissionid string) (*models.Submission, error)
}

type ExcelService struct {
	store ExcelStore
}

func NewExcelService(e ExcelStore) *ExcelService { return &ExcelService{store: e} }

func (es *ExcelService) CreateExcel(courseid string) (*excelize.File, error) {
	f := excelize.NewFile()
	defer f.Close()

	course, err := es.store.GetCourseByID(courseid)
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
			f.SetCellValue(courseid, fmt.Sprintf("%s%d", string(rune(65+i)), 1), header) // Set headers
		}
		for index, userid := range course.Roster {
			err = f.SetCellValue(courseid, fmt.Sprintf("%s%d", string(rune(65)), index), userid)                       // Add user in column A
			err = f.SetCellValue(courseid, fmt.Sprintf("%s%d", string(rune(66)), index), assignment.Submission[index]) // Add submission id in column B
			submission, err := es.store.GetSubmissionById(assignment.Submission[index])
			if err != nil {
				return nil, err
			}
			err = f.SetCellValue(courseid, fmt.Sprintf("%s%d", string(rune(67)), index), submission.Grade)    // Add submission grade in column C
			err = f.SetCellValue(courseid, fmt.Sprintf("%s%d", string(rune(68)), index), submission.Feedback) // Add submission feedback in column D
		}
	}
	if err := f.SaveAs(fmt.Sprintf("%s.xlsx", course.Title)); err != nil {
		return nil, err
	}
	return f, nil
}

func (cs *ExcelService) DownloadExcel()
