package models

import (
	"errors"
	"fmt"
	"time"
)

type member uint8

const (
	STUDENT member = iota
	TEACHER
	ADMIN
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`

	// The NetID serves as a UUID.
	Netid string `json:"netid"`

	// General user data
	FullName       string    `json:"full_name"`
	ProfilePicture Media     `json:"profile_picture"`
	JoinDate       time.Time `json:"join_date"`
	ModifiedDate   time.Time `json:"modified_date"`
	Projects       []Project `json:"projects"`
	Courses        []Course  `json:"courses"`
	Bio            string    `json:"bio"`

	ac *accessControl
}

func (u User) createDiscussion(title string, description string, media []Media, date time.Time) Discussion {
	return Discussion{
		Name:            title,
		Description:     description,
		MediaReferences: media,
	}
}

type userConfig struct {
	creds [4]credentials
}

// func newUserConfig() userConfig {
// 	return userConfig{
// 		creds: [4]credentials{
// 			password{},
// 			username("john"),
// 			netid("rra9981"),
// 			email("123@yahoo.com"),
// 		},
// 	}
// }

func (u userConfig) valid() error {
	for _, v := range u.creds {
		if err := v.valid(); err != nil {
			return fmt.Errorf("invalid user configuration: %s", err)
		}
	}
	return nil
}

type password struct{ value []byte }

func (p password) valid() error {
	valid := false
	if !valid {
		return errors.New("password does not match criteria")
	}
	return nil
}

type username string

func (u username) valid() error {
	// if username is already in DB, return false
	valid := false
	if !valid {
		return errors.New("username already taken")
	}
	return nil
}

type netid string

func (n netid) valid() error {
	valid := false
	if !valid {
		return errors.New("netid already in use")
	}
	return nil
}

type email string

func (e email) valid() error {
	valid := false
	if !valid {
		return errors.New("email already in use")
	}
	return nil
}

// Instantiate new users in database and session.
