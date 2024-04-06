package domain

import (
	"github.com/n30w/Darkspace/internal/models"
	"strconv"
	"testing"
)

func TestUserService_CreateUser(t *testing.T) {
	us := NewUserService(newMockUserStore())

	// cred is fake credentials.
	cred := models.Credentials{
		Username:   Username("snow"),
		Password:   Password("butter"),
		Email:      Email("snow@nyu.edu"),
		Membership: Membership(0),
	}

	newUser := models.NewUser("abc123", cred)

	got := us.CreateUser(newUser)

	if got != nil {
		t.Errorf("got %s", got)
	}
}

// ========= //
//   MOCKS   //
// ========= //

func newMockUserStore() *mockUserStore {
	return &mockUserStore{
		id:      0,
		byID:    make(map[string]*models.User),
		byEmail: make(map[string]int),
	}
}

type mockUserStore struct {
	id      int
	byID    map[string]*models.User
	byEmail map[string]int
}

func (mus *mockUserStore) InsertUser(u *models.User) error {
	mus.id += 1
	mus.byID[strconv.Itoa(mus.id)] = u
	mus.byEmail[u.Email.String()] = mus.id
	return nil
}

func (mus *mockUserStore) GetUserByID(id string) (
	*models.User,
	error,
) {
	u := mus.byID[id]
	return u, nil
}

func (mus *mockUserStore) GetUserByEmail(email string) (*models.User, error) {
	return nil, nil
}

func (mus *mockUserStore) GetUserByUsername(username string) (
	*models.User,
	error,
) {
	return nil, nil
}

func (mus *mockUserStore) DeleteCourseFromUser(
	courseid string,
	u *models.User,
) error {
	return nil
}
