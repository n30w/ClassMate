package domain

type Store interface {
	UserStore
	CourseStore
	MessageStore
	AssignmentStore
	AuthenticationStore
	SubmissionStore
	ExcelStore
	FileStore
}
