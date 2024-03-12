package models

import (
	"time"
)

type Post struct {
	Title       string
	ID          string
	Description string
	Media       []Media
	Date        time.Time
	Owner       User
}

type Assignment struct {
	Post       Post
	Submission []Submission
	Feedback   string
	Grade      int
	DueDate    time.Time `json:"due_date"`
}

func (a Assignment) addSubmission(submission Submission) {
	a.Submission = append(a.Submission, submission)
}

type Submission struct {
	User           User
	FileType       string
	SubmissionTime time.Time
	OnTime         bool
}

type Course struct {
	Name        string         `json:"name"`
	ID          int64          `json:"id"`
	Discussions [10]Discussion `json:"discussions"`
	Teachers    []User         `json:"teachers"`
	Roster      []User         `json:"roster"`
	Assignments []Assignment   `json:"assignments"`
	Archived    bool           `json:"archived"`
}

// Discussion contains anything related to communications,
// such as discussion posts and user messages.
type Discussion struct {
	Post     Post
	Comments []Comment
}

func (d Discussion) addComment(comment Comment) {
	d.Comments = append(d.Comments, comment)
}

// TODO: Linked lists
type Comment struct {
	Post    Post
	Replies []Comment
}

type Project struct {
	Name            string     `json:"name"`
	ID              string     `json:"id"`
	Deadline        time.Time  `json:"deadline"`
	MediaReferences []Media    `json:"media_references"`
	Members         []User     `json:"members"`
	Discussion      Discussion `json:"discussion"`
}
