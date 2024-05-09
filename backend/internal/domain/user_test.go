package domain

import (
	"errors"
	"strconv"
	"testing"

	"github.com/n30w/Darkspace/internal/models"
)

func TestUserService_CreateUser(t *testing.T) {
	us := NewUserService(newMockUserStore())

	// cred is fake credentials.
	cred := models.Credentials{
		Username:   Username("snow"),
		Password:   Password("buTter1290310923!09q3t"),
		Email:      Email("snow@nyu.edu"),
		Membership: Membership(0),
	}

	newUser, err := models.NewUser("abc123", cred, "donald duck")

	if err != nil {
		t.Errorf("%v", err)
	}

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
		id:         0,
		byID:       make(map[string]*models.User),
		byEmail:    make(map[string]int),
		byUsername: make(map[string]int),
	}
}

type mockUserStore struct {
	id         int
	byID       map[string]*models.User
	byEmail    map[string]int
	byUsername map[string]int
}

func (mus *mockUserStore) InsertUser(u *models.User) error {
	mus.id += 1
	mus.byID[strconv.Itoa(mus.id)] = u
	mus.byEmail[u.Email.String()] = mus.id
	mus.byUsername[u.Username.String()] = mus.id
	return nil
}

func (mus *mockUserStore) GetUserByID(u *models.User) (
	*models.User,
	error,
) {
	u = mus.byID[u.ID]
	return u, nil
}

func (mus *mockUserStore) GetUserByEmail(c models.Credential) (
	*models.User,
	error,
) {
	if u, ok := mus.byEmail[c.String()]; !ok {
		return mus.byID[strconv.Itoa(u)], errors.New("email already taken")
	}
	return nil, nil
}

func (mus *mockUserStore) GetUserByUsername(username models.Credential) (
	*models.User,
	error,
) {
	if u, ok := mus.byUsername[username.String()]; !ok {
		return mus.byID[strconv.Itoa(u)],
			errors.New("username already taken")
	}
	return nil, nil
}

func (mus *mockUserStore) DeleteCourseFromUser(
	u *models.User,
	courseid string,
) error {
	return nil
}

func (mus *mockUserStore) GetMembershipById(id string) (*models.Credential, error) {
	return nil, nil
}
