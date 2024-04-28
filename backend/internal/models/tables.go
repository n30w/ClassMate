package models

type Table int

const (
	USERS Table = iota
	COURSES
	MEDIA
	PROJECTS
	MESSAGES
	ASSIGNMENTS
	SUBMISSIONS
	USER_COURSES
	COURSE_MESSAGES
	COURSE_TEACHERS
	COURSE_ROSTER
	COURSE_ASSIGNMENTS
	ASSIGNMENT_SUBMISSIONS
	MESSAGE_MEDIA
)

func (t Table) String() string {
	switch t {
	case USERS:
		return "users"
	case COURSES:
		return "courses"
	case MEDIA:
		return "media"
	case PROJECTS:
		return "projects"
	case MESSAGES:
		return "messages"
	case ASSIGNMENTS:
		return "assignments"
	case SUBMISSIONS:
		return "submissions"
	case USER_COURSES:
		return "user_courses"
	case COURSE_MESSAGES:
		return "course_messages"
	case COURSE_TEACHERS:
		return "course_teachers"
	case COURSE_ROSTER:
		return "course_roster"
	case COURSE_ASSIGNMENTS:
		return "course_assignments"
	case ASSIGNMENT_SUBMISSIONS:
		return "assignment_submissions"
	case MESSAGE_MEDIA:
		return "message_media"
	default:
		return "invalid table"
	}
}
