package dal

import (
	"database/sql"
	"github.com/n30w/Darkspace/internal/domain"
	"log"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/n30w/Darkspace/internal/models"
)

func TestStore_InsertUser(t *testing.T) {
	db, err := sql.Open(
		"postgres",
		"postgres://reesedychiao:Parker84!@localhost/test_db?sslmode=disable",
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	store := NewStore(db)

	cred := models.Credentials{
		Username:   domain.Username("testuser"),
		Password:   domain.Password("testpassword"),
		Email:      domain.Email("test@example.com"),
		Membership: domain.Membership(0),
	}

	user := &models.User{
		Credentials: cred,
	}

	err = store.InsertUser(user)
	if err != nil {
		t.Errorf("Insert user unsuccessful")
	}
}

func Username(s string) {
	panic("unimplemented")
}

func TestGetUserByID(t *testing.T) {
	db, err := sql.Open(
		"postgres",
		"postgres://username:password@localhost/test_db?sslmode=disable",
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	store := NewStore(db)

	entity := models.Entity{
		ID:        "abc1234",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: sql.NullTime{},
	}

	cred := models.Credentials{
		Username:   domain.Username("okay"),
		Password:   domain.Password("password"),
		Email:      domain.Email("abc213@nyu.edu"),
		Membership: domain.Membership(0),
	}

	expectedUser := &models.User{
		Entity:      entity,
		Credentials: cred,
	}

	err = store.InsertUser(expectedUser)
	if err != nil {
		t.Errorf("Insert user unsuccessful")
	}

	user, err := store.GetUserByID("1")
	if expectedUser.ID != user.ID {
		t.Errorf("Wrong user ID")
	}
	if expectedUser.Email != user.Email {
		t.Errorf("Wrong user email")
	}
	if expectedUser.FullName != user.FullName {
		t.Errorf("Wrong user fullname")
	}
}

func TestGetUserByEmail(t *testing.T) {
	db, err := sql.Open(
		"postgres",
		"postgres://username:password@localhost/test_db?sslmode=disable",
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	store := dal.NewStore(db)

	expectedUser := &models.User{
		ID:       "1",
		Email:    "test@example.com",
		FullName: "Test User",
	}

	err = store.InsertUser(user)
	if err != nil {
		t.Errorf("Insert user unsuccessful")
	}

	user, err := store.GetUserByEmail("test@example.com")
	if expectedUser.ID != user.ID {
		t.Errorf("Wrong user ID")
	}
	if expectedUser.Email != user.Email {
		t.Errorf("Wrong user email")
	}
	if expectedUser.FullName != user.FullName {
		t.Errorf("Wrong user fullname")
	}
}

func TestGetUserByUsername(t *testing.T) {
	db, err := sql.Open(
		"postgres",
		"postgres://username:password@localhost/test_db?sslmode=disable",
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	store := dal.NewStore(db)

	expectedUser := &models.User{
		ID:       "1",
		Email:    "test@example.com",
		FullName: "Test User",
	}

	err = store.InsertUser(user)
	if err != nil {
		t.Errorf("Insert user unsuccessful")
	}

	user, err := store.GetUserByUsername("testuser")
	if expectedUser.ID != user.ID {
		t.Errorf("Wrong user ID")
	}
	if expectedUser.Email != user.Email {
		t.Errorf("Wrong user email")
	}
	if expectedUser.FullName != user.FullName {
		t.Errorf("Wrong user fullname")
	}
}
