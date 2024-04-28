package models

import "time"

type Post struct {
	Entity
	Title       string
	Description string
	Media       []string
	Date        string
	Course      string
	Owner       string
}

type Assignment struct {
	Post       `json:"post"`
	Submission []string  `json:"submission,omitempty"`
	Feedback   string    `json:"feedback,omitempty"`
	DueDate    time.Time `json:"due_date"`
}

type Submission struct {
	Entity
	Grade          float64 `json:"grade,omitempty"`
	AssignmentId   string
	User           User
	FileType       string
	SubmissionTime time.Time
	Media          *Media
	Feedback       string
	OnTime         bool
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

	// TODO write on time calculation method.
	OnTime bool
}

type Message struct {
	Post
	ID       string
	Comments []string
	Type     uint8 // 0 if discussion, 1 if announcement
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
