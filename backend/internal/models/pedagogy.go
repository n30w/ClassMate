package models

import "time"

type Assignment struct {
	Name        string     `json:"name"`
	Uuid        string     `json:"uuid"`
	Media       Media      `json:"media"`
	Discussion  Discussion `json:"discussion"`
	DueDate     time.Time  `json:"due_date"`
	Description string     `json:"description"`
}

type Course struct {
	Name        string         `json:"name"`
	Uuid        string         `json:"uuid"`
	Discussions [10]Discussion `json:"discussions"`
	Teachers    []Teacher      `json:"teachers"`
	Roster      []Student      `json:"roster"`
	Assignments []Assignment   `json:"assignments"`
	Archived    bool           `json:"archived"`
}

// Contains anything related to communications,
// such as discussion posts and user messages.
type Discussion struct {
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	Participants    []User   `json:"participants"`
	MediaReferences []Media  `json:"media_references"`
	CommentThreads  []string `json:"comment_threads"`
}

type Project struct {
	Name            string     `json:"name"`
	Uuid            string     `json:"uuid"`
	Deadline        time.Time  `json:"deadline"`
	MediaReferences []Media    `json:"media_references"`
	Members         []User     `json:"members"`
	Discussion      Discussion `json:"discussion"`
}
