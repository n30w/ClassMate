package main

import (
	"context"
	"database/sql"
	"github.com/n30w/Darkspace/internal/dal"
	"log"
	"time"

	"github.com/n30w/Darkspace/internal/domain"

	// This import fixes the error: "unknown driver "postgres" (forgotten import?)"
	_ "github.com/lib/pq"
)

type config struct {
	// Port the server will run on.
	port int

	// Runtime environment, either "development", "staging", or "production".
	env string

	// Database configurations
	db dal.DBConfig

	// limiter is limiter information for rate limiting.
	limiter struct {
		// rps is requests per second.
		rps   float64
		burst int

		// enabled either disables or enables rate limiting altogether.
		enabled bool
	}
}

// openDB opens a connection to the database using a certain config.
func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open(cfg.db.Driver, cfg.createDataSourceName())
	if err != nil {
		return nil, err
	}

	// Passing a value less than or equal to 0 means no limit.
	db.SetMaxOpenConns(cfg.db.MaxOpenConns)

	// Passing a value less than or equal to 0 means no limit.
	db.SetMaxIdleConns(cfg.db.MaxIdleConns)

	duration, err := time.ParseDuration(cfg.db.MaxIdleTime)
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

// createDataSourceName creates the dataSourceName parameter of the
// sql.Open function.
func (cfg config) createDataSourceName() string {
	return cfg.db.CreateDataSourceName()
}

func (cfg config) SetFromEnv() {
	cfg.db.SetFromEnv()
}

type application struct {
	config   config
	logger   *log.Logger
	services *domain.Service
}
