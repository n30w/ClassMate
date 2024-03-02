package models

import (
	"time"
)

type User struct {
	Username       string    `json:"username"`
	Password       string    `json:"password"`
	Uuid           string    `json:"uuid"`
	JoinDate       time.Time `json:"join_date"`
	ModifiedDate   time.Time `json:"modified_date"`
	FullName       string    `json:"full_name"`
	ProfilePicture Media     `json:"profile_picture"`
	Projects       []Project `json:"projects"`
	Courses        []Course  `json:"courses"`
	Bio            string    `json:"bio"`
}

type Student struct {
	User User
}

type Teacher struct {
	User User
}
