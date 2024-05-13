package dal

import (
	"fmt"
	"testing"
)

const (
	excelStorePath  = "../../resources/"
	excelOutputPath = "../../resources/test/"
	excelFileName   = "grade-offline-template.xlsx"
)

// defaultRow is the very top row of the Excel template file.
var defaultRow = []string{"Name", "Net ID", "Grade", "Feedback",
	"Submission ID", "", "Course ID", "Assignment ID"}
var templatePath = excelStorePath + excelFileName

func TestExcelStore_Open(t *testing.T) {
	es := NewExcelStore()

	f, err := es.Open(templatePath)
	if err != nil {
		t.Errorf("%+v", err)
	}

	defer f.Close()
}

func TestExcelStore_Get(t *testing.T) {
	es := NewExcelStore()

	want := [][]string{
		defaultRow,
	}

	got, err := es.Get(excelStorePath + excelFileName)
	if err != nil {
		t.Errorf("%+v", err)
	}

	err = multiDimComp(got, want)
	if err != nil {
		t.Errorf("%+v", err)
	}
}

func TestExcelStore_Save(t *testing.T) {
	es := NewExcelStore()

	f, err := es.Open(templatePath)
	if err != nil {
		t.Errorf("%+v", err)
	}

	fileName := "TestExcelStore_Save.xlsx"

	want := excelOutputPath + fileName
	got, err := es.Save(f, want)
	if err != nil {
		t.Errorf("%+v", err)
	}

	if want != got {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestExcelStore_AddRow(t *testing.T) {
	es := NewExcelStore()

	var got, want [][]string

	fileName := "TestExcelStore_AddRow.xlsx"
	row := []interface{}{"Joe Mama", "jm123", 86.3, "Well done.",
		"018f66c3-265d-7f2b-9b45-2d9606ad1d93"}
	dataSerial := []string{"Joe Mama", "jm123", "86.3", "Well done.", "018f66c3-265d-7f2b-9b45-2d9606ad1d93"}

	want = [][]string{
		defaultRow,
		dataSerial,
	}

	// Access and open the template.
	f, err := es.Open(templatePath)
	if err != nil {
		t.Errorf("%+v", err)
	}

	defer f.Close()

	// Add new rows to the template.
	err = es.AddRow(f, &row, "A2")
	if err != nil {
		t.Errorf("%+v", err)
	}

	// Save the changes to the template at a new destination.
	newSavePath := excelOutputPath + fileName

	p, err := es.Save(f, newSavePath)
	if err != nil {
		t.Errorf("%+v", err)
	}

	// Read the changes back.
	got, err = es.Get(p)
	if err != nil {
		t.Errorf("%+v", err)
	}

	err = multiDimComp(got, want)
	if err != nil {
		t.Errorf("%+v", err)
	}
}

// multiDimComp compares two multidimensional arrays, checking
// their length and each of their values to each other.
func multiDimComp(got, want [][]string) error {
	if len(got) != len(want) {
		return fmt.Errorf("length got %d, want %d", len(got), len(want))
	}

	if len(got[0]) != len(want[0]) {
		return fmt.Errorf("length got %d, want %d", len(got[0]), len(want[0]))
	}

	for i := range want {
		for j := range want[i] {
			if got[i][j] != want[i][j] {
				return fmt.Errorf("got %s, want %s", got[i][j], want[i][j])
			}
		}
	}

	return nil
}
