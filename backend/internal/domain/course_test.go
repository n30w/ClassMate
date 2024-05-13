package domain

import (
	"errors"
	"strconv"
	"testing"

	"github.com/n30w/Darkspace/internal/models"
)

func TestCourseService_CreateCourse(t *testing.T) {
	us := NewCourseService(newMockCourseStore())

	// cred is fake credentials.
	course := &models.Course{
		Title:    "Software Engineering",
		Teachers: make([]string, 1),
	}
	teacherid := "teacherid123"
	course.Teachers = append(course.Teachers, teacherid)

	got, _ := us.CreateCourse(course, teacherid)

	if got != nil {
		t.Errorf("got %s", got)
	}
}

// ========= //
//   MOCKS   //
// ========= //

func newMockCourseStore() *mockCourseStore {
	return &mockCourseStore{
		id:         0,
		byID:       make(map[string]*models.Course),
		byEmail:    make(map[string]int),
		byUsername: make(map[string]int),
	}
}

type mockCourseStore struct {
	id         int
	byID       map[string]*models.User
	byEmail    map[string]int
	byUsername map[string]int
}
func (mcs *mockCourseStore) InsertCourse(c *models.Course) (string, error)
{
	mus.id += 1
	mus.byID[strconv.Itoa(mus.id)] = u
	mus.byEmail[u.Email.String()] = mus.id
	mus.byUsername[u.Username.String()] = mus.id
	return nil
}

func (mcs *mockCourseStore) GetCourseByID(courseid string) (*models.Course, error)
{
	if u, ok := mus.byEmail[c.String()]; !ok {
		return mus.byID[strconv.Itoa(u)], errors.New("email already taken")
	}
	return nil, nil
}

func (mcs *mockCourseStore) GetRoster(c string) ([]models.User, error)
{
	if u, ok := mus.byUsername[username.String()]; !ok {
		return mus.byID[strconv.Itoa(u)],
			errors.New("username already taken")
	}
	return nil, nil
}

func (mcs *mockCourseStore) ChangeCourseName(c *models.Course, name string) error
{
	return nil
}

func (mcs *mockCourseStore) AddStudent(c *models.Course, userid string) (*models.Course, error)

{
	return nil, nil
}
func (mcs *mockCourseStore) AddTeacher(courseId, userId string) error
{
	return nil, nil
}
func (mcs *mockCourseStore) RemoveStudent(c *models.Course, userid string) (*models.Course, error)
{
	return nil, nil
}
func (mcs *mockCourseStore) CheckCourseProfessorDuplicate(courseName string, teacherId string) (bool, error)
{
	return nil, nil
}
func (mcs *mockCourseStore) InsertIntoUserCourses(c *models.Course, userid string) error
{
	return nil, nil
}

