package dal

import (
	"database/sql"
	"errors"
	"time"

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

type vEmail string

func (e vEmail) Valid() error {
	return nil
}

func (e vEmail) String() string {
	return string(e)
}

// InsertUser inserts into the database using a user model.
func (s *Store) InsertUser(u *models.User) error {
	stmt, err := s.db.Prepare("INSERT INTO users (net_id, created_at, updated_at, username, password, email, membership, full_name) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id")
	id := 0
	if err != nil {
		return err
	}
	defer stmt.Close()
	row := stmt.QueryRow(
		u.ID,
		u.CreatedAt,
		u.UpdatedAt,
		u.Username,
		u.Password,
		u.Email,
		u.Membership,
		u.FullName,
	)
	if err := row.Scan(&id); err != nil {
		return err
	}
	return nil
}

// GetUserByID retrieve's a user by their Net ID.
func (s *Store) GetUserByID(userid string) (*models.User, error) {

	u := &models.User{}

	query := `SELECT net_id, email, full_name FROM users WHERE net_id=$1`
	rows, err := s.db.Query(query, userid)

	if err != nil {
		return nil, err
	}

	var e, f string

	for rows.Next() {
		if err := rows.Scan(&u.ID, &e, &f); err != nil {
			switch {
			case errors.Is(err, sql.ErrNoRows):
				return nil, ERR_RECORD_NOT_FOUND
			default:
				return nil, err
			}
		}
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	u.Email = vEmail(e)
	u.FullName = f

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

	query := `SELECT net_id, email, full_name FROM users WHERE username=$1`
	rows, err := s.db.Query(query, username)

	if err != nil {
		return nil, err
	}

	var e, f string

	for rows.Next() {
		if err := rows.Scan(&u.ID, &e, &f); err != nil {
			switch {
			case errors.Is(err, sql.ErrNoRows):
				return nil, ERR_RECORD_NOT_FOUND
			default:
				return nil, err
			}
		}
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	u.Email = vEmail(e)
	u.FullName = f

	return u, nil
}

func (s *Store) DeleteCourseFromUser(
	u *models.User,
	courseid models.CourseId,
) error {
	return nil
}

func (s *Store) InsertCourse(c *models.Course) error { return nil }

func (s *Store) GetCourseByName(name string) (
	*models.Course,
	error,
) {
	return nil, nil
}

func (s *Store) GetCourseByID(courseid models.CourseId) (
	*models.Course,
	error,
) {
	return nil, nil
}

func (s *Store) GetRoster(courseid models.CourseId) (
	[]models.User,
	error,
) {
	return nil, nil
}

func (s *Store) DeleteCourse(c *models.Course) error { return nil }

func (s *Store) ChangeCourseName(c *models.Course, name string) error {
	return nil
}

func (s *Store) AddStudent(c *models.Course, userid string) (
	*models.Course,
	error,
) {
	return nil, nil
}
func (s *Store) RemoveStudent(c *models.Course, userid string) (
	*models.Course,
	error,
) {
	return nil, nil
}

func (s *Store) InsertMessage(
	m *models.Message,
	courseid models.CourseId,
) error {
	return nil
}
func (s *Store) GetMessageById(messageid models.MessageId) (
	*models.Message,
	error,
) {
	return nil, nil
}
func (s *Store) DeleteMessage(m *models.Message) error {
	return nil
}
func (s *Store) ChangeMessageTitle(m *models.Message) (*models.Message, error) {
	return nil, nil
}
func (s *Store) ChangeMessageBody(m *models.Message) (*models.Message, error) {
	return nil, nil
}
func (s *Store) GetAssignmentById(assignmentid models.AssignmentId) (
	*models.Assignment,
	error,
) {
	return nil, nil
}
func (s *Store) InsertAssignment(a *models.Assignment) error {
	return nil
}
func (s *Store) DeleteAssignment(a *models.Assignment) error {
	return nil
}
func (s *Store) ChangeAssignmentTitle(
	assignment *models.Assignment,
	title string,
) (*models.Assignment, error) {
	return nil, nil
}
func (s *Store) ChangeAssignmentBody(
	assignment *models.Assignment,
	body string,
) (*models.Assignment, error) {
	return nil, nil
}

func (s *Store) ChangeAssignmentDueDate(
	assignment *models.Assignment,
	duedate time.Time,
) (*models.Assignment, error) {
	return nil, nil
}
