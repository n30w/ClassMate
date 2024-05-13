package main

import "net/http"

func (app *application) routes() *http.ServeMux {

	// Enhanced routing patterns in Go 1.22. HandleFuncs now
	// accept a method and a route variable parameter.
	// https://tip.golang.org/doc/go1.22

	router := http.NewServeMux()

	router.HandleFunc("GET /v1/healthcheck", app.healthcheckHandler)
	router.HandleFunc("POST /v1/home", app.homeHandler)
	router.HandleFunc("GET /v1/course/{id}/homepage", app.courseHomepageHandler)

	router.HandleFunc(
		"POST /v1/course/{id}/announcement/create",
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
		"GET /v1/course/{id}/announcement/read",
		app.announcementReadHandler,
	)
	router.HandleFunc("POST /v1/course/addstudent", app.addStudentHandler)

	// Course CRUD operations
	router.HandleFunc("POST /v1/course/create", app.courseCreateHandler)
	router.HandleFunc("GET /v1/course/{id}/read/", app.courseReadHandler)
	router.HandleFunc("POST /v1/course/{id}/delete", app.courseDeleteHandler)

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
	router.HandleFunc(
		"GET /v1/course/{courseId}/assignment/read",
		app.assignmentReadHandler,
	)
	router.HandleFunc(
		"PATCH /v1/course/assignment/update",
		app.assignmentUpdateHandler,
	)
	router.HandleFunc(
		"DELETE /v1/course/assignment/{assignmentId}/delete",
		app.assignmentDeleteHandler,
	)

	// app.assignmentReadHandler switches its behavior based on the HTTP Method.
	router.HandleFunc(
		"/v1/course/{courseId}/assignment/read",
		app.assignmentReadHandler,
	)

	//router.HandleFunc(
	//	"POST /v1/course/assignment/{id}/upload",
	//	app.assignmentMediaUploadHandler,
	//)
	router.HandleFunc(
		"GET /v1/course/{courseId}/download/{mediaId}",
		app.mediaDownloadHandler,
	)

	// Submission operations
	router.HandleFunc(
		"POST /v1/course/assignment/{assignmentId}/submission/create",
		app.submissionCreateHandler,
	)
	router.HandleFunc(
		"POST /v1/course/assignment/submission/{id}/update",
		app.submissionUpdateHandler,
	)
	router.HandleFunc(
		"POST /v1/course/assignment/submission/{id}/delete",
		app.submissionDeleteHandler,
	)
	// Read submission from teacher view
	router.HandleFunc(
		"GET /v1/course/{courseId}/assignment/{assignmentId}/submission/{userId}/read",
		app.teachersubmissionReadHandler,
	)
	// Read submission from student view
	router.HandleFunc(
		"POST /v1/course/{courseId}/assignment/{assignmentId}/submission/read",
		app.studentsubmissionReadHandler,
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
		app.receiveOfflineGrades)
	router.HandleFunc(
		"POST /v1/course/assignment/submission/{id}/upload",
		app.submissionMediaUploadHandler,
	)

	return router
}
