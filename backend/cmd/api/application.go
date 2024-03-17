package main

import (
	"context"
	"database/sql"
	"fmt"
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
	db struct {
		// Database driver and DataSourceName
		driver string
		dsn    string

		// Database parameters, similar to .env file variables.
		name     string
		username string
		password string
		host     string
		port     string
		sslMode  string

		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
}

// openDB opens a connection to the database using a certain config.
func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open(cfg.db.driver, cfg.createDataSourceName())
	if err != nil {
		return nil, err
	}

	// Passing a value less than or equal to 0 means no limit.
	db.SetMaxOpenConns(cfg.db.maxOpenConns)

	// Passing a value less than or equal to 0 means no limit.
	db.SetMaxIdleConns(cfg.db.maxIdleConns)

	duration, err := time.ParseDuration(cfg.db.maxIdleTime)
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

func (cfg config) createDataSourceName() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.db.host,
		cfg.db.port,
		cfg.db.username,
		cfg.db.password,
		cfg.db.name,
		cfg.db.sslMode,
	)
}

type application struct {
	config   config
	logger   *log.Logger
	services *domain.Service
}
