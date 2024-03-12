package dal

import (
	"database/sql"
	"errors"
	"github.com/n30w/Darkspace/internal/models"
)

// Store will implement interfaces found in respective domain packages.
type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) InsertUser(u *models.User) error {

	query := `INSERT INTO users ... VALUES ($1, $2) RETURNING`

	err := s.db.QueryRow(query).Scan(
		&u.ID, &u.Username, &u.Password, &u.Email,
		&u.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetUserByID(id string) (*models.User, error) {

	u := &models.User{}

	query := `SELECT id, email, full_name FROM users WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(u.ID, u.Email, u.FullName)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ERR_RECORD_NOT_FOUND
		default:
			return nil, err
		}
	}

	return u, nil
}

func (s *Store) GetUserByEmail(email string) (*models.User, error) {

	u := &models.User{}

	query := `SELECT id, email, full_name FROM users WHERE email = $1`

	err := s.db.QueryRow(query, email).Scan(u.ID, u.Email, u.FullName)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ERR_RECORD_NOT_FOUND
		default:
			return nil, err
		}
	}

	return u, nil
}

func (s *Store) GetByUsername(username string) (*models.User, error) {

	u := &models.User{}

	query := `SELECT id, email, full_name FROM users WHERE username = $1`

	err := s.db.QueryRow(query, username).Scan(u.ID, u.Email, u.FullName)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ERR_RECORD_NOT_FOUND
		default:
			return nil, err
		}
	}

	return u, nil
}

func (s *Store) InsertCourse(c *models.Course) error { return nil }

func (s *Store) GetCourseByName(name string) (
	*models.Course,
	error,
) {
	return nil, nil
}

func (s *Store) GetCourseByID(id string) (
	*models.Course,
	error,
) {
	return nil, nil
}

func (s *Store) GetRoster(id string) ([]models.User, error) { return nil, nil }
