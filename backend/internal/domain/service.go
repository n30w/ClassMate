package domain

import "github.com/n30w/Darkspace/internal/dal"

type Service struct {
	UserService       *UserService
	CourseService     *CourseService
	MessageService    *MessageService
	AssignmentService *AssignmentService
	SubmissionService *SubmissionService
	ExcelService      *ExcelService
	MediaService      *MediaService
	UserService           *UserService
	CourseService         *CourseService
	MessageService        *MessageService
	AssignmentService     *AssignmentService
	AuthenticationService *AuthenticationService
}

func NewServices(s *dal.Store) *Service {
	return &Service{
		UserService:       NewUserService(s),
		CourseService:     NewCourseService(s),
		MessageService:    NewMessageService(s),
		AssignmentService: NewAssignmentService(s),
		SubmissionService: NewSubmissionService(s),
		ExcelService:      NewExcelService(s),
		MediaService:      NewMediaService(s),
	}
}
		UserService:           NewUserService(s),
		CourseService:         NewCourseService(s),
		MessageService:        NewMessageService(s),
		AssignmentService:     NewAssignmentService(s),
		AuthenticationService: NewAuthenticationService(s),
	}
}

type action int

const (
	Add action = iota
	Delete
)
