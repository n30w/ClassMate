package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/n30w/Darkspace/internal/domain"

	"github.com/n30w/Darkspace/internal/dal"
)

const version = "1.0.0"

func main() {

	var cfg config

	flag.IntVar(&cfg.port, "port", 6789, "API server port")
	flag.StringVar(
		&cfg.env,
		"env",
		"development",
		"Environment (development|staging|production)",
	)

	// Database driver.
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("DB_DSN"), "PostgreSQL DSN")

	// Database configuration for connection settings.
	flag.IntVar(
		&cfg.db.maxOpenConns, "db-max-open-conns", 25,
		"PostgreSQL max open connections",
	)
	flag.IntVar(
		&cfg.db.maxIdleConns, "db-max-idle-conns", 25,
		"PostgreSQL max idle connections",
	)
	flag.StringVar(
		&cfg.db.maxIdleTime, "db-max-idle-time", "15m",
		"PostgreSQL max connection idle time",
	)

	flag.Parse()

	logger := log.New(os.Stdout, "[DKSE] ", log.Ldate|log.Ltime)

	cfg.db.driver = "postgres"

	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()

	store := dal.NewStore(db)

	app := &application{
		config:   cfg,
		logger:   logger,
		services: domain.NewServices(store),
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.env, server.Addr)

	err = server.ListenAndServe()

	logger.Fatal(err)
}
