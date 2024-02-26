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
	Projects       []project
	Courses        []Course
	Bio            string
}
