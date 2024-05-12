package models

import (
	"database/sql"
	"time"
)

// Credentials are user credentials gathered from the JSON request body.
// They represent custom types that implement the credential interface method.
type Credentials struct {
	Username   Credential `json:"username,omitempty"`
	Password   Credential `json:"password,omitempty"`
	Email      Credential `json:"email,omitempty"`
	Membership Credential `json:"membership,omitempty"`
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
	FullName string `json:"full_name"`

	ProfilePicture Media  `json:"profile_picture,omitempty"`

	// Projects       []Project `json:"projects,omitempty"`
	Courses []string `json:"courses,omitempty"`
	Bio     string   `json:"bio,omitempty"`
}

// NewUser creates a new user based on provided parameter
// information. It also sets the default access permissions
// and membership.
func NewUser(netId string, c Credentials, fullName string) (*User, error) {
	var err error

	err = c.Username.Valid()
	if err != nil {
		return nil, err
	}

	err = c.Password.Valid()
	if err != nil {
		return nil, err
	}

	err = c.Email.Valid()
	if err != nil {
		return nil, err
	}

	err = c.Membership.Valid()
	if err != nil {
		return nil, err
	}

	return &User{
		Entity: Entity{
			ID:        netId,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: sql.NullTime{},
		},
		Credentials: c,
		FullName:    fullName,
	}, nil
}
