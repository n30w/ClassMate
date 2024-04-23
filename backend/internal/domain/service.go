package domain

import "github.com/n30w/Darkspace/internal/dal"

type Service struct {
	UserService       *UserService
	CourseService     *CourseService
	MessageService    *MessageService
	AssignmentService *AssignmentService
}

func NewServices(s *dal.Store) *Service {
	return &Service{
		UserService:    NewUserService(s),
		CourseService:  NewCourseService(s),
		MessageService: NewMessageService(s),
		//AssignmentService: NewAssignmentService(s),
	}
}

type action int

const (
	Add action = iota
	Delete
)
