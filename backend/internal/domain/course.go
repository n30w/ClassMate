package domain

import "github.com/n30w/Darkspace/internal/models"

type CourseStore interface {
	InsertCourse(c *models.Course) error
	GetCourseByName(name string) (*models.Course, error)
	GetCourseByID(courseid string) (*models.Course, error)
	GetRoster(courseid string) ([]models.User, error)
	ChangeCourseName(c *models.Course, name string) error
	DeleteCourse(c *models.Course) error
	AddStudent(c *models.Course, userid string) (*models.Course, error)
	RemoveStudent(c *models.Course, userid string) (*models.Course, error)
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

func (cs *CourseService) RetrieveCourse(courseid string) (*models.Course, error) {
	c, err := cs.store.GetCourseByID(courseid)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (cs *CourseService) RetrieveRoster(courseid string) ([]models.User, error) {
	// TODO fix this implementation
	c, err := cs.store.GetRoster(courseid)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (cs *CourseService) AddToRoster(courseid string, userid string) (*models.Course, error) {
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

func (cs *CourseService) RemoveFromRoster(courseid string, userid string) (*models.Course, error) {
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

func (cs *CourseService) UpdateCourseName(courseid string, name string) (*models.Course, error) {
	c, err := cs.RetrieveCourse(courseid)
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
	return nil
}
