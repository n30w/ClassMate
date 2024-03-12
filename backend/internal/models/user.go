package models

import (
	"database/sql"
	"time"
)

type member uint8

const (
	STUDENT member = iota
	TEACHER
	ADMIN
)

// User represents the concept of a User in a database.
// It is composed of an entity, the basic unit of a database object.
// The ID in this case will be the NetID,
// which is retrieved from a consumer during an API hit.
type User struct {
	Entity
	Credentials

	// General user information.
	ProfilePicture Media     `json:"profile_picture,omitempty"`
	Projects       []Project `json:"projects,omitempty"`
	Courses        []Course  `json:"courses,omitempty"`
	Bio            string    `json:"bio,omitempty"`
}

func NewUser(netid string, c Credentials) *User {
	return &User{
		Entity: Entity{
			ID:        netid,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: sql.NullTime{},
		},
		Credentials: c,
	}
}

// Credentials are user credentials gathered from the JSON request body.
// They represent custom types that implement the credential interface method.
type Credentials struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
	FullName string `json:"full_name,omitempty"`
}

func NewCredentials(
	username, password, email string,
) Credentials {
	return Credentials{
		Username: username,
		Password: password,
		Email:    email,
	}
}
