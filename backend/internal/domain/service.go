package domain

import "github.com/n30w/Darkspace/internal/dal"

type Service struct {
	UserService           *UserService
	CourseService         *CourseService
	MessageService        *MessageService
	AssignmentService     *AssignmentService
	SubmissionService     *SubmissionService
	ExcelService          *ExcelService
	MediaService          *MediaService
	AuthenticationService *AuthenticationService
}

func NewServices(s *dal.Store, e *dal.ExcelStore) *Service {
	return &Service{
		UserService:           NewUserService(s),
		CourseService:         NewCourseService(s),
		MessageService:        NewMessageService(s),
		AssignmentService:     NewAssignmentService(s),
		SubmissionService:     NewSubmissionService(s),
		ExcelService:          NewExcelService(e),
		MediaService:          NewMediaService(s),
		AuthenticationService: NewAuthenticationService(s),
	}
}

type action int

const (
	Add action = iota
	Delete
)
