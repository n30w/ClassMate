package models

import (
	"database/sql"
	"time"
)

// Credentials are user credentials gathered from the JSON request body.
// They represent custom types that implement the credential interface method.
type Credentials struct {
	Username   Credential
	Password   Credential
	Email      Credential
	Membership Credential
}

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

// NewUser creates a new user based on provided parameter
// information. It also sets the default access permissions
// and membership.
func NewUser(netid string, membership int, c Credentials) (*User, error) {
	// Check if the membership is valid
	if !
	return &User{
		Entity: Entity{
			ID:        netid,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: sql.NullTime{},
		},
		Credentials: c,
	}, nil
}
