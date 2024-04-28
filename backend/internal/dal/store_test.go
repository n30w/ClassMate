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
func setupDatabaseTest(t *testing.T) (*sql.DB, error) {
	t.Helper()
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
		t.Errorf("%v", err)
	}

	// Passing a value less than or equal to 0 means no limit.
	db.SetMaxOpenConns(25)

	// Passing a value less than or equal to 0 means no limit.
	db.SetMaxIdleConns(25)

	duration, err := time.ParseDuration("15m")
	if err != nil {
		t.Errorf("%v", err)
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		t.Errorf("%v", err)
	}

	t.Cleanup(
		func() {
			db.Close()
		},
	)

	store := NewStore(db)

	ei := "abc123"

	// This should match the dev-init.sql file's first entry.
	expectedUser := &models.User{
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

	expectedCourse := &models.Course{
		Entity:      models.Entity{ID: "c3b34a9f-8f59-4818-a684-9cda56f42d02"},
		Title:       "Clown Foundations",
		Description: "Learn how to be a clown",
		Messages:    [10]string{},
		Teachers:    nil,
		Roster:      nil,
		Assignments: nil,
		Archived:    false,
	}

	// #################
	// 	 REMOVAL TESTS
	// #################

	t.Run(
		"delete user by id", func(t *testing.T) {
			id := "ghi987"
			n, err := store.DeleteUserByNetID(id)
			if err != nil {
				t.Errorf("%v", err)
			}

			if n == 0 {
				t.Errorf("no rows deleted")
			}
		},
	)

	t.Run(
		"delete course by title", func(t *testing.T) {
			title := "Delete This Course"
			n, err := store.DeleteCourseByTitle(title)
			if err != nil {
				t.Errorf("%v", err)
			}

			if n == 0 {
				t.Errorf("no rows deleted")
			}
		},
	)

	// #################
	// 	RETRIEVAL TESTS
	// #################

	t.Run(
		"get user by id", func(t *testing.T) {
			u, err := store.GetUserByID("abc123")
			if err != nil {
				t.Errorf("%s", err)
			}

			if u.ID != expectedUser.ID {
				t.Errorf("got %s, want %s", u.ID, expectedUser.ID)
			}

			if u.Email.String() != expectedUser.Email.String() {
				t.Errorf("got %s, want %s", u.Email, expectedUser.Email)
			}

			if u.FullName != expectedUser.FullName {
				t.Errorf("got %s, want %s", u.FullName, expectedUser.FullName)
			}
		},
	)

	t.Run(
		"get user by email", func(t *testing.T) {
			var e email = "abc123@nyu.edu"
			u, err := store.GetUserByEmail(e)
			if err != nil {
				t.Errorf("%v", err)
			}

			if u.ID != expectedUser.ID {
				t.Errorf("got %s, want %s", u.ID, expectedUser.ID)
			}

			if u.Email.String() != expectedUser.Email.String() {
				t.Errorf("got %s, want %s", u.Email, expectedUser.Email)
			}

			if u.FullName != expectedUser.FullName {
				t.Errorf("got %s, want %s", u.FullName, expectedUser.FullName)
			}
		},
	)

	t.Run(
		"get user by username", func(t *testing.T) {
			var n username = "jcena"
			u, err := store.GetUserByUsername(n)
			if err != nil {
				t.Errorf("%v", err)
			}

			if u.ID != expectedUser.ID {
				t.Errorf("got %s, want %s", u.ID, expectedUser.ID)
			}

			if u.Email.String() != expectedUser.Email.String() {
				t.Errorf("got %s, want %s", u.Email, expectedUser.Email)
			}

			if u.FullName != expectedUser.FullName {
				t.Errorf("got %s, want %s", u.FullName, expectedUser.FullName)
			}
		},
	)

	t.Run(
		"get course by title", func(t *testing.T) {
			title := "Clown Foundations"
			c, err := store.GetCourseByName(title)
			if err != nil {
				t.Errorf("%v", err)
			}

			if c.ID != expectedCourse.ID {
				t.Errorf("got %s, want %s", c.ID, expectedCourse.ID)
			}

			if c.Description != expectedCourse.Description {
				t.Errorf(
					"got %s, want %s",
					c.Description,
					expectedCourse.Description,
				)
			}
		},
	)

	t.Run(
		"get course by id", func(t *testing.T) {
			id := "c3b34a9f-8f59-4818-a684-9cda56f42d02"
			c, err := store.GetCourseByID(id)
			if err != nil {
				t.Errorf("%v", err)
			}

			if c.Title != expectedCourse.Title {
				t.Errorf("got %s, want %s", c.Title, expectedCourse.Title)
			}

			if c.Description != expectedCourse.Description {
				t.Errorf(
					"got %s, want %s",
					c.Description,
					expectedCourse.Description,
				)
			}
		},
	)

	// #################
	//  INSERTION TESTS
	// #################

	t.Run(
		"insert user", func(t *testing.T) {
			var n username = "testuser"
			var id string = "xyz123"

			cred := models.Credentials{
				Username:   n,
				Password:   password("testpassword"),
				Email:      email("test@example.com"),
				Membership: membership(0),
			}

			t.Cleanup(
				func() {
					store.DeleteUserByNetID(id)
				},
			)

			u := &models.User{
				Entity: models.Entity{
					ID: id,
				},
				Credentials: cred,
			}

			err := store.InsertUser(u)
			if err != nil {
				t.Errorf("%v", err)
			}

			_, err = store.GetUserByUsername(n)
			if err != nil {
				t.Errorf("%v", err)
			}
		},
	)

	t.Run(
		"insert course", func(t *testing.T) {
			c := &models.Course{
				Entity: models.Entity{ID: "623208ea-7d83-4cf3-9f31" +
					"-b21d0ce151ff"},
				Title: "Nonsense and BS: An Introduction",
				Description: "Everything you need to know about getting your" +
					" way without knowing anything. " +
					"Pre-requisite to Intermediate Conning",
			}

			id, err := store.InsertCourse(c)
			if err != nil {
				t.Errorf("%v", err)
			}

			t.Cleanup(
				func() {
					store.DeleteCourseByID(id)
				},
			)

			_, err = store.GetCourseByID(c.ID)
			if err != nil {
				t.Errorf("%v", err)
			}
		},
	)

	// ################
	//  JUNCTION TESTS
	// ################

	t.Run(
		"add teacher to course", func(t *testing.T) {
			userId := "uvw321"
			err := store.AddTeacher(expectedCourse.ID, userId)
			if err != nil {
				t.Errorf("%v", err)
			}

			// Check for bidirectional representations.

			//t.Run(
			//	"teacher has course", func(t *testing.T) {
			//
			//	},
			//)
			//
			//t.Run(
			//	"course has teacher", func(t *testing.T) {
			//
			//	},
			//)
		},
	)
}
