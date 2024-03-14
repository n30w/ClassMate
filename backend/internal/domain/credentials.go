package domain

import (
	"errors"
	"fmt"

	"github.com/n30w/Darkspace/internal/models"
)

// validateCredentials validates credentials using credentials interface
// method. Firstly, the credentials are checked if they are blank.
// Then, in each Valid() method, specific requirements for the credentials are
// checked.
func validateCredentials(c *models.User) error {
	var err error

	err = c.Credentials.Username.Valid()
	if err != nil {
		return err
	}

	err = c.Credentials.Password.Valid()
	if err != nil {
		return err
	}

	err = c.Credentials.Email.Valid()
	if err != nil {
		return err
	}

	err = c.Credentials.Membership.Valid()
	if err != nil {
		return err
	}

	return nil
}

type Password string

func (p Password) Valid() error {
	if p == "" {
		return errors.New("password field empty")
	}

	return nil
}

func (p Password) String() string {
	return fmt.Sprintf("%s", p)
}

type Username string

func (u Username) Valid() error {
	if u == "" {
		return errors.New("username field empty")
	}

	return nil
}

func (u Username) String() string {
	return fmt.Sprintf("%s", u)
}

type Email string

func (e Email) Valid() error {
	if e == "" {
		return errors.New("email field empty")
	}

	return nil
}

func (e Email) String() string {
	return fmt.Sprintf("%s", e)
}

type Membership int

func (m Membership) Valid() error {
	if m < 0 || m > 1 {
		return errors.New("membership must either be 0 or 1")
	}

	return nil
}

func (m Membership) String() string {
	return fmt.Sprintf("%d", m)
}
