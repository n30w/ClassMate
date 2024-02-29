package models

import (
	"time"
)

type User struct {
	Username       string
	Password       string
	Uuid           string
	JoinDate       time.Time
	ModifiedDate   time.Time
	FullName       string
	ProfilePicture Media
	Projects       []Project
	Courses        []Course
	Bio            string
}

type Student struct {
	User User
}

type Teacher struct {
	User User
}
