package main

import "net/http"

func (app *application) routes() *http.ServeMux {

	// Enhanced routing patterns in Go 1.22. HandleFuncs now
	// accept a method and a route variable parameter.
	// https://tip.golang.org/doc/go1.22

	router := http.NewServeMux()

	router.HandleFunc("GET /v1/healthcheck", app.healthcheckHandler)
	router.HandleFunc("POST /v1/home", app.homeHandler)
	router.HandleFunc("GET /v1/course/homepage/{id}", app.courseHomepageHandler)

	router.HandleFunc(
		"POST /v1/course/announcement/create/{id}",
		app.announcementCreateHandler,
	)
	router.HandleFunc(
		"POST /v1/course/announcement/update",
		app.announcementUpdateHandler,
	)
	router.HandleFunc(
		"POST /v1/course/announcement/delete",
		app.announcementDeleteHandler,
	)
	// ID is message ID
	router.HandleFunc(
		"GET /v1/course/announcement/read/{id}",
		app.announcementReadHandler,
	)
	router.HandleFunc("POST /v1/course/addstudent", app.addStudentHandler)

	// Course CRUD operations
	router.HandleFunc("POST /v1/course/create", app.courseCreateHandler)
	router.HandleFunc("GET /v1/course/read/{id}", app.courseReadHandler)
	router.HandleFunc(
		"PATCH /v1/course/update/{id}/{action}",
		app.courseUpdateHandler,
	)
	router.HandleFunc("POST /v1/course/delete/{id}", app.courseDeleteHandler)

	// User CRUD operations
	router.HandleFunc("POST /v1/user/create", app.userCreateHandler)
	router.HandleFunc("GET /v1/user/read/{id}", app.userReadHandler)
	router.HandleFunc("PATCH /v1/user/update/{id}", app.userUpdateHandler)
	router.HandleFunc("POST /v1/user/delete/{id}", app.userDeleteHandler)

	// A user posts something to a discussion
	router.HandleFunc("/v1/user/post", app.userPostHandler)

	// Login will require authorization, body will contain the credential info
	router.HandleFunc("POST /v1/user/login", app.userLoginHandler)

	// Assignment CRUD operations
	router.HandleFunc(
		"/v1/course/assignment/create",
		app.assignmentCreateHandler,
	)
	router.HandleFunc(
		"GET /v1/course/assignment/read/{id}",
		app.assignmentReadHandler,
	)
	router.HandleFunc(
		"PATCH /v1/course/assignment/update",
		app.assignmentUpdateHandler,
	)
	router.HandleFunc(
		"POST /v1/course/assignment/delete",
		app.assignmentDeleteHandler,
	)

	// Discussion CRUD operations
	router.HandleFunc(
		"/v1/course/discussion/create",
		app.discussionCreateHandler,
	)
	router.HandleFunc("/v1/course/discussion/read", app.discussionReadHandler)
	router.HandleFunc(
		"PATCH /v1/course/discussion/update",
		app.discussionUpdateHandler,
	)
	router.HandleFunc(
		"/v1/course/discussion/delete",
		app.discussionDeleteHandler,
	)

	// Media operations
	router.HandleFunc(
		"POST /v1/course/{post}/media/create",
		app.mediaCreateHandler,
	)
	router.HandleFunc(
		"POST /v1/course/{post}/media/delete",
		app.mediaDeleteHandler,
	)

	// Authentication
	// router.HandleFunc(
	// 	"POST /v1/tokens/authentication",
	// 	app.createAuthenticationTokenHandler,
	// )

	// Comment operations
	router.HandleFunc(
		"POST /v1/course/{post}/comment/create",
		app.commentCreateHandler,
	)
	router.HandleFunc(
		"POST /v1/course/{post}/comment/delete",
		app.commentDeleteHandler,
	)

	// Submission operations
	router.HandleFunc(
		"POST /v1/course/assignment/submission/create",
		app.submissionCreateHandler,
	)
	router.HandleFunc(
		"POST /v1/course/assignment/submission/update",
		app.submissionUpdateHandler,
	)
	router.HandleFunc(
		"POST /v1/course/assignment/submission/delete",
		app.submissionUpdateHandler,
	)
	router.HandleFunc(
		"POST /v1/course/assignment/submission/read",
		app.submissionUpdateHandler,
	)

	// Image operations
	// router.HandlerFunc("POST /v1/course/image", app.courseImageHandler)

	// Offline grading operations
	// Subtle difference, one is a GET, one is a POST. The POST expects
	// data to be sent along with request. The GET just sends back data.
	// The system does not need to know the ID of the course or the ID
	// of the assignment, because this should be inside the sheet
	// of the Excel document, under columns G2 and H2.
	router.HandleFunc(
		"GET /v1/course/{id}/assignment/{post}/offline",
		app.sendOfflineTemplate,
	)
	router.HandleFunc(
		"POST /v1/course/{id}/assignment/{post}/offline",
		app.receiveOfflineGrades,
	)

	return router
}
