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
	router.HandleFunc("GET /v1/course/announcement/read/{id}", app.announcementReadHandler)
	router.HandleFunc("POST /v1/course/addstudent", app.addStudentHandler)

	// Course CRUD operations
	router.HandleFunc("POST /v1/course/create", app.courseCreateHandler)
	router.HandleFunc("GET /v1/course/read/{id}", app.courseReadHandler)
	router.HandleFunc("PATCH /v1/course/update/{id}/{action}", app.courseUpdateHandler)
	router.HandleFunc("POST /v1/course/delete/{id}", app.courseDeleteHandler)
	router.HandleFunc(
		"POST /v1/course/{mediaId}/banner/create",
		app.bannerCreateHandler,
	)
	router.HandleFunc(
		"GET /v1/course/{mediaId}/banner/read",
		app.bannerReadHandler,
	)

	// User CRUD operations
	router.HandleFunc("POST /v1/user/create", app.userCreateHandler)
	router.HandleFunc("GET /v1/user/read/{id}", app.userReadHandler)
	router.HandleFunc("PATCH /v1/user/update/{id}", app.userUpdateHandler)
	router.HandleFunc("POST /v1/user/delete/{id}", app.userDeleteHandler)

	// Login will require authorization, body will contain the credential info
	router.HandleFunc("POST /v1/user/login", app.userLoginHandler)

	// Assignment CRUD operations
	router.HandleFunc(
		"POST /v1/course/assignment/create",
		app.assignmentCreateHandler,
	)

	// app.assignmentReadHandler switches its behavior based on the HTTP Method.
	router.HandleFunc(
		"/v1/course/assignment/read/{id}",
		app.assignmentReadHandler,
	)
	//router.HandleFunc(
	//	"POST /v1/course/assignment/read/{id}",
	//	app.assignmentReadHandler,
	//)
	//
	router.HandleFunc(
		"PATCH /v1/course/assignment/update",
		app.assignmentUpdateHandler,
	)
	router.HandleFunc(
		"POST /v1/course/assignment/delete",
		app.assignmentDeleteHandler,
	)
	//router.HandleFunc(
	//	"POST /v1/course/assignment/{id}/upload",
	//	app.assignmentMediaUploadHandler,
	//)
	//router.HandleFunc(
	//	"GET /v1/course/download/{mediaId}",
	//	app.mediaDownloadHandler,
	//)

	// Submission operations
	router.HandleFunc(
		"POST /v1/course/assignment/{assignmentId}/submission/{userId}/create",
		app.submissionCreateHandler,
	)
	router.HandleFunc(
		"POST /v1/course/assignment/submission/{id}/update",
		app.submissionUpdateHandler,
	)
	router.HandleFunc(
		"POST /v1/course/assignment/submission/{id}/delete",
		app.submissionUpdateHandler,
	)
	router.HandleFunc(
		"POST /v1/course/assignment/submission/{id}/read",
		app.submissionUpdateHandler,
	)
	router.HandleFunc(
		"POST /v1/course/assignment/submission/{id}/upload",
		app.submissionMediaUploadHandler,
	)

	return router
}
