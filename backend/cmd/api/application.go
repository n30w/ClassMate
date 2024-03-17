package main

import (
	"github.com/n30w/Darkspace/internal/domain"
	"log"
)

type config struct {
	// Port the server will run on.
	port int

	// Runtime environment, either "development", "staging", or "production".
	env string

	// Database configurations
	db struct {
		// Database driver and DataSourceName
		driver       string
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
}

type application struct {
	config   config
	logger   *log.Logger
	services *domain.Service
}
