package models

import (
	"database/sql"
	"time"
)

// User represents the concept of a User in a database.
// It is composed of an entity, the basic unit of a database object.
// The ID in this case will be the NetID,
// which is retrieved from a consumer during an API hit.
type User struct {
	Entity
	Credentials
	*AccessControl

	// General user information.
	FullName       string `json:"fullName,omitempty"`
	ProfilePicture Media  `json:"profile_picture,omitempty"`
	// Projects       []Project `json:"projects,omitempty"`
	Courses []Course `json:"courses,omitempty"`
	Bio     string   `json:"bio,omitempty"`
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
	Username   Credential
	Password   Credential
	Email      Credential
	Membership Credential
}
