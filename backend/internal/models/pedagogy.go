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

type Student struct {
	User User
}

type Teacher struct {
	User User
}
