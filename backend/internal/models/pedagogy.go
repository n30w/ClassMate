package models

import (
	"github.com/google/uuid"
)

type CustomId uuid.UUID
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
	Post
	Submission []string
	Feedback   string
	Grade      float64
	DueDate    time.Time `json:"due_date"`
}

type Submission struct {
	Entity
	AssignmentId string
	UserId       string
	Feedback     string
	Grade        int
	Media        *Media

	SubmissionTime string
	OnTime         bool
}

type Course struct {
	Name        string     `json:"name"`
	ID          string     `json:"id"`
	User           User
	FileType       string
	SubmissionTime time.Time

	// TODO write on time calculation method.
	OnTime bool
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
	UserNetID   string     `json:"user_net_id"`
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
