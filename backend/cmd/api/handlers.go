package main

import (
	"net/http"
)

// Course handlers. Any course handler requires a page type from the body
// of the request. This page type must include a page type
// like "home", "course", "section", etc.

// createCourseHandler creates a course.
// Request: course
// Response: course + course UUID
func (app *application) courseCreateHandler(w http.ResponseWriter, r *http.Request) {
}

// courseReadHandler relays information back to the requester
// about a certain course.
// Request: course UUID
// Response: course data
func (app *application) courseReadHandler(w http.ResponseWriter, r *http.Request) {
}

// courseUpdateHandler updates information about a course.
// Request: course UUID + fields to update
// Response: course
func (app *application) courseUpdateHandler(w http.ResponseWriter, r *http.Request) {
}

// courseDeleteHandler deletes a course.
// Request: course UUID
// Response: course list
func (app *application) courseDeleteHandler(w http.ResponseWriter, r *http.Request) {}

// User handlers, deals with anything user side.

// userCreateHandler creates a user.
// Request: email, password, full name
// Response: home page
func (app *application) userCreateHandler(w http.ResponseWriter, r *http.Request) {
}

// userReadHandler reads user data from a requester.
// Request: user UUID
// Response: user
func (app *application) userReadHandler(w http.ResponseWriter, r *http.Request) {
}

// userUpdateHandler updates a user's data.
// Request: user UUID + data to update
// Response: new user data
func (app *application) userUpdateHandler(w http.ResponseWriter, r *http.Request) {
}

// userDeleteHandler deletes a user. A request must come from
// a user themselves to delete themselves.
// Request: user UUID
// Response: logout page
func (app *application) userDeleteHandler(w http.ResponseWriter, r *http.Request) {
}

// userPostHandler handles post requests. When a user posts
// something to a discussion, this is the handler that is called.
// A post consists of a body, media, and author. The request therefore
// requires an author of who posted it, what discussion it exists under,
// and if it is a reply or not. To find the author of who sent it,
// we can check with middleware authorization headers.
// Request: user post
// Response: user post
func (app *application) userPostHandler(w http.ResponseWriter, r *http.Request) {}

// userLoginHandler handles login requests from any user. It requires
// a username and a password. A login must occur from a genuine domain. This
// means that the request comes from the frontend server rather than the
// user's browser. Written to the http response is an authorized
// login cookie.
// Request: username/email, password
// Response: auth cookie/login session
func (app *application) userLoginHandler(w http.ResponseWriter, r *http.Request) {

}

// Assignment handlers. Only teachers should be able to request the use of
// these handlers. Therefore, teacher permission/authorization is
// a necessity.

// assignmentCreateHandler creates an assignment based on the request values. To create an assignment, a request must contain an assignment: title, author, body, and media. The return value is the assignment data along with a UUID.
// Request: title, author, body, media
// Response: assignment
func (app *application) assignmentCreateHandler(w http.ResponseWriter, r *http.Request) {
}

// assignmentReadHandler relays assignment data back to the requester. To read
// one specific assignment, one must only request the UUID of an assignment.
// Request: UUID
// Response: assignment
func (app *application) assignmentReadHandler(w http.ResponseWriter, r *http.Request) {
}

// assignmentUpdateHandler updates the information of an assignment.
// Request: UUID
// Response: assignment
func (app *application) assignmentUpdateHandler(w http.ResponseWriter, r *http.Request) {
}

// assignmentDeleteHandler deletes an assignment.
// Request: UUID
// Response: 200 OK
func (app *application) assignmentDeleteHandler(w http.ResponseWriter, r *http.Request) {
}

// discussionCreateHandler creates a discussion.
// Request: where (the discussion is being created), title, body, media, poster
// Response: discussion data
func (app *application) discussionCreateHandler(w http.ResponseWriter, r *http.Request) {

}

// discussionReadHandler reads a discussion.
// Request: discussion UUID
// Response: discussion
func (app *application) discussionReadHandler(w http.ResponseWriter, r *http.Request) {

}

// discussionUpdateHandler updates a discussion's information. For example,
// the title or body or media and author.
// Request: discussion UUID + information to update
// Response: discussion
func (app *application) discussionUpdateHandler(w http.ResponseWriter, r *http.Request) {

}

// discussionDeleteHandler deletes a discussion.
// Request: discussion UUID
// Response: deleted discussion page
func (app *application) discussionDeleteHandler(w http.ResponseWriter, r *http.Request) {

}
