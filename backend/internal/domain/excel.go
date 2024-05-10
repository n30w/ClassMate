package domain

import (
	"io"
	"path"
	"strconv"
	"strings"

	"github.com/n30w/Darkspace/internal/models"
	"github.com/xuri/excelize/v2"
)

type ExcelStore interface {
	Get(path ...string) ([][]string, error)
	Save(file *excelize.File, to string) (string, error)
	Open(path ...string) (*excelize.File, error)
	AddRow(f *excelize.File, row *[]interface{}, start string) error
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
	var submissions []models.Submission
	rows, err := es.store.Get(path)
	if err != nil {
		return nil, err
	}

	// [1:] to ignore the first row,
	// which is just column headers for data in the
	// Excel template.
	for _, row := range rows[1:] {
		submission := models.Submission{}
		submission.User.ID = row[0]
		submission.Grade, _ = strconv.ParseFloat(row[1], 64)
		submission.Feedback = row[2]
		submission.ID = row[3]
		submissions = append(submissions, submission)
	}

	return submissions, nil
}

// WriteSubmissions writes to an Excel file which will be sent to the
// teacher for their offline grading use. It writes to an Excel file.
// p is the path to save the file to. The name of the file is automatically
// generated. The path to the generated file is returned along with an error.
func (es *ExcelService) WriteSubmissions(
	p, fileName string,
	submissions []models.Submission,
) (string, error) {
	savePath := path.Join(p, fileName)

	// Open template.
	f, err := es.store.Open()
	if err != nil {
		return "", err
	}

	defer f.Close()

	// Write rows to template.
	for i, submission := range submissions {
		row := &[]interface{}{submission.User.FullName, submission.User.ID, submission.Grade, submission.Feedback, submission.ID}

		// Start in column A, increment downward. i+2 because
		// i starts at 0, Excel rows start at 1, and the first
		// row is used by column headers.
		err = es.store.AddRow(f, row, "A"+strconv.Itoa(i+2))
		if err != nil {
			return "", err
		}
	}

	// Write Course ID and Assignment ID to template. Uses the
	// fileName to retrieve the Course ID and Assignment ID.
	caId := strings.Split(fileName, "-")
	row := &[]interface{}{
		caId[0],
		caId[1],
	}
	err = es.store.AddRow(f, row, "G2")

	// Save the file to disk.
	s, err := es.store.Save(f, savePath)
	if err != nil {
		return "", err
	}

	// s should be a complete path with the generated file name.
	return s, nil
}

func (es *ExcelService) Save(f *excelize.File, to string) (string, error) {
	p, err := es.store.Save(f, to)
	if err != nil {
		return "", err
	}

	return p, nil
}

// SendFile writes an Excel file to an io.Writer interface. For our
// use case, this will be an HTTP stream.
func (es *ExcelService) SendFile(path string, w io.Writer) error {
	f, err := es.store.Open(path)
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
