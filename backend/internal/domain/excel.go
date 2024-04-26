package domain

import (
	"github.com/google/uuid"
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
	defer func() {
        if err := f.Close(); err != nil {
            return nil, err
        }
    }()
	course, err := cs.store.GetCourseByID(courseid)
	if err != nil {
		return nil, err
	}
	for _, courseid := range course.Assignments {
		f.NewSheet(id)
		assignment := cs.store.GetAssignmentById(id)
		headers := []string{"NetID", "#SID", "Numeric Grade", "Feedback"}
		for i, header := range headers {
				file.SetCellValue(courseid, fmt.Sprintf("%s%d", string(rune(65+i)), 1), header)
			}
		for index, userid := range course.Roster { // Add user in column A
			err = f.SetCellValue(courseid, fmt.Sprintf("%s%d", string(rune(65)),index), userid)
			err = f.SetCellValue(courseid, fmt.Sprintf("%s%d", string(rune(66)),index), assignment.Submission[index])
			err = submission, _ := cs.store.GetSubmissionById(assignment.Submission[index])
			err = f.SetCellValue(courseid, fmt.Sprintf("%s%d", string(rune(67)),index), submission.Grade)
			err = f.SetCellValue(courseid, fmt.Sprintf("%s%d", string(rune(68)),index), submission.Feedback)
		}
	}
	if err := file.SaveAs(fmt.Sprintf("%s.xlsx"),course.Title); err != nil {
       return nil, err
    }
	return file, nil
}


func (cs *ExcelService) DownloadExcel()