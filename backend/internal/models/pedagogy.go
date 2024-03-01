package models

import "time"

type Assignment struct {
	Name        string
	Uuid        string
	Media       Media
	Discussion  Discussion
	DueDate     time.Time
	Description string
}

type Course struct {
	Name        string
	Uuid        string
	Discussions [10]Discussion
	Teachers    []Teacher
	Roster      []Student
	Assignments []Assignment
	Archived    bool
}

// Contains anything related to communications,
// such as discussion posts and user messages.
type Discussion struct {
	name            string
	description     string
	participants    []User
	mediaReferences []Media
	commentThreads  []string
}

type Project struct {
	Name            string
	Uuid            string
	Deadline        time.Time
	MediaReferences []Media
	Members         []User
	Discussion      Discussion
}
