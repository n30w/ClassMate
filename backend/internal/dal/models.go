package dal

import (
	"database/sql"
	"errors"
	"github.com/n30w/Darkspace/internal/models"

	"github.com/lib/pq"
)

var (
	ERR_RECORD_NOT_FOUND = errors.New("record not found")
	ERR_INVALID_BY       = errors.New("invalid get type received")
)

type Models struct {
	Course *CourseModel
}

func NewModels(db *sql.DB) *Models {
	return &Models{
		Course: &CourseModel{db},
	}
}

type CourseModel struct{ db *sql.DB }

func (m CourseModel) Insert(c *models.Course) error {
	return nil
}

func (m CourseModel) Get(by string) (*models.Course, error) {

	if by == "" {
		return nil, ERR_INVALID_BY
	}

	var course models.Course

	q := ``

	// args is the type of stuff we're scanning for.
	args := []any{}

	err := m.db.QueryRow(q, args).Scan(
		&course.Name,
		&course.ID,
		pq.Array(&course.Teachers),
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ERR_RECORD_NOT_FOUND
		default:
			return nil, err
		}
	}

	return &course, nil
}

func (m CourseModel) Update(c *models.Course) error {
	return nil
}

func (m CourseModel) Delete(ID string) error {
	return nil
}
