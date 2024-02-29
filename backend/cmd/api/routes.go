package main

import "net/http"

func (app *application) routes() *http.ServeMux {

	router := http.NewServeMux()

	router.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

	// Course CRUD operations
	router.HandleFunc("/v1/course/create", app.courseCreateHandler)
	router.HandleFunc("/v1/course/read", app.courseReadHandler)
	router.HandleFunc("/v1/course/update", app.courseUpdateHandler)

	// User CRUD operations
	router.HandleFunc("/v1/user/create", app.userCreateHandler)
	router.HandleFunc("/v1/user/read", app.userReadHandler)
	router.HandleFunc("/v1/user/update", app.userUpdateHandler)
	router.HandleFunc("/v1/user/delete", app.userDeleteHandler)

	// Assignment CRUD operations
	router.HandleFunc("/v1/course/assignment/create", app.assignmentCreateHandler)
	router.HandleFunc("/v1/course/assignment/read", app.assignmentReadHandler)
	router.HandleFunc("/v1/course/assignment/update", app.assignmentUpdateHandler)
	router.HandleFunc("/v1/course/assignment/delete", app.assignmentDeleteHandler)

	// Discussion CRUD operations
	router.HandleFunc("/v1/course/discussion/create", app.discussionCreateHandler)
	router.HandleFunc("/v1/course/discussion/read" app.discussionReadHandler)
	router.HandleFunc("/v1/course/discussion/update", app.discussionUpdateHandler)
	router.HandleFunc("/v1/course/discussion/delete", app.discussionDeleteHandler)

	// Login will require authorization, body will contain the credential info
	router.HandleFunc("/v1/user/login", app.userLoginHandler)

	return router
}
