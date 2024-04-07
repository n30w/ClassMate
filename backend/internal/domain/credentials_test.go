package domain

import (
	"github.com/n30w/Darkspace/internal/models"
	"testing"
)

func Test_validateCredentials(t *testing.T) {
	valid := models.Credentials{
		Username:   Username("smartbunnypants123"),
		Password:   Password("validPass12@vaso(#0jlkm.Q"),
		Email:      Email("scamyu@nyu.edu"),
		Membership: Membership(0),
	}

	u := &models.User{Credentials: valid}
	err := validateCredentials(u)
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestEmail_Valid(t *testing.T) {
	var e Email

	t.Run(
		"empty field", func(t *testing.T) {
			e = ""
			err := e.Valid()
			if err == nil {
				t.Errorf("invalid validity")
			}
		},
	)

	t.Run(
		"does not contain nyu.edu", func(t *testing.T) {
			e = "randomguy@yahoo.com"
			err := e.Valid()
			if err == nil {
				t.Errorf("invalid validity")
			}
		},
	)

	t.Run(
		"no TLD", func(t *testing.T) {
			e = "randomguy@nyu"
			err := e.Valid()
			if err == nil {
				t.Errorf("invalid validity")
			}
		},
	)

	t.Run(
		"two character email", func(t *testing.T) {
			e = "ab@nyu.edu"
			err := e.Valid()
			if err == nil {
				t.Errorf("invalid validity")
			}
		},
	)
}

func TestUsername_Valid(t *testing.T) {
	var u Username

	t.Run(
		"empty field", func(t *testing.T) {
			u = ""
			err := u.Valid()
			if err == nil {
				t.Errorf("invalid validity")
			}
		},
	)

	t.Run(
		"less than 3 characters", func(t *testing.T) {
			u = "abc"
			err := u.Valid()
			if err != nil {
				t.Errorf("invalid validity")
			}
		},
	)
}

func TestPassword_Valid(t *testing.T) {
	var p Password

	t.Run(
		"empty field", func(t *testing.T) {
			p = ""
			err := p.Valid()
			if err == nil {
				t.Errorf("invalid validity")
			}
		},
	)

	t.Run(
		"too short", func(t *testing.T) {
			p = "abc"
			err := p.Valid()
			if err == nil {
				t.Errorf("invalid validity")
			}
		},
	)

	t.Run(
		"No numbers", func(t *testing.T) {
			p = "aBcdefghijk"
			err := p.Valid()
			if err == nil {
				t.Errorf("invalid validity")
			}
		},
	)

	t.Run(
		"One number", func(t *testing.T) {
			p = "aBcdefghijk3"
			err := p.Valid()
			if err == nil {
				t.Errorf("invalid validity")
			}
		},
	)

	t.Run(
		"no special characters", func(t *testing.T) {
			p = "aBcdefghijk39"
			err := p.Valid()
			if err == nil {
				t.Errorf("invalid validity")
			}
		},
	)

	t.Run(
		"all lowercase", func(t *testing.T) {
			p = "abcdefghijk39"
			err := p.Valid()
			if err == nil {
				t.Errorf("invalid validity")
			}
		},
	)
}

func TestMembership_Valid(t *testing.T) {
	var m Membership

	t.Run(
		"less than 0", func(t *testing.T) {
			m = -1
			err := m.Valid()
			if err == nil {
				t.Errorf("invalid validity")
			}
		},
	)

	t.Run(
		"greater than 1", func(t *testing.T) {
			m = 2
			err := m.Valid()
			if err == nil {
				t.Errorf("invalid validity")
			}
		},
	)

	t.Run(
		"equal to 0", func(t *testing.T) {
			m = 0
			err := m.Valid()
			if err != nil {
				t.Errorf("%s", err)
			}
		},
	)

	t.Run(
		"equal to 1", func(t *testing.T) {
			m = 1
			err := m.Valid()
			if err != nil {
				t.Errorf("%s", err)
			}
		},
	)
}
