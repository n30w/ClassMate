package dal

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/n30w/Darkspace/internal/models"
)

// Credential interface implementations. These implementations may seem
// somewhat redundant, but they are helpful, because it lets us test and
// validate the input once more to verify data integrity across boundaries.

type username string
type password string
type email string
type membership int
type ID string

func (i ID) String() string { return string(i) }
func (i ID) Valid() error   { return nil }

func (u username) String() string { return string(u) }
func (u username) Valid() error   { return nil }

func (p password) String() string { return string(p) }
func (p password) Valid() error   { return nil }

func (e email) String() string { return string(e) }
func (e email) Valid() error   { return nil }

func (m membership) String() string { return fmt.Sprintf("%d", m) }
func (m membership) Valid() error   { return nil }

// Store implements interfaces found in respective domain packages.
type Store struct {
	db *sql.DB
}

func (s *Store) InsertMediaReference(media *models.Media) error {
	return nil
}

func (s *Store) UploadMedia(
	file multipart.File,
	submission *models.Submission,
) {
	//TODO implement me
	panic("implement me")
}

func (s *Store) GetSubmissionById(submissionId string) (
	*models.Submission,
	error,
) {
	//TODO implement me
	panic("implement me")
}

func (s *Store) UpdateSubmission(submission *models.Submission) error {
	//TODO implement me
	panic("implement me")
}

func (s *Store) SubmitAssignment(assignment *models.Assignment) (
	*models.Assignment,
	error,
) {
	//TODO implement me
	panic("implement me")
}

func (s *Store) ChangeAssignment(
	assignment *models.Assignment,
	updatedfield string,
	action string,
) (*models.Assignment, error) {
	//TODO implement me
	panic("implement me")
}

var err error

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// InsertUser inserts into the database using a user model.
func (s *Store) InsertUser(u *models.User) error {
	id := 0
	stmt, err := s.db.Prepare("INSERT INTO users (net_id, created_at, updated_at, username, password, email, membership, full_name) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id")
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
func (s *Store) GetUserByID(u *models.User) (*models.User, error) {
	var e string
	query := `SELECT net_id, password FROM users WHERE net_id = $1`
	row := s.db.QueryRow(query, u.ID)
	if err := row.Scan(&u.ID, &e); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ERR_RECORD_NOT_FOUND
		default:
			return nil, err
		}
	}

	u.Password = password(e)

	return u, nil
}

func (s *Store) GetUserById_2(c models.Credential) (*models.User, error) {
	u := &models.User{}
	var e string

	query := `SELECT net_id, password FROM users WHERE net_id = $1`
	row := s.db.QueryRow(query, c.String())
	if err := row.Scan(&u.ID, &e); err != nil {
		return nil, err
	}

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ERR_RECORD_NOT_FOUND
		default:
			return nil, err
		}
	}

	u.Password = password(e)

	return u, nil
}

// GetUserByEmail retrieves a user using a credential, returning
// a user model and error.
func (s *Store) GetUserByEmail(c models.Credential) (*models.User, error) {
	u := &models.User{}
	var e string
	var f string

	query := `SELECT net_id, email, full_name FROM users WHERE email = $1`
	row := s.db.QueryRow(query, c.String())
	if err := row.Scan(&u.ID, &e, &f); err != nil {
		return nil, err
	}

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ERR_RECORD_NOT_FOUND
		default:
			return nil, err
		}
	}

	u.Email = email(e)
	u.FullName = f

	return u, nil
}

// func (s *Store) GetUserByUsername(username models.Credential) (
// 	*models.User,
// 	error,
// ) {
// 	u := &models.User{}

// 	query := `SELECT net_id, email, full_name FROM users WHERE username=$1`
// 	rows, err := s.db.Query(query, username.String())

// 	if err != nil {
// 		return nil, err
// 	}

// 	var e string

// 	for rows.Next() {
// 		if err := rows.Scan(&u.ID, &e, &u.FullName); err != nil {
// 			switch {
// 			case errors.Is(err, sql.ErrNoRows):
// 				return nil, ERR_RECORD_NOT_FOUND
// 			default:
// 				return nil, err
// 			}
// 		}
// 	}

// 	err = rows.Err()
// 	if err != nil {
// 		return nil, err
// 	}

// 	u.Email = email(e)

// 	return u, nil
// }

func (s *Store) DeleteUserByNetID(netId string) (int64, error) {
	query := `DELETE FROM users WHERE net_id = $1`
	var result sql.Result
	var err error

	result, err = s.db.Exec(query, netId)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return 0, ERR_RECORD_NOT_FOUND
		default:
			return 0, err
		}
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rows, nil
}

func (s *Store) DeleteCourseByID(id string) (int64, error) {
	query := `DELETE FROM courses WHERE id = $1`
	var result sql.Result
	var err error

	result, err = s.db.Exec(query, id)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return 0, ERR_RECORD_NOT_FOUND
		default:
			return 0, err
		}
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rows, nil
}

func (s *Store) DeleteCourseByTitle(title string) (int64, error) {
	query := `DELETE FROM courses WHERE title = $1`
	var result sql.Result
	var err error

	result, err = s.db.Exec(query, title)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return 0, ERR_RECORD_NOT_FOUND
		default:
			return 0, err
		}
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rows, nil
}

func (s *Store) DeleteMediaByID(id string) (int64, error) {
	query := `DELETE FROM media WHERE id = $1`
	var result sql.Result
	var err error

	result, err = s.db.Exec(query, id)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return 0, ERR_RECORD_NOT_FOUND
		default:
			return 0, err
		}
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rows, nil
}

func (s *Store) DeleteAssignmentByID(id string) (int64, error) {
	query := `DELETE FROM assignments WHERE id = $1`
	var result sql.Result
	var err error

	result, err = s.db.Exec(query, id)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return 0, ERR_RECORD_NOT_FOUND
		default:
			return 0, err
		}
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rows, nil
}

func (s *Store) DeleteSubmissionByID(id string) (int64, error) {
	query := `DELETE FROM submissions WHERE id = $1`
	var result sql.Result
	var err error

	result, err = s.db.Exec(query, id)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return 0, ERR_RECORD_NOT_FOUND
		default:
			return 0, err
		}
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rows, nil
}

func (s *Store) DeleteCourseFromUser(
	u *models.User,
	courseid string,
) error {
	var indexToRemove = -1
	for i, course := range u.Courses {
		if course.ID == courseid {
			indexToRemove = i
			break
		}
	}

	if indexToRemove == -1 {
		return errors.New("course not found in user's list")
	}

	u.Courses = append(
		u.Courses[:indexToRemove],
		u.Courses[indexToRemove+1:]...,
	)

	return nil
}

func (s *Store) GetUserCourses(u *models.User) ([]models.Course, error) {
	courses := make([]models.Course, 0)

	query := `
	SELECT
		c.id AS course_id,
		c.title AS course_title,
		t.full_name AS teacher_name,
	FROM users u
	JOIN user_courses uc ON u.net_id = uc.user_net_id
	JOIN courses c ON uc.course_id = c.id
	JOIN course_teachers ct ON c.id = ct.course_id
	JOIN users t ON ct.teacher_id = t.net_id
	WHERE u.net_id = $1;
	`
	rows, err := s.db.Query(query, u.ID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		c := models.Course{}
		t := models.User{}
		if err := rows.Scan(&c.ID, &c.Title, &t.FullName); err != nil {
			switch {
			case errors.Is(err, sql.ErrNoRows):
				return nil, ERR_RECORD_NOT_FOUND
			default:
				return nil, err
			}
		}
		courses = append(courses, c)
	}

	return courses, err
}

// func (s *Store) GetCourseProfessors(u *models.User) ([]models.User, error) {
// 	professors := make([]models.User, 0)
// 	query := `
// 	SELECT c.id, c.title, c.description, c.created_at, c.updated_at
// 	FROM users u
// 	JOIN user_courses uc ON u.net_id = uc.user_net_id
// 	JOIN courses c ON uc.course_id = c.id
// 	WHERE u.net_id = $1`

// 	rows, err := s.db.Query(query, u.ID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	for rows.Next() {
// 		p := models.Course{}
// 		if err := rows.Scan(&c.ID, &c.Title, &c.Description, &c.CreatedAt); err != nil {
// 			switch {
// 			case errors.Is(err, sql.ErrNoRows):
// 				return nil, ERR_RECORD_NOT_FOUND
// 			default:
// 				return nil, err
// 			}
// 		}
// 		courses = append(courses, c)
// 	}

// }

// InsertCourse inserts a course into the database based on a model,
// then returns a string value that is the UUID.
func (s *Store) InsertCourse(c *models.Course) (string, error) {
	query := `INSERT INTO courses (title, description, created_at, updated_at
) VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) RETURNING id`
	var err error
	var id string

	err = s.db.QueryRow(query, c.Title, c.Description).Scan(&id)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return id, ERR_RECORD_NOT_FOUND
		default:
			return "", err
		}
	}

	return id, nil
}

func (s *Store) GetCourseByName(name string) (
	*models.Course,
	error,
) {
	c := &models.Course{}

	query := `SELECT id, description, created_at FROM courses WHERE title=$1`
	rows, err := s.db.Query(query, name)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.Scan(&c.ID, &c.Description, &c.CreatedAt); err != nil {
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

	return c, nil
}

func (s *Store) GetCourseByID(courseid string) (
	*models.Course,
	error,
) {
	c := &models.Course{}

	query := `SELECT title, description, created_at FROM courses WHERE id=$1`
	rows, err := s.db.Query(query, courseid)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.Scan(
			&c.Title,
			&c.Description,
			&c.CreatedAt,
		); err != nil {
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

	return c, nil
}

func (s *Store) GetRoster(courseid string) (
	[]models.User,
	error,
) {
	var roster []models.User

	query := `SELECT users.net_id, users.email, users.full_name FROM users INNER JOIN courses ON courses.id = $1 AND users.net_id = ANY(courses.roster)`

	rows, err := s.db.Query(query, courseid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Email, &user.FullName); err != nil {
			return nil, err
		}
		roster = append(roster, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return roster, nil
}

func (s *Store) DeleteCourse(c *models.Course) error {
	query := `
        DELETE FROM courses
        WHERE id = $1
    `

	_, err := s.db.Exec(query, c.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) ChangeCourseName(c *models.Course, name string) error {
	query := `UPDATE courses SET title = $1 WHERE id = $2`

	_, err := s.db.Exec(query, name, c.ID)

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) InsertMessage(
	m *models.Message,
	courseid string,
) error {
	query := `INSERT INTO messages (id, title, description, media, date, course, owner) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := s.db.Exec(
		query,
		m.ID,
		m.Title,
		m.Description,
		m.Media,
		m.Date,
		courseid,
		m.Owner,
	)
	if err != nil {
		return err
	}

	courseQuery := `UPDATE courses SET messages = array_append(messages, $1) WHERE id = $2`
	_, err = s.db.Exec(courseQuery, m.ID, courseid)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetMessageById(messageid string) (
	*models.Message,
	error,
) {
	message := &models.Message{}

	query := `SELECT id, title, description, media, date, course, owner FROM messages WHERE id = $1`
	row := s.db.QueryRow(query, messageid)

	err := row.Scan(
		&message.ID,
		&message.Title,
		&message.Description,
		&message.Media,
		&message.Course,
		&message.Owner,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ERR_RECORD_NOT_FOUND
		}
		return nil, err
	}

	return message, nil
}

func (s *Store) DeleteMessage(m *models.Message) error {
	query := `
        DELETE FROM messages
        WHERE id = $1
    `

	_, err := s.db.Exec(query, m.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) ChangeMessageTitle(m *models.Message) (*models.Message, error) {
	query := `UPDATE messages SET title = $1 WHERE id = $2 RETURNING id, title, description, media, date, course, owner`

	row := s.db.QueryRow(query, m.Title, m.ID)

	updatedMessage := &models.Message{}
	err := row.Scan(
		&updatedMessage.ID,
		&updatedMessage.Title,
		&updatedMessage.Description,
		&updatedMessage.Media,
		&updatedMessage.Course,
		&updatedMessage.Owner,
	)
	if err != nil {
		return nil, err
	}

	return updatedMessage, nil
}

func (s *Store) ChangeMessageBody(m *models.Message) (*models.Message, error) {
	query := `UPDATE messages SET description = $1 WHERE id = $2 RETURNING id, title, description, media, date, course, owner`

	row := s.db.QueryRow(query, m.Description, m.ID)

	updatedMessage := &models.Message{}
	err := row.Scan(
		&updatedMessage.ID,
		&updatedMessage.Title,
		&updatedMessage.Description,
		&updatedMessage.Media,
		&updatedMessage.Course,
		&updatedMessage.Owner,
	)
	if err != nil {
		return nil, err
	}

	return updatedMessage, nil
}

func (s *Store) GetAssignmentById(assignmentid string) (
	*models.Assignment,
	error,
) {
	assignment := &models.Assignment{}

	query := `SELECT id, title, description, due_date, course_id FROM assignments WHERE id = $1`
	row := s.db.QueryRow(query, assignmentid)

	err := row.Scan(
		&assignment.ID,
		&assignment.Title,
		&assignment.Description,
		&assignment.DueDate,
		&assignment.Course,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ERR_RECORD_NOT_FOUND
		}
		return nil, err
	}

	return assignment, nil
}

func (s *Store) InsertAssignment(a *models.Assignment) error {
	query := `INSERT INTO assignments (id, title, description, due_date, course_id) VALUES ($1, $2, $3, $4, $5)`

	_, err := s.db.Exec(
		query,
		a.ID,
		a.Title,
		a.Description,
		a.DueDate,
		a.Course,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteAssignment(a *models.Assignment) error {
	query := `DELETE FROM assignments WHERE id = $1`

	_, err := s.db.Exec(query, a.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) ChangeAssignmentTitle(
	assignment *models.Assignment,
	title string,
) (*models.Assignment, error) {
	query := `UPDATE assignments SET title = $1 WHERE id = $2 RETURNING id, title, description, due_date, course_id`

	row := s.db.QueryRow(query, title, assignment.ID)

	updatedAssignment := &models.Assignment{}
	err := row.Scan(
		&updatedAssignment.ID,
		&updatedAssignment.Title,
		&updatedAssignment.Description,
		&updatedAssignment.DueDate,
		&updatedAssignment.Course,
	)
	if err != nil {
		return nil, err
	}

	return updatedAssignment, nil
}

func (s *Store) ChangeAssignmentBody(
	assignment *models.Assignment,
	body string,
) (*models.Assignment, error) {
	query := `UPDATE assignments SET description = $1 WHERE id = $2 RETURNING id, title, description, due_date, course_id`

	row := s.db.QueryRow(query, body, assignment.ID)

	updatedAssignment := &models.Assignment{}
	err := row.Scan(
		&updatedAssignment.ID,
		&updatedAssignment.Title,
		&updatedAssignment.Description,
		&updatedAssignment.DueDate,
		&updatedAssignment.Course,
	)
	if err != nil {
		return nil, err
	}

	return updatedAssignment, nil
}

// InsertToken inserts a created token for a user.
func (s *Store) InsertToken(t models.Token) error {
	query := `INSERT INTO tokens (hash, net_id, expiry,
scope) VALUES ($1. $2. $3. $4)`

	args := []any{t.Hash, t.NetID, t.Expiry, t.Scope}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, args...)
	return err
}

//func (s *Store) GetToken()

// DeleteTokenFrom deletes a user's authentication Token using their
// Net ID.
func (s *Store) DeleteTokenFrom(netId, scope string) error {
	query := `DELETE FROM tokens WHERE scope = $1 AND net_id = $2`

	args := []any{scope, netId}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, args...)
	return err
}

// ##################
//  JUNCTION METHODS
// ##################
//
// Junction methods function upon junction tables. They change
// the relationships between database objects.

// AddTeacher adds a teacher to a specified course, using the teacher's
// userId. This method uses junction tables to assign relationships.
func (s *Store) AddTeacher(courseId string, userId string) error {
	// Start a transaction
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %v", err)
	}

	// Check if the course exists
	var exists bool
	err = tx.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM courses WHERE id = $1)",
		courseId,
	).Scan(&exists)

	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error checking course existence: %v", err)
	}

	if !exists {
		tx.Rollback()
		return fmt.Errorf("course with ID %s does not exist", courseId)
	}

	// Check if the teacher exists
	err = tx.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM users WHERE net_id = $1)",
		userId,
	).Scan(&exists)

	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error checking teacher existence: %v", err)
	}

	if !exists {
		tx.Rollback()
		return fmt.Errorf("teacher with ID %s does not exist", userId)
	}

	// Insert the new relationship into the junction table
	_, err = tx.Exec(
		"INSERT INTO course_teachers (course_id, teacher_id) VALUES ($1, $2)",
		courseId,
		userId,
	)

	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error inserting into course_teachers: %v", err)
	}

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return fmt.Errorf("error committing transaction: %v", err)
	}

	return nil
}

func (s *Store) ChangeAssignmentDueDate(
	assignment *models.Assignment,
	duedate time.Time,
) (*models.Assignment, error) {
	return nil, nil
}

func (s *Store) GetMediaReferenceById(media *models.Media) error {
	return nil
}

// AddStudent uses junction tables to insert a new student
// into a course.
func (s *Store) AddStudent(c *models.Course, userid string) (
	*models.Course,
	error,
) {
	query := `UPDATE courses SET roster = array_append(roster, $1) WHERE id = $2`

	_, err := s.db.Exec(query, userid, c.ID)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *Store) RemoveStudent(c *models.Course, userid string) (
	*models.Course,
	error,
) {
	query := `UPDATE courses SET roster = array_remove(roster, $1) WHERE id = $2`

	_, err := s.db.Exec(query, userid, c.ID)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *Store) InsertSubmission(
	sub *models.Submission,
	file string,
) (
	*models.Submission,
	error,
) {
	query := `INSERT INTO submissions (id, file_type, submission_time, on_time, grade, feedback) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := s.db.Exec(query, file, sub.ID)
	if err != nil {
		return nil, err
	}

	return sub, nil
}

func (s *Store) GradeSubmission(grade float64, submission *models.Submission) error {
	return nil
}

func (s *Store) InsertSubmissionFeedback(feedback string, submission *models.Submission) error {
	return nil
}
