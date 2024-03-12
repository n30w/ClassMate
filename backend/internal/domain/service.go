package domain

import "github.com/n30w/Darkspace/internal/dal"

type Service struct {
	UserService   *UserService
	CourseService *CourseService
}

func NewServices(s *dal.Store) *Service {
	return &Service{
		UserService:   NewUserService(s),
		CourseService: NewCourseService(s),
	}
}
