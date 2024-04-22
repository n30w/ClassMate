package dal

import (
	"encoding/csv"
	"github.com/n30w/Darkspace/internal/models"
	"os"
)

// CSV defines access operations for accessing data from a CSV file.
// This exists because we currently do not have a functioning database just yet.
// General overview of CSV handling in Go: https://earthly.dev/blog/golang-csv-files/

type CSVStore struct {
	path string
}

func NewCSVStore(p string) *CSVStore {
	return &CSVStore{path: p}
}

// readCSV reads a CSV file at the specified path.
// It returns a multidimensional array of strings.
func (cs *CSVStore) readCSV() ([][]string, error) {
	f, err := os.Open(cs.path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	r := csv.NewReader(f)
	data, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

// writeCSV creates a new CSV file.
// This can be used in tandem with readCSV to read a CSV,
// delete a row from the slices,
// then write the new slices to a new CSV file that overwrites the original.
// Here is a helpful article on the writing to a CSV pattern: https
// ://gosamples.dev/write-csv/
func (cs *CSVStore) writeCSV(data [][]string) error {
	f, err := os.Create(cs.path)
	if err != nil {
		return err
	}

	defer f.Close()

	writer := csv.NewWriter(f)

	defer writer.Flush()

	err = writer.WriteAll(data)
	if err != nil {
		return err
	}

	return nil
}

// updateCSV appends a line to a CSV file.
func (cs *CSVStore) updateCSV(row []string) error {

	f, err := os.OpenFile(cs.path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()

	err = writer.Write(row)
	if err != nil {
		return err
	}

	return nil
}

func (cs *CSVStore) InsertCourse(c *models.Course) error {
	//TODO implement me
	panic("implement me")
}

func (cs *CSVStore) GetCourseByName(name string) (*models.Course, error) {
	//TODO implement me
	panic("implement me")
}

func (cs *CSVStore) GetCourseByID(id string) (*models.Course, error) {
	//TODO implement me
	panic("implement me")
}

func (cs *CSVStore) GetRoster(id string) ([]models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (cs *CSVStore) InsertUser(u *models.User) error {
	//TODO implement me
	panic("implement me")
}

func (cs *CSVStore) GetUserByID(id string) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (cs *CSVStore) GetUserByEmail(email string) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (cs *CSVStore) GetByUsername(username string) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}
