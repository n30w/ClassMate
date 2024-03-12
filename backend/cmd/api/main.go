package main

import (
	"flag"
	"fmt"
	"github.com/n30w/Darkspace/internal/domain"
	"log"
	"net/http"
	"os"
	"time"

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

	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()

	store := dal.NewStore(db)

	app := &application{
		config:   cfg,
		logger:   logger,
		models:   dal.NewModels(db),
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
