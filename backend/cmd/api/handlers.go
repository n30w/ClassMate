package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/n30w/Darkspace/internal/models"
)

func (app *application) downloadExcelHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		CourseId string `json:"courseid"`
	}
	// Get the Excel file with the user input data
	file, err := app.services.ExcelService.CreateExcel(input.CourseId)
	if err != nil {
		app.serverError(w, r, err)
	}

	// Set the headers necessary to get browsers to interpret the downloadable file
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, file.GetFileName()))
	w.Header().Set("File-Name", fmt.Sprintf("%s"))
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")
	err = file.Write(w)
}

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

	course, err := app.services.CourseService.RetrieveCourse(id)
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
// REQUEST: course title, user id
// RESPONSE: course id, name, teacher, assignments
func (app *application) courseCreateHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	var input struct {
		Title     string `json:"title"`
		TeacherID string `json:"teacherid"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverError(w, r, err)
	}
	teachers := []string{input.TeacherID}

	course := &models.Course{
		Title:    input.Title,
		Teachers: teachers,
	}

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
		CourseId string `json:"courseid"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverError(w, r, err)
	}

	course, err := app.services.CourseService.RetrieveCourse(input.CourseId)
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
	id := r.PathValue("id")

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
			course, err := app.services.CourseService.AddToRoster(id, input.UserId)
			if err != nil {
				app.serverError(w, r, err)
			}
			res := jsonWrap{"course": course}
			err = app.writeJSON(w, http.StatusOK, res, nil)
			if err != nil {
				app.serverError(w, r, err)
			}
		} else if action == "delete" {
			course, err := app.services.CourseService.RemoveFromRoster(id, input.UserId)
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

		course, err := app.services.CourseService.UpdateCourseName(id, input.Name)
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
		app.serverError(w, r, fmt.Errorf("%s is an invalid action", action)) //need to format error, input field is not one of the 3 options
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

	err = app.services.UserService.UnenrollUserFromCourse(input.UserId, input.CourseId) // delete course from user
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	_, err = app.services.CourseService.RemoveFromRoster(input.CourseId, input.UserId) // delete user from course
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	courses, err := app.services.UserService.RetrieveFromUser(input.UserId, "courses")
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	res := jsonWrap{"courses": courses}
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
		CourseId    string   `json:"courseid"`
		TeacherId   string   `json:"teacherid"`
		Title       string   `json:"title"`
		Date        string   `json:"date"`
		Description string   `json:"description"`
		Media       []string `json:"media"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	post := models.Post{
		Title:       input.Title,
		Description: input.Description,
		Owner:       input.TeacherId,
		Media:       input.Media,
		Date:        input.Date,
	}
	msg := &models.Message{
		Post: post,
		Type: 1,
	}

	msg, err = app.services.MessageService.CreateMessage(msg, input.CourseId)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	res := jsonWrap{"announcement": msg}
	err = app.writeJSON(w, http.StatusOK, res, nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}

// REQUEST: announcement ID
// RESPONSE: announcement
func (app *application) announcementReadHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	var input struct {
		MsgId string `json:"announcementid"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	msg, err := app.services.MessageService.ReadMessage(input.MsgId)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	res := jsonWrap{"announcement": msg}
	err = app.writeJSON(w, http.StatusOK, res, nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}

// REQUEST: announcement ID, action (title, body), updated field
// RESPONSE: announcement
func (app *application) announcementUpdateHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	var input struct {
		MsgId        string `json:"announcementid"`
		Action       string `json:"action"`
		UpdatedField string `json:"updatedfield"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	msg, err := app.services.MessageService.UpdateMessage(input.MsgId, input.Action, input.UpdatedField)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	res := jsonWrap{"announcement": msg}
	err = app.writeJSON(w, http.StatusOK, res, nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) announcementDeleteHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	var input struct {
		CourseId  string `json:"courseid"`
		TeacherId string `json:"teacherid"`
		MsgId     string `json:"announcementid"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = app.services.MessageService.DeleteMessage(input.MsgId)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	res := jsonWrap{"announcement": nil}
	err = app.writeJSON(w, http.StatusOK, res, nil)
	if err != nil {
		app.serverError(w, r, err)
	}
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

	// Perform a database lookup of user.
	user, err := app.services.UserService.GetByID(id)
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
	// id := r.PathValue("id")

	// var input struct {

	// }

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
	var input struct {
		Title       string    `json:"title"`
		TeacherId   string    `json:"teacherid"`
		Description string    `json:"description"`
		Media       []string  `json:"media"`
		DueDate     time.Time `json:"time"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverError(w, r, err)
	}

	post := models.Post{
		Title:       input.Title,
		Description: input.Description,
		Owner:       input.TeacherId,
		Media:       input.Media,
	}
	assignment := &models.Assignment{
		Post:    post,
		DueDate: input.DueDate,
	}

	assignment, err = app.services.AssignmentService.CreateAssignment(assignment)
	if err != nil {
		app.serverError(w, r, err)
	}

	res := jsonWrap{"assignment": assignment}

	err = app.writeJSON(w, http.StatusOK, res, nil)
	if err != nil {
		app.serverError(w, r, err)
	}

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
	var input struct {
		Uuid string `json:"uuid"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverError(w, r, err)
	}

	assignment, err := app.services.AssignmentService.ReadAssignment(input.Uuid)
	if err != nil {
		app.serverError(w, r, err)
	}

	res := jsonWrap{"assignment": assignment}

	err = app.writeJSON(w, http.StatusOK, res, nil)
	if err != nil {
		app.serverError(w, r, err)
	}

}

// assignmentUpdateHandler updates the information of an assignment.
//
// REQUEST: uuid, updated information, type (title, description, duedate)
// RESPONSE: assignment
func (app *application) assignmentUpdateHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	var input struct {
		Uuid         string      `json:"uuid"`
		UpdatedField interface{} `json:"updatedfield"`
		Action       string      `json:"action"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverError(w, r, err)
	}

	assignment, err := app.services.AssignmentService.UpdateAssignment(input.Uuid, input.UpdatedField, input.Action)
	if err != nil {
		app.serverError(w, r, err)
	}

	res := jsonWrap{"assignment": assignment}

	err = app.writeJSON(w, http.StatusOK, res, nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}

// assignmentDeleteHandler deletes an assignment.
//
// REQUEST: uuid
// RESPONSE: 200 OK
func (app *application) assignmentDeleteHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	var input struct {
		Uuid string `json:"uuid"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverError(w, r, err)
	}

	err = app.services.AssignmentService.DeleteAssignment(input.Uuid)
	if err != nil {
		app.serverError(w, r, err)
	}

	err = app.writeJSON(w, http.StatusOK, nil, nil)
	if err != nil {
		app.serverError(w, r, err)
	}

}

// discussionCreateHandler creates a discussion.
//
// REQUEST: where (the discussion is being created), title, body, media, poster
// RESPONSE: discussion data
func (app *application) discussionCreateHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	var input struct {
		CourseId   string   `json:"courseid"`
		PosterId   string   `json:"posterid"`
		Discussion string   `json:"discussion"`
		Media      []string `json:"media"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	post := models.Post{
		Description: input.Discussion,
		Owner:       input.PosterId,
		Media:       input.Media,
	}
	msg := &models.Message{
		Post: post,
		Type: 0,
	}

	msg, err = app.services.MessageService.CreateMessage(msg, input.CourseId)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	res := jsonWrap{"discussion": msg}
	err = app.writeJSON(w, http.StatusOK, res, nil)
	if err != nil {
		app.serverError(w, r, err)
	}
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

// Media handlers
func (app *application) mediaCreateHandler(w http.ResponseWriter,
	r *http.Request,
) {

}
func (app *application) mediaDeleteHandler(w http.ResponseWriter,
	r *http.Request,
) {

}

// Comment handlers
//
// REQUEST: discussion/announcement uuid + comment + author netid
// RESPONSE: comment
func (app *application) commentCreateHandler(w http.ResponseWriter,
	r *http.Request,
) {
	var input struct {
		MessageId string `json:"messageid"`
		Comment   string `json:"comment"`
		Netid     string `json:"netid"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverError(w, r, err)
	}

}
func (app *application) commentDeleteHandler(w http.ResponseWriter,
	r *http.Request) {
	var input struct {
		Uuid  string `json:"uuid"`
		Netid string `json:"netid"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverError(w, r, err)
	}
}

// Submission handlers
//
// REQUEST: assignmentid + userid + filetype + submissiontime
// RESPONSE: submission
func (app *application) submissionCreateHandler(w http.ResponseWriter,
	r *http.Request) {

	r.ParseMultipartForm(10 << 20)
	file, header, err := r.FormFile("file")
	if err != nil {
		app.serverError(w, r, err)
	}

	defer file.Close()
	media := &models.Media{
		FileName:           header.Filename,
		CreatedAt:          r.FormValue("submissiontime"),
		AttributionsByType: make(map[string]string),
		FileType:           r.FormValue("filetype"),
	}

	media.AttributionsByType["assignment"] = r.FormValue("assignmentid")
	media.AttributionsByType["user"] = r.FormValue("userid")

	submission := &models.Submission{
		AssignmentId:   r.FormValue("assignmentid"),
		UserId:         r.FormValue("userid"),
		SubmissionTime: r.FormValue("submissiontime"),
		Media:          media,
	}

	submission, err = app.services.SubmissionService.CreateSubmission(submission)
	if err != nil {
		app.serverError(w, r, err)
	}
	media, err = app.services.MediaService.UploadMedia(file, submission.ID) // implement cloud storage of file and add reference to submission ID, return media struct (metadata)
	if err != nil {
		app.serverError(w, r, err)
	}

	_, err = app.services.AssignmentService.UpdateAssignment(submission.AssignmentId, true, "submit") // assignment is now completed
	if err != nil {
		app.serverError(w, r, err)
	}

	res := jsonWrap{"submission": submission}
	err = app.writeJSON(w, http.StatusOK, res, nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}
