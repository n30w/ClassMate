package dal

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
	"github.com/n30w/Darkspace/internal/models"
)

// setupDatabaseTest creates a connection to an already running
// postgresql database, for running tests.
func setupDatabaseTest() (*sql.DB, error) {
	var dbConf DBConfig
	dbConf.Driver = "postgres"
	dbConf.SetFromEnv()

	db, err := sql.Open(dbConf.Driver, dbConf.CreateDataSourceName())
	if err != nil {
		return nil, err
	}

	// Passing a value less than or equal to 0 means no limit.
	db.SetMaxOpenConns(25)

	// Passing a value less than or equal to 0 means no limit.
	db.SetMaxIdleConns(25)

	duration, err := time.ParseDuration("15m")
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

func TestDB(t *testing.T) {
	var dbConf DBConfig
	dbConf.Driver = "postgres"

	rootEnvFile := "../../.env"
	err := godotenv.Load(rootEnvFile)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbConf.Dsn = os.Getenv("DB_DSN")
	dbConf.Name = os.Getenv("DB_NAME")
	dbConf.Username = os.Getenv("DB_USERNAME")
	dbConf.Password = os.Getenv("DB_PASSWORD")
	dbConf.Host = os.Getenv("DB_HOST")
	dbConf.Port = os.Getenv("DB_PORT")
	dbConf.SslMode = os.Getenv("DB_SSL_MODE")

	db, err := sql.Open(dbConf.Driver, dbConf.CreateDataSourceName())
	if err != nil {
		t.Errorf("%s", err)
	}

	// Passing a value less than or equal to 0 means no limit.
	db.SetMaxOpenConns(25)

	// Passing a value less than or equal to 0 means no limit.
	db.SetMaxIdleConns(25)

	duration, err := time.ParseDuration("15m")
	if err != nil {
		t.Errorf("%s", err)
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		t.Errorf("%s", err)
	}

	defer db.Close()

	store := NewStore(db)

	var ei models.ID = models.ID{Serialized: "abc123"}

	// This should match the dev-init.sql file's first entry.
	expected := &models.User{
		Entity: models.Entity{ID: ei},
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

	// #################
	// 	RETRIEVAL TESTS
	// #################

	t.Run(
		"get user by id", func(t *testing.T) {
			u, err := store.GetUserByID("abc123")
			if err != nil {
				t.Errorf("%s", err)
			}

			if u.ID != expected.ID {
				t.Errorf("got %s, want %s", u.ID, expected.ID)
			}

			if u.Email.String() != expected.Email.String() {
				t.Errorf("got %s, want %s", u.Email, expected.Email)
			}

			if u.FullName != expected.FullName {
				t.Errorf("got %s, want %s", u.FullName, expected.FullName)
			}
		},
	)

	t.Run("get user by email", func(t *testing.T) {
		var e email = "abc123@nyu.edu"
		u, err := store.GetUserByEmail(e)
		if err != nil {
			t.Errorf("%s", err)
		}

		if u.ID.Serialized != expected.ID.Serialized {
			t.Errorf("got %s, want %s", u.ID, expected.ID)
		}

		if u.Email.String() != expected.Email.String() {
			t.Errorf("got %s, want %s", u.Email, expected.Email)
		}

		if u.FullName != expected.FullName {
			t.Errorf("got %s, want %s", u.FullName, expected.FullName)
		}
	})

	t.Run(
		"get user by username", func(t *testing.T) {
			var n username = "jcena"
			u, err := store.GetUserByUsername(n)
			if err != nil {
				t.Errorf("%s", err)
			}

			if u.ID.Serialized != expected.ID.Serialized {
				t.Errorf("got %s, want %s", u.ID, expected.ID)
			}

			if u.Email.String() != expected.Email.String() {
				t.Errorf("got %s, want %s", u.Email, expected.Email)
			}

			if u.FullName != expected.FullName {
				t.Errorf("got %s, want %s", u.FullName, expected.FullName)
			}
		},
	)

	t.Run("get course by id", func(t *testing.T) {
		c, err := store.GetCourseByID()
		if err != nil {
			t.Errorf("%s", err)
		}

	})

	t.Run("get course by name", func(t *testing.T) {})

	// #################
	//  INSERTION TESTS
	// #################

	t.Run(
		"insert user", func(t *testing.T) {
			var n username = "testuser"
			cred := models.Credentials{
				Username:   n,
				Password:   password("testpassword"),
				Email:      email("test@example.com"),
				Membership: membership(0),
			}

			u := &models.User{
				Entity: models.Entity{
					ID: models.ID{Serialized: "xyz123"},
				},
				Credentials: cred,
			}

			err := store.InsertUser(u)
			if err != nil {
				t.Errorf("%s", err)
			}

			_, err = store.GetUserByUsername(n)
			if err != nil {
				t.Errorf("%s", err)
			}
		},
	)

	t.Run("insert course", func(t *testing.T) {
		c := &models.Course{}

		err := store.InsertCourse(c)
		if err != nil {
			t.Errorf("%s", err)
		}

		_, err = store.GetCourseByID(c.ID.String())
		if err != nil {
			t.Errorf("%s", err)
		}

	})
}
