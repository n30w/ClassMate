package models

import (
	"time"
)

type UserId string
type TeacherId string
type CourseId int64
type AssignmentId int64
type MediaId int64
type SubmissionId int64
type DiscussionId int64
type CommentId int64
type AnnouncementId int64

type Post struct {
	Title       string
	Description string
	Media       []MediaId
	Date        time.Time
	Owner       string
}

type Assignment struct {
	Post       Post
	ID         AssignmentId
	Submission []SubmissionId
	Feedback   string
	Grade      int
	DueDate    time.Time `json:"due_date"`
}

type Submission struct {
	User           User
	ID             SubmissionId
	FileType       string
	SubmissionTime time.Time
	OnTime         bool
}

type Course struct {
	Name        string           `json:"name"`
	ID          CourseId         `json:"id"`
	Discussions [10]DiscussionId `json:"discussions"`
	Teachers    []TeacherId      `json:"teachers"`
	Roster      []UserId         `json:"roster"`
	Assignments []AssignmentId   `json:"assignments"`
	Archived    bool             `json:"archived"`
}

// Discussion contains anything related to communications,
// such as discussion posts and user messages.
type Discussion struct {
	Post     Post
	ID       DiscussionId
	Comments []CommentId
}

// Announcements have the same structure as Discussions but they are displayed differently
type Announcement struct {
	Post     Post
	ID       AnnouncementId
	Comments []CommentId
}

// TODO: Linked lists
type Comment struct {
	ID      CommentId
	Post    Post
	Replies []CommentId
}

// type Project struct {
// 	Name            string     `json:"name"`
// 	ID              string     `json:"id"`
// 	Deadline        time.Time  `json:"deadline"`
// 	MediaReferences []MediaId  `json:"media_references"`
// 	Members         []User     `json:"members"`
// 	Discussion      Discussion `json:"discussion"`
// }
