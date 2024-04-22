package domain

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/n30w/Darkspace/internal/models"
)

type CourseStore interface {
	InsertCourse(c *models.Course) error
	GetCourseByName(name string) (*models.Course, error)
	GetCourseByID(courseid models.CourseId) (*models.Course, error)
	GetRoster(courseid models.CourseId) ([]models.User, error)
	ChangeCourseName(c *models.Course, name string) error
	DeleteCourse(c *models.Course) error
	AddStudent(c *models.Course, userid string) (*models.Course, error)
	RemoveStudent(c *models.Course, userid string) (*models.Course, error)
}

type CourseService struct {
	store CourseStore
}

func NewCourseService(c CourseStore) *CourseService { return &CourseService{store: c} }

func (cs *CourseService) ValidateID(id models.CourseId) bool {
	return true
}

func (cs *CourseService) CreateCourse(c *models.Course) error {
	// Check if course already exists. Can also try and do fuzzy name matching.
	_, err := cs.store.GetCourseByName(c.Name)
	if err != nil {
		return err
	}
	newUUID := uuid.New()
	c.ID = models.CourseId(newUUID)
	// Create the course.
	err = cs.store.InsertCourse(c)
	if err != nil {
		return err
	}
	return nil
}

func (cs *CourseService) RetrieveCourse(courseid models.CourseId) (
	*models.Course,
	error,
) {
	if !cs.ValidateID(courseid) {
		return nil, fmt.Errorf("invalid course ID: %s", courseid)
	}
	c, err := cs.store.GetCourseByID(courseid)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (cs *CourseService) RetrieveRoster(courseid models.CourseId) (
	[]models.User,
	error,
) {
	if !cs.ValidateID(courseid) {
		return nil, fmt.Errorf("invalid course ID: %s", courseid)
	}
	c, err := cs.store.GetRoster(courseid)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (cs *CourseService) AddToRoster(
	courseid models.CourseId,
	userid string,
) (*models.Course, error) {
	if !cs.ValidateID(courseid) {
		return nil, fmt.Errorf("invalid course ID: %s", courseid)
	}
	c, err := cs.store.GetCourseByID(courseid)
	if err != nil {
		return nil, err
	}
	c, err = cs.store.AddStudent(c, userid)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (cs *CourseService) RemoveFromRoster(
	courseid models.CourseId,
	userid string,
) (*models.Course, error) {
	if !cs.ValidateID(courseid) {
		return nil, fmt.Errorf("invalid course ID: %s", courseid)
	}
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
	courseid models.CourseId,
	name string,
) (*models.Course, error) {
	if !cs.ValidateID(courseid) {
		return nil, fmt.Errorf("invalid course ID: %s", courseid)
	}
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

func (cs *CourseService) DeleteCourse(courseid models.CourseId) error {
	if !cs.ValidateID(courseid) {
		return fmt.Errorf("invalid course ID: %s", courseid)
	}
	c, err := cs.store.GetCourseByID(courseid)
	if err != nil {
		return err
	}
	err = cs.store.DeleteCourse(c)
	if err != nil {
		return err
	}
	return nil
}
