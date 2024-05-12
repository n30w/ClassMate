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
	stmt, err := s.db.Prepare(
		`
		INSERT INTO users (net_id, created_at, updated_at,
		username, password, email, membership, full_name)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`,
	)
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
	var p string
	var e string
	var m int
	query := `SELECT net_id, full_name, password, email, membership FROM users WHERE net_id = $1`
	row := s.db.QueryRow(query, u.ID)
	if err := row.Scan(&u.ID, &u.FullName, &p, &e, &m); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ERR_RECORD_NOT_FOUND
		default:
			return nil, err
		}
	}
	var courses []string
	query2 := `SELECT uc.course_id FROM users u JOIN user_courses uc ON u.net_id = uc.user_net_id WHERE u.net_id = $1`
	rows, err := s.db.Query(query2, u.ID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var courseID string
		if err := rows.Scan(&courseID); err != nil {
			return nil, err
		}
		courses = append(courses, courseID)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	u.Courses = courses
	u.Password = password(p)
	u.Email = email(e)
	u.Membership = membership(m)

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
	for i, id := range u.Courses {
		if id == courseid {
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
	for _, courseID := range u.Courses {
		coursequery := `SELECT id, title, banner_id FROM courses WHERE id = $1;`
		row := s.db.QueryRow(coursequery, courseID)

		var course models.Course
		err = row.Scan(&course.ID, &course.Title, &course.Banner)
		if err != nil {
			switch {
			case errors.Is(err, sql.ErrNoRows):
				return nil, ERR_RECORD_NOT_FOUND
			default:
				return nil, err
			}
		}

		teacherquery := `SELECT * FROM course_teachers WHERE course_id=$1`

		rows, err := s.db.Query(teacherquery, courseID)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var teacherIDs []string

		for rows.Next() {
			var teacherID string
			if err := rows.Scan(&teacherID); err != nil {
				return nil, err
			}
			teacherIDs = append(teacherIDs, teacherID)
		}
		if err := rows.Err(); err != nil {
			return nil, err
		}
		course.Teachers = teacherIDs
		courses = append(courses, course)
	}

	return courses, nil
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

func (s *Store) InsertIntoUserCourses(c *models.Course, userid string) error {
	query := `INSERT INTO user_courses (user_net_id, course_id) VALUES ($1, $2);`
	_, err = s.db.Query(query, userid, c.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) InsertTeacherToCourse(c *models.Course, t string) error {
	var id string
	// query := `INSERT INTO user_course (user_net_id, course_id) VALUES ($1, $2) RETURNING id`
	query := `INSERT INTO courses (title, description, created_at, updated_at
		) VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) RETURNING id`
	// args := []interface{}{t, c.ID}
	// err := s.db.QueryRow(query, args).Scan(&id)
	err := s.db.QueryRow(query, c.Title, c.Description).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) CheckCourseProfessorDuplicate(
	courseName string,
	teacherid string,
) (
	bool,
	error,
) {
	var n int
	query := `SELECT COUNT(*) AS course_count
	FROM user_courses uc
	JOIN courses c ON uc.course_id = c.id
	WHERE uc.user_net_id = $1
	AND c.title = $2;`
	row := s.db.QueryRow(query, teacherid, courseName)
	if err := row.Scan(&n); err != nil {
		return false, err
	}
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return false, ERR_RECORD_NOT_FOUND
		default:
			return false, err
		}
	}
	if n > 0 {
		return true, nil
	} else {
		return false, nil
	}

}

func (s *Store) GetMessagesByCourse(courseid string) ([]string, error) {
	query := `SELECT message_id FROM course_messages WHERE course_id = $1`
	rows, err := s.db.Query(query, courseid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messageIds []string

	for rows.Next() {
		var messageId string
		if err := rows.Scan(&messageId); err != nil {
			return nil, err
		}
		messageIds = append(messageIds, messageId)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return messageIds, nil
}

func (s *Store) GetCourseByID(courseid string) (
	*models.Course,
	error,
) {
	c := &models.Course{}
	c.ID = courseid

	query := `SELECT title, description, created_at, banner_id FROM courses WHERE id=$1`
	rows, err := s.db.Query(query, courseid)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.Scan(
			&c.Title,
			&c.Description,
			&c.CreatedAt,
			&c.Banner,
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
	var e string

	query := `SELECT u.net_id, u.email, u.full_name FROM users u
              JOIN course_roster cr ON u.net_id = cr.student_id
              WHERE cr.course_id = $1`

	rows, err := s.db.Query(query, courseid)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &e, &user.FullName); err != nil {
			return nil, err
		}
		user.Email = email(e)
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
	query := `INSERT INTO messages (title, description, type, date) VALUES ($1, 
$2, $3, $4) RETURNING id`
	row := s.db.QueryRow(
		query,
		m.Title,
		m.Description,
		m.Type,
		m.CreatedAt,
	)

	err := row.Scan(&m.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return ERR_RECORD_NOT_FOUND
		}
		return err
	}
	courseQuery := `INSERT INTO course_messages (course_id, message_id)
VALUES ($1, $2)`

	_, err = s.db.Exec(courseQuery, courseid, m.Post.ID)
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

	query := `SELECT id, title, description, type, 
date FROM messages WHERE id = $1`
	row := s.db.QueryRow(query, messageid)

	err := row.Scan(
		&message.Post.ID,
		&message.Title,
		&message.Description,
		&message.Type,
		&message.Date,
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

	_, err := s.db.Exec(query, m.Post.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) ChangeMessageTitle(m *models.Message) (*models.Message, error) {
	query := `UPDATE messages SET title = $1 WHERE id = $2 RETURNING id, title, description, media, date, course, owner`

	row := s.db.QueryRow(query, m.Title, m.Post.ID)

	updatedMessage := &models.Message{}
	err := row.Scan(
		&updatedMessage.Post.ID,
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

	row := s.db.QueryRow(query, m.Description, m.Post.ID)

	updatedMessage := &models.Message{}
	err := row.Scan(
		&updatedMessage.Post.ID,
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
	assignment := models.NewAssignment()

	query := `SELECT id, title, description, due_date FROM assignments WHERE id = $1`
	row := s.db.QueryRow(query, assignmentid)

	err := row.Scan(
		&assignment.ID,
		&assignment.Title,
		&assignment.Description,
		&assignment.DueDate,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ERR_RECORD_NOT_FOUND
		}
		return nil, err
	}

	return assignment, nil
}

func (s *Store) InsertAssignment(a *models.Assignment) (
	*models.Assignment,
	error,
) {
	query := `INSERT INTO assignments (title, description, due_date) VALUES ($1, $2, $3) RETURNING id`

	row := s.db.QueryRow(query, a.Title, a.Description, a.DueDate)
	if err != nil {
		return nil, err
	}
	err = row.Scan(&a.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ERR_RECORD_NOT_FOUND
		}
		return nil, err
	}
	return a, err
}

func (s *Store) InsertIntoCourseAssignments(a *models.Assignment) (
	*models.Assignment,
	error,
) {
	coursequery := `INSERT INTO course_assignments (course_id, assignment_id) VALUES ($1, $2)`
	_, err = s.db.Exec(coursequery, a.Course, a.ID)
	if err != nil {
		return nil, err
	}
	return a, err
}

func (s *Store) InsertAssignmentIntoUser(a *models.Assignment) (
	*models.Assignment,
	error,
) {
	userquery := `INSERT INTO user_assignments (user_net_id, assignment_id) VALUES ($1, $2)`
	_, err = s.db.Exec(userquery, a.Owner, a.ID)
	if err != nil {
		return nil, err
	}
	return a, err
}

func (s *Store) DeleteAssignment(a *models.Assignment) error {
	query := `DELETE FROM assignments WHERE id = $1`

	_, err := s.db.Exec(query, a.ID)
	if err != nil {
		return err
	}

	return nil
}
func (s *Store) GetAssignmentsByCourse(courseid string) ([]string, error) {
	query := `SELECT assignment_id FROM course_assignments WHERE course_id = $1`
	rows, err := s.db.Query(query, courseid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var assignmentIds []string

	for rows.Next() {
		var assignmentId string
		if err := rows.Scan(&assignmentId); err != nil {
			return nil, err
		}
		assignmentIds = append(assignmentIds, assignmentId)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return assignmentIds, nil
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
func (s *Store) InsertToken(t *models.Token) error {
	query := `INSERT INTO tokens (hash, net_id, expiry, scope) VALUES ($1, $2, $3, $4)`
	args := []any{t.Hash, t.NetID, t.Expiry, t.Scope}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := s.db.ExecContext(ctx, query, args...)

	return err
}

func (s *Store) GetTokenFromNetId(t *models.Token) (*models.Token, error) {
	query := `SELECT hash, expiry, scope FROM tokens WHERE hash = $1;`
	row := s.db.QueryRow(query, t.NetID)

	err := row.Scan(
		&t.Hash,
		&t.Expiry,
		&t.Scope,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ERR_RECORD_NOT_FOUND
		}
		return nil, err
	}

	return t, nil

}

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
	query := `INSERT INTO course_roster (course_id, student_id) VALUES ($1, $2)`

	_, err := s.db.Exec(query, c.ID, userid)
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
) (
	*models.Submission,
	error,
) {
	query := `INSERT INTO submissions (submission_time, on_time, grade, feedback) VALUES ($1, $2, $3, $4) RETURNING id`

	row := s.db.QueryRow(
		query,
		sub.SubmissionTime,
		sub.IsOnTime,
		sub.Grade,
		sub.Feedback,
	)
	err := row.Scan(
		&sub.ID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ERR_RECORD_NOT_FOUND
		}
		return nil, err
	}
	return sub, nil
}

func (s *Store) InsertSubmissionIntoAssignment(sub *models.Submission) (*models.Submission, error) {
	query := `INSERT INTO assignment_submissions (assignment_id, submission_id) VALUES ($1, $2)`

	_, err := s.db.Exec(query, sub.AssignmentId, sub.ID)
	if err != nil {
		return nil, err
	}
	return sub, nil
}

func (s *Store) GradeSubmission(
	grade float64,
	submission *models.Submission,
) error {
	return nil
}

func (s *Store) InsertSubmissionFeedback(
	feedback string,
	submission *models.Submission,
) error {
	return nil
}

func (s *Store) GetMembershipById(userid string) (
	*models.Credential,
	error,
) {
	u := &models.User{}

	var m int

	query := `SELECT id, membership FROM users WHERE net_id = $1`
	row := s.db.QueryRow(query, userid)

	err := row.Scan(
		&u.ID,
		&m,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ERR_RECORD_NOT_FOUND
		}
		return nil, err
	}
	u.Membership = membership(m)
	return &u.Membership, nil
}

func (s *Store) GetNetIdFromHash(hash []byte) (
	string,
	error,
) {
	u := &models.User{}
	query := `SELECT net_id FROM tokens WHERE hash = $1`
	row := s.db.QueryRow(query, hash)

	err = row.Scan(
		&u.ID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", ERR_RECORD_NOT_FOUND
		}
		return "", err
	}
	return u.ID, nil
}

func (s *Store) GetNameById(userid string) (
	*models.Credential,
	error,
) {
	u := &models.User{}

	var m int

	query := `SELECT id, membership FROM users WHERE net_id = $1`
	row := s.db.QueryRow(query, userid)

	err := row.Scan(
		&u.ID,
		&m,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ERR_RECORD_NOT_FOUND
		}
		return nil, err
	}
	u.Membership = membership(m)
	return &u.Membership, nil
}

func (s *Store) InsertMedia(
	m *models.Media,
) (
	*models.Media,
	error,
) {
	query := `INSERT INTO media (type, path, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id`

	row := s.db.QueryRow(
		query,
		m.FileType,
		m.FilePath,
		m.CreatedAt,
		m.UpdatedAt,
	)
	err := row.Scan(&m.ID)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (s *Store) GetMediaById(mediaid string) (
	*models.Media,
	error,
) {
	media := &models.Media{}

	query := `SELECT id, type, path FROM media WHERE id = $1`
	row := s.db.QueryRow(query, mediaid)

	err := row.Scan(
		&media.ID,
		&media.FileType,
		&media.FilePath,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ERR_RECORD_NOT_FOUND
		}
		return nil, err
	}

	return media, nil
}

func (s *Store) InsertMediaIntoCourse(
	m *models.Media,
) error {
	query := `INSERT INTO course_media (course_id, media_id, media_path) VALUES ($1, $2, $3)`

	_, err := s.db.Exec(query, m.AttributionsByType["course"], m.ID, m.FilePath)
	if err != nil {
		return err
	}
	return nil
}
func (s *Store) InsertMediaIntoCourseBanner(
	m *models.Media,
) error {

	query := `UPDATE courses SET banner_id = $2 WHERE id = $1;`
	_, err = s.db.Exec(query, m.AttributionsByType["course"], m.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) InsertMediaIntoAssignment(
	m *models.Media,
) error {
	query := `INSERT INTO assignment_media (assignment_id, media_id, media_path) VALUES ($1, $2, $3)`

	_, err := s.db.Exec(
		query,
		m.AttributionsByType["assignment"],
		m.ID,
		m.FilePath,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) InsertMediaIntoSubmission(
	m *models.Media,
) error {
	query := `INSERT INTO submission_media (submission_id, media_id, media_path) VALUES ($1, $2, $3)`

	_, err := s.db.Exec(
		query,
		m.AttributionsByType["submission"],
		m.ID,
		m.FilePath,
	)
	if err != nil {
		return err
	}

	return nil
}
