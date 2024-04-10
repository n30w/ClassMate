package dal

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/lib/pq"
	"github.com/n30w/Darkspace/internal/models"
)

func TestStore_InsertUser(t *testing.T) {
	db, err := sql.Open("postgres", "postgres://reesedychiao:Parker84!@localhost/test_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	store := NewStore(db)

	cred := &models.Credentials{
		Username:  Username("testuser"),
		Password:  Password("testpassword"),
		Email:     Email("test@example.com"),
		CreatedAt: Membership("2022-04-15"),
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
	db, err := sql.Open("postgres", "postgres://username:password@localhost/test_db?sslmode=disable")
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
	db, err := sql.Open("postgres", "postgres://username:password@localhost/test_db?sslmode=disable")
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
	db, err := sql.Open("postgres", "postgres://username:password@localhost/test_db?sslmode=disable")
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
