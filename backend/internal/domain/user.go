package domain

import (
	"fmt"

	"github.com/n30w/Darkspace/internal/models"
)

type UserStore interface {
	InsertUser(u *models.User) error
	GetUserByID(u *models.User) (*models.User, error)
	GetUserByEmail(c models.Credential) (*models.User, error)
	DeleteCourseFromUser(u *models.User, courseid string) error
	GetMembershipById(netid string) (*models.Credential, error)
	GetUserCourses(u *models.User) ([]models.Course, error)
}

type UserService struct {
	store UserStore
}

func NewUserService(us UserStore) *UserService {
	return &UserService{store: us}
}

func (us *UserService) ValidateUser(netid string, password string) error {
	u := &models.User{
		Entity: models.Entity{
			ID: netid,
		},
	}

	user, err := us.store.GetUserByID(u)
	if err != nil {
		return err
	}

	if user.Password.String() != password {
		return fmt.Errorf("password mismatch")
	}

	return nil
}

// CreateUser validates User model values, and if all is well,
// creates the user in the database.
func (us *UserService) CreateUser(um *models.User) error {
	// TEMP
	// m := &models.User{}
	// First check if user exists.
	// _, err := us.store.GetUserByID(m)
	// _, err := us.store.GetUserByEmail(m.Email)
	// if err != nil {
	// 	return err
	// }

	// Check if credentials are valid.
	err := validateCredentials(um)
	if err != nil {
		return err
	}

	// Check if email is already in use.
	_, err = us.store.GetUserByEmail(um.Email)
	if err == nil {
		return fmt.Errorf("email already in use")
	}

	// // Check if username is already in use.
	// _, err = us.store.GetUserByUsername(um.Username)
	// // Notice that err IS EQUAL TO nil and not NOT EQUAL TO.
	// if err == nil {
	// 	return fmt.Errorf("username already in use")
	// }

	// If all is well...
	err = us.store.InsertUser(um)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) RetrieveHomepage() (*models.Homepage, error) {
	// hp := &models.Homepage{}

	return nil, nil
}

func (us *UserService) GetByID(userid string) (*models.User, error) {
	// TEMP
	m := &models.User{}
	m.ID = userid
	user, err := us.store.GetUserByID(m)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// RetrieveFromUser currently is used for the homepage. It just retrieves
// the user's courses.
func (us *UserService) RetrieveFromUser(
	userid string,
) ([]models.Course, error) {
	// TEMP
	m := &models.User{}
	m.ID = userid
	courses, err := us.store.GetUserCourses(m)
	if err != nil {
		return nil, err
	}
	return courses, err
}

func (us *UserService) GetUserCourses(userId string) ([]models.Course, error) {
	m := &models.User{
		Entity: models.Entity{
			ID: userId,
		},
	}

	user, err := us.store.GetUserByID(m)
	if err != nil {
		return nil, err
	}

	courses, err := us.store.GetUserCourses(user)
	if err != nil {
		return nil, err
	}

	return courses, err
}

func (us *UserService) GetMembership(netid string) (*models.Credential, error) {
	membership, err := us.store.GetMembershipById(netid)
	if err != nil {
		return nil, err
	}
	return membership, nil
}

func (us *UserService) UnenrollUserFromCourse(
	userid string,
	courseid string,
) error {
	// TEMP
	m := &models.User{}
	user, err := us.store.GetUserByID(m)
	if err != nil {
		return err
	}
	err = us.store.DeleteCourseFromUser(user, courseid)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) NewUsername(s string) Username {
	return Username(s)
}

func (us *UserService) NewPassword(s string) Password {
	return Password(s)
}

func (us *UserService) NewEmail(s string) Email {
	return Email(s)
}

func (us *UserService) NewMembership(d int) Membership {
	return Membership(d)
}
