package main

import (
	"net/http"

	models "github.com/n30w/Darkspace/internal/domain"
)

// homeHandler returns a set template of information needed for the home
// page.
//
// REQUEST: Netid
// RESPONSE: Active course data [name, 3 most recent assignments uncompleted, ]
func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {

}

// courseHomepageHandler returns data related to the homepage of a course.
//
// REQUEST: course id
// RESPONSE:
func (app *application) courseHomepageHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	// id := r.PathValue("id")

	// var course *models.Course
	// var err error

	// course, err = app.models.Course.Get(id)
	// if err != nil {
	// 	app.serverError(w, r, err)
	// }

	// res := jsonWrap{"course": course}

	// err = app.writeJSON(w, http.StatusOK, res, nil)
	// if err != nil {
	// 	app.serverError(w, r, err)
	// }

	// If the course ID exists in the database AND the user requesting this
	// data has the appropriate permissions, retrieve the course data requested.

}

// createCourseHandler creates a course.
//
// REQUEST: course
// RESPONSE: course + course uuid
func (app *application) courseCreateHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	var input struct {
		Title       string `json:"title"`
		TeacherName string `json:"username"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverError(w, r, err)
	}

	var course *models.Course

	// Validate if there is a name associated with the course.
	// We can also do a fuzzy match of course names.
	course, err = app.models.Course.Get(input.Title)

	// If course already exists, send error.
	if err != nil {
		app.serverError(w, r, err)
	}

	// if not, proceed with course creation.
	err = app.models.Course.Insert(course)
	if err != nil {
		app.serverError(w, r, err)
	}
	// Return success.
}

// courseReadHandler relays information back to the requester
// about a certain course.
//
// REQUEST: course ID
// RESPONSE: course data
func (app *application) courseReadHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
}

// courseUpdateHandler updates information about a course.
//
// REQUEST: course ID + fields to update
// RESPONSE: course
func (app *application) courseUpdateHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
}

// courseDeleteHandler deletes a course.
//
// REQUEST: course ID.
// RESPONSE: updated list of courses
func (app *application) courseDeleteHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
}

// User handlers, deals with anything user side.

// userCreateHandler creates a user.
//
// REQUEST: email, password, full name, netid
// RESPONSE: home page
func (app *application) userCreateHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	// use credential validation
}

// userReadHandler reads a specific user's data,
// which is specified by the requester.
//
// REQUEST: user uuid
// RESPONSE: user
func (app *application) userReadHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Create a new user model, this will be used to perform SQL actions.

	// Create a new User to read the JSON into.
	u := models.User{}

	// read the JSON from the client into the User Model.
	err := app.readJSON(w, r, &u)
	if err != nil {
		app.serverError(w, r, err)
	}

	// Perform a database lookup of user.

	// If user exists, read permissions of the user using accessControl,
	// then send back the information requested.
	//if u.perms[SELF].read {
	//	// Send back the requested user's bio/media/name/email,
	//	// by first wrapping the JSON into a readable map,
	//	// then writing to the http writer stream.
	//	d := jsonWrap{}
	//	err := app.writeJSON(w, http.StatusOK, d, nil)
	//} else {
	//	app.serverError(w, r, http.ErrHandlerTimeout)
	//}
}

// userUpdateHandler updates a user's data.
//
// REQUEST: user UUID + data to update
// RESPONSE: new user data
func (app *application) userUpdateHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
}

// userDeleteHandler deletes a user. A request must come from
// a user themselves to delete themselves.
//
// REQUEST: user uuid
// RESPONSE: logout page
func (app *application) userDeleteHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
}

// userPostHandler handles post requests. When a user posts
// something to a discussion, this is the handler that is called.
// A post consists of a body, media, and author. The request therefore
// requires an author of who posted it, what discussion it exists under,
// and if it is a reply or not. To find the author of who sent it,
// we can check with middleware authorization headers.
//
// REQUEST: user post
// RESPONSE: user post
func (app *application) userPostHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
}

// userLoginHandler handles login requests from any user. It requires
// a username and a password. A login must occur from a genuine domain. This
// means that the request comes from the frontend server rather than the
// user's browser. Written to the http response is an authorized
// login cookie.
//
// REQUEST: username/email, password
// RESPONSE: auth cookie/login session
func (app *application) userLoginHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

}

// Assignment handlers. Only teachers should be able to request the use of
// these handlers. Therefore, teacher permission/authorization is
// a necessity.

// assignmentCreateHandler creates an assignment based on the request values.
// To create an assignment, a request must contain an assignment: title,
// author, body, and media. The return value is the assignment data along
// with a uuid.
//
// REQUEST: title, author, body, media
// RESPONSE: assignment
func (app *application) assignmentCreateHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
}

// assignmentReadHandler relays assignment data back to the requester. To read
// one specific assignment, one must only request the UUID of an assignment.
//
// REQUEST: uuid
// RESPONSE: assignment
func (app *application) assignmentReadHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
}

// assignmentUpdateHandler updates the information of an assignment.
//
// REQUEST: uuid
// RESPONSE: assignment
func (app *application) assignmentUpdateHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
}

// assignmentDeleteHandler deletes an assignment.
//
// REQUEST: uuid
// RESPONSE: 200 OK
func (app *application) assignmentDeleteHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
}

// discussionCreateHandler creates a discussion.
//
// REQUEST: where (the discussion is being created), title, body, media, poster
// RESPONSE: discussion data
func (app *application) discussionCreateHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

}

// discussionReadHandler reads a discussion.
//
// REQUEST: discussion uuid
// RESPONSE: discussion
func (app *application) discussionReadHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

}

// discussionUpdateHandler updates a discussion's information. For example,
// the title or body or media and author.
//
// REQUEST: discussion uuid + information to update
// RESPONSE: discussion
func (app *application) discussionUpdateHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

}

// discussionDeleteHandler deletes a discussion.
//
// REQUEST: discussion uuid
// RESPONSE: 200 or 500 response
func (app *application) discussionDeleteHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

}
