package domain

import (
	"fmt"

	"github.com/n30w/Darkspace/internal/models"
)

type CourseStore interface {
	InsertCourse(c *models.Course) (string, error)
	GetCourseByID(courseid string) (*models.Course, error)
	GetRoster(c string) ([]models.User, error)
	ChangeCourseName(c *models.Course, name string) error
	DeleteCourseByID(courseid string) error
	AddStudent(c *models.Course, userid string) (*models.Course, error)
	AddTeacher(courseId, userId string) error
	RemoveStudent(c *models.Course, userid string) (*models.Course, error)
	CheckCourseProfessorDuplicate(courseName string, teacherId string) (bool, error)
	InsertIntoUserCourses(c *models.Course, userid string) error
}

type CourseService struct {
	store CourseStore
}

func NewCourseService(c CourseStore) *CourseService { return &CourseService{store: c} }

// CreateCourse creates a new course in the database,
// then assigns a UUID to it. This is not an idempotent method!
func (cs *CourseService) CreateCourse(c *models.Course, teacherid string) (*models.Course, error) {
	// Check if course already exists. Can also try and do fuzzy name matching.
	duplicate, err := cs.store.CheckCourseProfessorDuplicate(c.Title, teacherid)
	if err != nil {
		return nil, err
	}

	if duplicate {
		return nil, fmt.Errorf("course already exists")
	}

	// c.ID = uuid.New().String()

	// Create the course.
	id, err := cs.store.InsertCourse(c)
	if err != nil {
		return nil, err
	}
	c.ID = id
	err = cs.store.InsertIntoUserCourses(c, teacherid)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (cs *CourseService) RetrieveCourse(courseid string) (
	*models.Course,
	error,
) {
	c, err := cs.store.GetCourseByID(courseid)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (cs *CourseService) RetrieveRoster(courseid string) (
	[]models.User,
	error,
) {
	c, err := cs.store.GetRoster(courseid)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (cs *CourseService) AddToRoster(
	courseid string,
	userid string,
) (*models.Course, error) {
	c, err := cs.store.GetCourseByID(courseid)
	if err != nil {
		return nil, err
	}
	c, err = cs.store.AddStudent(c, userid)
	if err != nil {
		return nil, err
	}
	err = cs.store.InsertIntoUserCourses(c, userid)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (cs *CourseService) RemoveFromRoster(
	courseid string,
	userid string,
) (*models.Course, error) {
	c, err := cs.store.GetCourseByID(courseid)
	if err != nil {
		return nil, err
	}
	c, err = cs.store.RemoveStudent(c, userid)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (cs *CourseService) UpdateCourseName(
	courseid string,
	name string,
) (*models.Course, error) {
	c, err := cs.store.GetCourseByID(courseid)
	if err != nil {
		return nil, err
	}

	err = cs.store.ChangeCourseName(c, name)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (cs *CourseService) DeleteCourse(courseid string) error {
	err := cs.store.DeleteCourseByID(courseid)
	if err != nil {
		return err
	}
	return nil
}
