package models

import (
	"time"

	"github.com/google/uuid"
)

type CustomId uuid.UUID

type UserId string
type TeacherId string
type CourseId CustomId
type AssignmentId CustomId
type MediaId CustomId
type SubmissionId CustomId
type MessageId CustomId
type CommentId CustomId

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
	Name        string         `json:"name"`
	ID          CourseId       `json:"id"`
	Messages    [10]MessageId  `json:"discussions"` //announcements + discussions
	Teachers    []TeacherId    `json:"teachers"`
	Roster      []UserId       `json:"roster"`
	Assignments []AssignmentId `json:"assignments"`
	Archived    bool           `json:"archived"`
}

type Message struct {
	Post     *Post
	ID       MessageId
	Comments []CommentId
	Type     uint8 // 0 if discussion, 1 if announcement
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
