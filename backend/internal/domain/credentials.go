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

	return nil
}

type Password struct{ Value string }

func (p Password) Valid() error {
	if p.Value == "" {
		return errors.New("password field empty")
	}

	return nil
}

func (p Password) String() string {
	return fmt.Sprintf("%s", p.Value)
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
