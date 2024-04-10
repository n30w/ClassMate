package dal

import (
	"database/sql"
	"errors"

	"github.com/n30w/Darkspace/internal/models"
)

// Store implements interfaces found in respective domain packages.
type Store struct {
	db *sql.DB
}

var err error

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) InsertUser(u *models.User) error {
	stmt, err := s.db.Prepare("INSERT INTO users (username, password, email, created_at) VALUES ($1, $2, $3, $4) RETURNING id, created_at")
	if err != nil {return err}
	defer stmt.Close()
	row := stmt.QueryRow(u.Username, u.Password, u.Email, u.CreatedAt)
	if err := row.Scan(&u.ID, &u.CreatedAt); err != nil {
		return err
	}
    return nil
}

func (s *Store) GetUserByID(userid string) (*models.User, error) {
	u := &models.User{}

	query := `SELECT id, email, full_name FROM users WHERE id = $1`
	row := s.db.QueryRow(query, userid)
	if err := row.Scan(&u.ID, &u.Email, &u.FullName); err != nil {
		return u, err
	}

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
	row := s.db.QueryRow(query, email)
	if err := row.Scan(&u.ID, &u.Email, &u.FullName); err != nil {
		return u, err
	}

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

func (s *Store) GetUserByUsername(username string) (*models.User, error) {

	u := &models.User{}

	query := `SELECT id, email, full_name FROM users WHERE username = $1`
	row := s.db.QueryRow(query, username)
	if err := row.Scan(&u.ID, &u.Email, &u.FullName); err != nil {
		return u, err
	}

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

func (s *Store) InsertCourse(c *models.Course) error {
	return nil
}

func (s *Store) GetCourseByName(name string) (
	*models.Course,
	error,
) {
	return nil, nil
}

func (s *Store) GetCourseByID(courseid string) (
	*models.Course,
	error,
) {
	return nil, nil
}

func (s *Store) GetRoster(courseid string) (
	[]models.User,
	error,
) {
	return nil, nil
}

func (s *Store) DeleteUserFromCourse(c *models.Course) error { return nil }

func (s *Store) DeleteCourseFromUser(
	courseid string,
	u *models.User,
) error {
	return nil
}

func (s *Store) ChangeCourseName(c *models.Course, name string) error {
	return nil
}

func (s *Store) AddStudent(c *models.Course, userid string) error {
	return nil
}
func (s *Store) RemoveStudent(c *models.Course, userid string) error {
	return nil
}

func (s *Store) InsertMessage(m *models.Message) (
	*models.Message,
	error,
) {
	return nil, nil
}

func (s *Store) GetMessageById(id string) (
	*models.Message,
	error,
) {
	return nil, nil
}

func (s *Store) DeleteMessage(id string) error { return nil }

func (s *Store) EditMessage(id string) (
	*models.Message,
	error,
) {
	return nil, nil
}
