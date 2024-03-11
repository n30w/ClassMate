package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/n30w/Darkspace/internal/dao"
)

const version = "1.0.0"

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
	}
}

type application struct {
	config config
	logger *log.Logger
	models *dao.Models
}

func main() {

	var cfg config

	flag.IntVar(&cfg.port, "port", 6789, "API server port")
	flag.StringVar(
		&cfg.env,
		"env",
		"development",
		"Environment (development|staging|production)",
	)

	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()

	app := &application{
		config: cfg,
		logger: logger,
		models: dao.NewModels(db),
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
