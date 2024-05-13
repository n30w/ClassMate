package models

import "time"

type Post struct {
	Entity
	Title       string   `json:"title,omitempty"`
	Description string   `json:"description,omitempty"`
	Media       []string `json:"media,omitempty"`
	Date        string   `json:"date,omitempty"`
	Course      string   `json:"course,omitempty"`

	// Owner is often the Net ID.
	Owner string `json:"owner,omitempty"`
}

func NewPost(title, description, owner string) *Post {
	return &Post{
		Title:       title,
		Description: description,
		Owner:       owner,
	}
}

type Assignment struct {
	Post
	Submission []string  `json:"submission,omitempty"`
	DueDate    time.Time `json:"due_date"`
}

func NewAssignment() *Assignment {
	return &Assignment{}
}

type Submission struct {
	Entity
	Grade          float64 `json:"grade,omitempty"`
	AssignmentId   string
	User           User
	SubmissionTime time.Time
	Media          []string
	Feedback       string
	OnTime         bool
}

func NewSubmission() *Submission {
	return &Submission{}
}

// IsOnTime checks if an assignment's submission time is
// submitted on or before its due date, returning either true or false.
// This function is a variation of the one found here:
// https://stackoverflow.com/a/34100548/20087581
func (s *Submission) IsOnTime(due time.Time) bool {
	loc, _ := time.LoadLocation("UTC")

	dueDate := due.In(loc)
	s.SubmissionTime = s.SubmissionTime.In(loc)
	dur := dueDate.Sub(s.SubmissionTime)

	// If the duration is less than 0, that means the assignment is
	// not on time.
	return dur.Seconds() < 0
}

type Course struct {
	Entity
	Title       string     `json:"name"`
	Description string     `json:"description"`
	Messages    [10]string `json:"discussions"` //announcements + discussions
	Teachers    []string   `json:"teachers"`
	Roster      []string   `json:"roster"`
	Assignments []string   `json:"assignments"`
	Archived    bool       `json:"archived"`

	// UUID of the banner
	Banner string `json:"banner"`

	OnTime bool `json:"on_time"`
}

type Message struct {
	Post
	Comments []string
	Type     bool // false if discussion, true if announcement
}

func NewMessage(title, description, owner string, t bool) *Message {
	return &Message{
		Post: *NewPost(title, description, owner),
		Type: t,
	}
}

type Comment struct {
	Post
	ID      string
	Replies []string
}

// type Project struct {
// 	Name            string     `json:"name"`
// 	string              string     `json:"id"`
// 	Deadline        time.Time  `json:"deadline"`
// 	MediaReferences []MediaId  `json:"media_references"`
// 	Members         []User     `json:"members"`
// 	Discussion      Discussion `json:"discussion"`
// }
