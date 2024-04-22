package dal

import (
	"context"
	"database/sql"
	"github.com/n30w/Darkspace/internal/domain"
	"log"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/n30w/Darkspace/internal/models"
)

// setupDatabaseTest creates a connection to an already running
// postgresql database, for running tests.
func setupDatabaseTest() (*sql.DB, error) {
	var dbConf DBConfig

	dbConf.SetFromEnv()

	db, err := sql.Open(dbConf.Driver, dbConf.CreateDataSourceName())
	if err != nil {
		return nil, err
	}

	// Passing a value less than or equal to 0 means no limit.
	db.SetMaxOpenConns(dbConf.MaxOpenConns)

	// Passing a value less than or equal to 0 means no limit.
	db.SetMaxIdleConns(dbConf.MaxIdleConns)

	duration, err := time.ParseDuration(dbConf.MaxIdleTime)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}

type username string
type password string
type email string
type membership int

func (u username) String() string { return string(u) }
func (u username) Valid() error   { return nil }

func (p password) String() string { return string(p) }
func (p password) Valid() error   { return nil }

func (e email) String() string { return string(e) }
func (e email) Valid() error   { return nil }

func (m membership) String() string { return string(m) }
func (m membership) Valid() error   { return nil }

func TestDB(t *testing.T) {
	db, err := setupDatabaseTest()
	if err != nil {
		t.Errorf("%s", err)
	}
	defer db.Close()

	store := NewStore(db)

	// This should match the dev-init.sql file's first entry.
	expected := &models.User{
		Entity: models.Entity{ID: "abc123"},
		Credentials: models.Credentials{
			Username:   username("jcena"),
			Password:   password("password123"),
			Email:      email("abc123@nyu.edu"),
			Membership: membership(0),
		},
		FullName:       "John Cena",
		ProfilePicture: models.Media{},
		Bio:            "Can you see me?",
	}

	t.Run(
		"get user by id", func(t *testing.T) {
			u, err := store.GetUserByID("abc123")
			if err != nil {
				t.Errorf("%s", err)
			}

			if u.ID != expected.ID {
				t.Errorf("got %s, want %s", u.ID, expected.ID)
			}

			if u.Email != expected.Email {
				t.Errorf("got %s, want %s", u.Email, expected.Email)
			}

			if u.ID != expected.ID {
				t.Errorf("got %s, want %s", u.FullName, expected.FullName)
			}
		},
	)

	t.Run(
		"get user by username", func(t *testing.T) {
			u, err := store.GetUserByUsername("jcena")
			if err != nil {
				t.Errorf("%s", err)
			}

			if u.ID != expected.ID {
				t.Errorf("got %s, want %s", u.ID, expected.ID)
			}

			if u.Email != expected.Email {
				t.Errorf("got %s, want %s", u.Email, expected.Email)
			}

			if u.ID != expected.ID {
				t.Errorf("got %s, want %s", u.FullName, expected.FullName)
			}
		},
	)

	t.Run(
		"insert user", func(t *testing.T) {
			cred := models.Credentials{
				Username:   username("testuser"),
				Password:   password("testpassword"),
				Email:      email("test@example.com"),
				Membership: membership(0),
			}

			u := &models.User{Credentials: cred}

			err := store.InsertUser(u)
			if err != nil {
				t.Errorf("%s", err)
			}

			_, err = store.GetUserByUsername("testuser")
			if err != nil {
				t.Errorf("%s", err)
			}
		},
	)
}

func TestStore_InsertUser(t *testing.T) {
	db, err := setupDatabaseTest()
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
		t.Errorf("%s", err)
	}
}
