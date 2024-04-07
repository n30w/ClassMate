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

// The functions pass and fail denote assertions.
// In other words, a test that calls pass() expects
// the invalidity to be successfully caught. A test
// that calls fail() expects the validity to be invalid.

func TestEmail_Valid(t *testing.T) {
	var e Email
	fail := func() {
		err := e.Valid()
		if err == nil {
			t.Errorf("invalid validity")
		}
	}
	fail()
}

func TestUsername_Valid(t *testing.T) {
	var u Username
	fail := func() {
		err := u.Valid()
		if err == nil {
			t.Errorf("invalid validity")
		}
	}
	fail()
}

func TestPassword_Valid(t *testing.T) {
	var p Password

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
