package models

import (
	"time"
)

type User struct {
	username       string
	password       string
	uuid           string
	joinDate       time.Time
	modifiedDate   time.Time
	fullName       string
	profilePicture Media
	projects       []project
	courses        []Course
	bio            string
}
