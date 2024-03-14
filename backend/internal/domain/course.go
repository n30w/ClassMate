package domain

import "github.com/n30w/Darkspace/internal/models"

type CourseStore interface {
	InsertCourse(c *models.Course) error
	GetCourseByName(name string) (*models.Course, error)
	GetCourseByID(id string) (*models.Course, error)
	GetRoster(id string) ([]models.User, error)
}

type CourseService struct {
	store CourseStore
}

func NewCourseService(c CourseStore) *CourseService { return &CourseService{store: c} }

func (cs *CourseService) CreateCourse(c *models.Course) error {
	// Check if course already exists. Can also try and do fuzzy name matching.
	_, err := cs.store.GetCourseByName(c.Name)
	if err != nil {
		return err
	}

	// Create the course.
	err = cs.store.InsertCourse(c)
	if err != nil {
		return err
	}

	return nil
}

func (cs *CourseService) RetrieveCourse(id string) (*models.Course, error) {
	c, err := cs.store.GetCourseByID(id)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (cs *CourseService) RetrieveRoster(id string) ([]models.User, error) {
	// TODO fix this implementation
	c, err := cs.store.GetRoster(id)
	if err != nil {
		return nil, err
	}

	return c, nil
}
