package domain

import (
	"fmt"
	"reflect"

	"github.com/n30w/Darkspace/internal/models"
)

type UserStore interface {
	InsertUser(u *models.User) error
	GetUserByID(u *models.User) (*models.User, error)
	GetUserByEmail(c models.Credential) (*models.User, error)
	// GetUserByUsername(username models.Credential) (*models.User, error)
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

func (us *UserService) ValidateUser(email string, password string) error {
	user, err := us.store.GetUserByEmail(Email(email))
	fmt.Printf("%s :: %s", user.Password.String(), password)
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

// What if we want only some information from Assignments or Courses?
func (us *UserService) RetrieveFromUser(
	userid string,
	field string,
) (interface{}, error) {
	// TEMP
	m := &models.User{}
	m.ID = userid
	user, err := us.store.GetUserByID(m)
	if err != nil {
		return nil, err
	}
	if field == "Courses" {
		courses, err := us.store.GetUserCourses(m)
		if err != nil {
			return nil, err
		}
		return courses, err
	}

	model := reflect.ValueOf(user).Elem()
	fieldValue := model.FieldByName(field)

	if !fieldValue.IsValid() {
		return nil, fmt.Errorf(
			"field %s does not exist or is uninitialized",
			field,
		)
	}

	return fieldValue.Interface(), nil
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
