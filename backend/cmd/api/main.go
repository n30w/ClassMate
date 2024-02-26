package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	// port the server will run on.
	port int

	// runtime environment, either "development", "staging", or "production".
	env string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {

	var cfg config

	flag.IntVar(&cfg.port, "port", 6789, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

	// Course CRUD operations
	mux.HandleFunc("/v1/course/create", app.courseCreateHandler)
	mux.HandleFunc("/v1/course/{id}", app.courseReadHandler)
	mux.HandleFunc("/v1/course/{id}/update", app.courseUpdateHandler)

	// User CRUD operations
	mux.HandleFunc("/v1/user/create", app.userCreateHandler)
	mux.HandleFunc("/v1/user/{id}", app.userReadHandler)
	mux.HandleFunc("/v1/user/{id}/update", app.userUpdateHandler)
	mux.HandleFunc("/v1/user/{id}/delete", app.userDeleteHandler)

	// Assignment CRUD operations
	mux.HandleFunc("/v1/course/assignment/create", app.assignmentCreateHandler)
	mux.HandleFunc("/v1/course/assignment/{id}", app.assignmentReadHandler)
	mux.HandleFunc("/v1/course/assignment/{id}/update", app.assignmentUpdateHandler)
	mux.HandleFunc("/v1/course/assignment/{id}/delete", app.assignmentDeleteHandler)

	// Discussion CRUD operations
	mux.HandleFunc("/v1/course/discussion/create", app.discussionCreateHandler)
	mux.HandleFunc("/v1/course/discussion/{id}", app.discussionReadHandler)
	mux.HandleFunc("/v1/course/discussion/{id}/update", app.discussionUpdateHandler)
	mux.HandleFunc("/v1/course/discussion/{id}/delete", app.discussionDeleteHandler)

	// Login will require authorization, body will contain the credential info
	mux.HandleFunc("/v1/user/login", app.userLoginHandler)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.env, server.Addr)

	err := server.ListenAndServe()

	logger.Fatal(err)

}
