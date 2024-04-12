package main

import (
	"net/http"

	"github.com/n30w/Darkspace/internal/models"
)

// An input struct is used for ushering in data because it makes it explicit
// as to what we are getting from the incoming request.

// homeHandler returns a set template of information needed for the home
// page.
//
// REQUEST: Netid
// RESPONSE: Active course data [name, 3 most recent assignments uncompleted, ]
func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	// Get user's enrolled courses
}

// courseHomepageHandler returns data related to the homepage of a course.
//
// REQUEST: course id
// RESPONSE:
func (app *application) courseHomepageHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	id := r.PathValue("id")

	var course *models.Course

	course, err = app.services.CourseService.RetrieveCourse(id)
	if err != nil {
		app.serverError(w, r, err)
	}

	res := jsonWrap{"course": course}

	err = app.writeJSON(w, http.StatusOK, res, nil)
	if err != nil {
		app.serverError(w, r, err)
	}

	// If the course ID exists in the database AND the user requesting this
	// data has the appropriate permissions, retrieve the course data requested.
}

// createCourseHandler creates a course.
//
// REQUEST: course title, username
// RESPONSE: course id, name, teacher, assignments
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

	// Might need to reconsider how we store teachers in course model, currently by user struct
	teacher, err := app.services.UserService.GetByUsername(input.TeacherName)

	course.Name = input.Title
	course.Teachers = append(course.Teachers, teacher)

	err = app.services.CourseService.CreateCourse(course)
	if err != nil {
		app.serverError(w, r, err)
	}

	// Return success.
	res := jsonWrap{"course": course}
	err = app.writeJSON(w, http.StatusOK, res, nil)
	if err != nil {
		app.serverError(w, r, err)
	}
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
	var input struct {
		Id string `json:"id"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverError(w, r, err)
	}

	course, err := app.services.CourseService.RetrieveCourse(input.Id)
	if err != nil {
		app.serverError(w, r, err)
	}

	res := jsonWrap{"course": course}
	err = app.writeJSON(w, http.StatusOK, res, nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}

// courseUpdateHandler updates information about a course.
// REQUEST: course ID + fields to update (add user, delete user, rename)
// RESPONSE: course
func (app *application) courseUpdateHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	courseid := r.PathValue("id")
	action := r.PathValue("action")

	switch action {

	case "add", "delete":
		var input struct {
			UserId string `json:"userid"`
		}
		err := app.readJSON(w, r, &input)
		if err != nil {
			app.serverError(w, r, err)
		}
		if action == "add" {
			course, err := app.services.CourseService.AddToRoster(courseid, input.UserId)
			if err != nil {
				app.serverError(w, r, err)
			}
			res := jsonWrap{"course": course}
			err = app.writeJSON(w, http.StatusOK, res, nil)
			if err != nil {
				app.serverError(w, r, err)
			}
		} else if action == "delete" {
			course, err := app.services.CourseService.RemoveFromRoster(courseid, input.UserId)
			if err != nil {
				app.serverError(w, r, err)
			}
			res := jsonWrap{"course": course}
			err = app.writeJSON(w, http.StatusOK, res, nil)
			if err != nil {
				app.serverError(w, r, err)
			}
		}

	case "rename":
		var input struct {
			Name string
		}
		err := app.readJSON(w, r, &input)
		if err != nil {
			app.serverError(w, r, err)
		}

		course, err := app.services.CourseService.UpdateCourseName(courseid, input.Name)
		if err != nil {
			app.serverError(w, r, err)
			return
		}
		res := jsonWrap{"course": course}
		err = app.writeJSON(w, http.StatusOK, res, nil)
		if err != nil {
			app.serverError(w, r, err)
		}

	default:
		app.serverError(w, r, err) //need to format error, input field is not one of the 3 options
	}

}

// courseDeleteHandler deletes a course
//
// REQUEST: course ID, user id
// RESPONSE: updated list of course
func (app *application) courseDeleteHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	var input struct {
		CourseId string `json:"courseid"`
		UserId   string `json:"userid"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = app.services.CourseService.DeleteCourse(input.CourseId, input.UserId) // needs work
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	// RetrieveUserCourse requires overlapping store functions
	// Need to implement specific field retrieval for Users, i.e. retrieving courses of a user
	courses, err := app.services.UserService.RetrieveFromUser(input.UserId, "courses")
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	res := jsonWrap{"course": courses}
	err = app.writeJSON(w, http.StatusOK, res, nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}

// REQUEST: course ID, teacher ID, announcement description
// RESPONSE: announcement
func (app *application) announcementCreateHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	var input struct {
		CourseId     string           `json:"courseid"`
		TeacherId    string           `json:"teacherid"`
		Announcement string           `json:"announcement"`
		Media        []models.MediaId `json:"media"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	var anno *models.Announcement
	var msg *models.Message
	msg.Post.Description, msg.Post.Owner, msg.Post.Media = input.Announcement, input.TeacherId, input.Media

	msg, err = app.services.MessageService.CreateMessage(msg)
	anno.Message = *msg

}

func (app *application) announcementUpdateHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

}

func (app *application) announcementDeleteHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

}

// User handlers, deals with anything user side.

// userCreateHandler creates a user.
//
// REQUEST: email, password, full name, netid, membership
// RESPONSE: home page
func (app *application) userCreateHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	var input struct {
		Username   string `json:"username"`
		Password   string `json:"password"`
		Email      string `json:"email"`
		Netid      string `json:"netid"`
		Membership int    `json:"membership"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverError(w, r, err)
	}

	// Map the input fields to the appropriate credentials fields.
	c := models.Credentials{
		Username:   app.services.UserService.NewUsername(input.Username),
		Password:   app.services.UserService.NewPassword(input.Password),
		Email:      app.services.UserService.NewEmail(input.Email),
		Membership: app.services.UserService.NewMembership(input.Membership),
	}

	user := models.NewUser(input.Netid, c)

	err = app.services.UserService.CreateUser(user)
	if err != nil {
		app.serverError(w, r, err)
	}

	// Here we would generate a session token, but not now.

	// Send back home page.

}

// userReadHandler reads a specific user's data,
// which is specified by the requester.
//
// REQUEST: user netid
// RESPONSE: user
func (app *application) userReadHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	id := r.PathValue("id")

	var err error
	var user *models.User

	// Perform a database lookup of user.
	user, err = app.services.UserService.GetByID(id)
	if err != nil {
		app.serverError(w, r, err)
	}

	res := jsonWrap{"user": user}

	err = app.writeJSON(w, http.StatusOK, res, nil)
	if err != nil {
		app.serverError(w, r, err)
	}

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
