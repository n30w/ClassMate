package domain

import "github.com/n30w/Darkspace/internal/dal"

type Service struct {
	UserService           *UserService
	CourseService         *CourseService
	MessageService        *MessageService
	AssignmentService     *AssignmentService
	AuthenticationService *AuthenticationService
}

func NewServices(s *dal.Store) *Service {
	return &Service{
		UserService:           NewUserService(s),
		CourseService:         NewCourseService(s),
		MessageService:        NewMessageService(s),
		AssignmentService:     NewAssignmentService(s),
		AuthenticationService: NewAuthenticationService(s),
	}
}

// validates UUID before dal operations
type Validator interface {
	ValidateID(id string) bool
}
