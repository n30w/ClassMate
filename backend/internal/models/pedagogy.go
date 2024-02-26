package models

import "time"

type Assignment struct {
	name        string
	uuid        string
	media       Media
	discussion  Discussion
	dueDate     time.Time
	description string
}

type Course struct {
	name        string
	uuid        string
	discussions [10]Discussion
	teachers    []Teacher
	roster      []Student
	assignments []Assignment
	archived    bool
}

type Student struct {
	user User
}

type Teacher struct {
	user User
}
