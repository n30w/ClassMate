package domain

import (
	"errors"
	"github.com/n30w/Darkspace/internal/models"
)

type UserStore interface {
	InsertUser(u *models.User) error
	GetUserByID(id string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
}

type UserService struct {
	store UserStore
}

func NewUserService(s UserStore) *UserService {
	return &UserService{store: s}
}

// CreateUser validates User model values, and if all is well,
// creates the user in the database.
func (us *UserService) CreateUser(um *models.User) error {
	// First check if user exists.
	_, err := us.GetByID(um.ID)
	if err != nil {
		return err
	}

	// Check if credentials are valid.
	//err = validateCredentials(um)
	//if err != nil {
	//	return err
	//}

	// Check if email is already in use.
	_, err = us.store.GetUserByEmail(um.Email)
	if err == nil {
		return errors.New("email already taken")
	}

	// Check if username is already in use.
	_, err = us.store.GetByUsername(um.Username)
	// Notice that err IS EQUAL TO nil and not NOT EQUAL TO.
	if err == nil {
		return errors.New("username already taken")
	}

	// If all is well...
	err = us.store.InsertUser(um)

	return nil
}

func (us *UserService) GetByID(id string) (*models.User, error) {
	user, err := us.store.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
