package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/n30w/Darkspace/internal/models"
	"github.com/xuri/excelize/v2"
)

func (app *application) downloadExcelHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	var input struct {
		CourseId string `json:"courseid"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Get the Excel file with the user input data
	file, err := app.services.ExcelService.CreateExcel(input.CourseId)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Set the headers necessary to get browsers to interpret the downloadable file
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set(
		"Content-Disposition",
		fmt.Sprintf(`attachment; filename="%s"`, "todo"),
	) // TODO FIX ME
	w.Header().Set("File-Name", fmt.Sprintf("%s"))
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")
	err = file.Write(w)
	file.Close()
}

func (app *application) uploadExcelHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	r.ParseMultipartForm(10 << 20)
	file, _, err := r.FormFile("excelfile")
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	defer file.Close()
	// Read the Excel file content
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	// Create a temporary file to store the uploaded Excel file
	tempFile, err := os.CreateTemp("temp-excel", "upload-*.xlsx")
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	// Write the Excel file content to the temporary file
	_, err = tempFile.Write(fileBytes)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Open the Excel file using excelize
	excelFile, err := excelize.OpenFile(tempFile.Name())
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = app.services.ExcelService.ParseExcel(excelFile)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, nil, nil)
	if err != nil {
		app.serverError(w, r, err)
	}
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
	var input struct {
		Token string `json:"token"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	netid, err := app.services.AuthenticationService.GetNetIdFromToken(input.Token)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	courses, err := app.services.UserService.RetrieveFromUser(netid, "Courses")
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

// courseHomepageHandler returns data related to the homepage of a course.
//
// REQUEST: course id
// RESPONSE:
func (app *application) courseHomepageHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	id := r.PathValue("id")

	course, err := app.services.CourseService.RetrieveCourse(id)
	if err != nil {
		app.serverError(w, r, err)
	}

	res := jsonWrap{"course_info": course}

	err = app.writeJSON(w, http.StatusOK, res, nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}

// createCourseHandler creates a course.
func (app *application) courseCreateHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	var input struct {
		Title     string `json:"title"`
		Token     string `json:"token"`
		BannerUrl string `json:"banner"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	teacherid, err := app.services.AuthenticationService.GetNetIdFromToken(input.Token)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	teachers := []string{teacherid}

	course := &models.Course{
		Title:    input.Title,
		Teachers: teachers,
	}

	course, err = app.services.CourseService.CreateCourse(course, teacherid, input.BannerUrl)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Return success.
	res := jsonWrap{"course": course}
	err = app.writeJSON(w, http.StatusOK, res, nil)
	if err != nil {
		app.serverError(w, r, err)
		return
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
		return
	}

	course, err := app.services.CourseService.RetrieveCourse(input.CourseId)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	res := jsonWrap{"course": course}
	err = app.writeJSON(w, http.StatusOK, res, nil)
	if err != nil {
		app.serverError(w, r, err)
		return
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
			return
		}
		if action == "add" {
			course, err := app.services.CourseService.AddToRoster(
				id,
				input.UserId,
			)
			if err != nil {
				app.serverError(w, r, err)
				return
			}
			res := jsonWrap{"course": course}
			err = app.writeJSON(w, http.StatusOK, res, nil)
			if err != nil {
				app.serverError(w, r, err)
				return
			}
		} else if action == "delete" {
			course, err := app.services.CourseService.RemoveFromRoster(
				id,
				input.UserId,
			)
			if err != nil {
				app.serverError(w, r, err)
				return
			}
			res := jsonWrap{"course": course}
			err = app.writeJSON(w, http.StatusOK, res, nil)
			if err != nil {
				app.serverError(w, r, err)
				return
			}
		}

	case "rename":
		var input struct {
			Name string
		}
		err := app.readJSON(w, r, &input)
		if err != nil {
			app.serverError(w, r, err)
			return
		}

		course, err := app.services.CourseService.UpdateCourseName(
			id,
			input.Name,
		)
		if err != nil {
			app.serverError(w, r, err)
			return
		}
		res := jsonWrap{"course": course}
		err = app.writeJSON(w, http.StatusOK, res, nil)
		if err != nil {
			app.serverError(w, r, err)
			return
		}

	default:
		app.serverError(
			w,
			r,
			fmt.Errorf("%s is an invalid action", action),
		) //need to format error, input field is not one of the 3 options
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

	err = app.services.UserService.UnenrollUserFromCourse(
		input.UserId,
		input.CourseId,
	) // delete course from user
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	_, err = app.services.CourseService.RemoveFromRoster(
		input.CourseId,
		input.UserId,
	) // delete user from course
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	courses, err := app.services.UserService.RetrieveFromUser(
		input.UserId,
		"courses",
	)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	res := jsonWrap{"courses": courses}
	err = app.writeJSON(w, http.StatusOK, res, nil)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

// REQUEST: course ID, teacher ID, announcement description
// RESPONSE: announcement
func (app *application) announcementCreateHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	cId := r.PathValue("id")
	var input struct {
		CourseId    string   `json:"courseid"`
		Token       string   `json:"token"`
		Title       string   `json:"title"`
		Description string   `json:"description"`
		Media       []string `json:"media"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	netid, err := app.services.AuthenticationService.GetNetIdFromToken(input.Token)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	msg := &models.Message{
		Post: models.Post{
			Title:       input.Title,
			Description: input.Description,
			Owner:       netid,
			Media:       input.Media,
		},
		Type: 1,
	}
	msg, err = app.services.MessageService.CreateMessage(msg, cId)

	if err != nil {
		app.serverError(w, r, err)
		return
	}
	res := jsonWrap{"announcement": msg}
	err = app.writeJSON(w, http.StatusOK, res, nil)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

// REQUEST: announcement ID
// RESPONSE: announcement
func (app *application) announcementReadHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	courseId := r.PathValue("id")
	msgids, err := app.services.MessageService.RetrieveMessages(courseId)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	var msgs []models.Message
	for _, msgid := range msgids {
		msg, err := app.services.MessageService.ReadMessage(msgid)
		if err != nil {
			app.serverError(w, r, err)
			return
		}
		msgs = append(msgs, *msg)
	}

	res := jsonWrap{"announcements": msgs}
	err = app.writeJSON(w, http.StatusOK, res, nil)
	if err != nil {
		app.serverError(w, r, err)
		return
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

	msg, err := app.services.MessageService.UpdateMessage(
		input.MsgId,
		input.Action,
		input.UpdatedField,
	)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	res := jsonWrap{"announcement": msg}
	err = app.writeJSON(w, http.StatusOK, res, nil)
	if err != nil {
		app.serverError(w, r, err)
		return
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
		return
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
		FullName   string `json:"fullname"`
		Password   string `json:"password"`
		Email      string `json:"email"`
		Netid      string `json:"netid"`
		Membership int    `json:"membership"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Map the input fields to the appropriate credentials fields.
	c := models.Credentials{
		Username:   app.services.UserService.NewUsername(input.Netid),
		Password:   app.services.UserService.NewPassword(input.Password),
		Email:      app.services.UserService.NewEmail(input.Email),
		Membership: app.services.UserService.NewMembership(input.Membership),
	}

	user, err := models.NewUser(input.Netid, c, input.FullName)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = app.services.UserService.CreateUser(user)
	if err != nil {
		app.serverError(w, r, err)
		return
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
		return
	}

	res := jsonWrap{"user": user}

	err = app.writeJSON(w, http.StatusOK, res, nil)
	if err != nil {
		app.serverError(w, r, err)
		return
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
	id := r.PathValue("id")
	fmt.Printf("id: %s", id)
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
// login token.
//
// REQUEST: username/email, password
// RESPONSE: auth cookie/login session
func (app *application) userLoginHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	var input struct {
		NetId    string `json:"netid"`
		Password string `json:"password"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	// 	Validate credentials
	err = app.services.UserService.ValidateUser(input.NetId, input.Password)
	if err != nil {
		app.serverError(w, r, err)
		err = app.writeJSON(
			w, http.StatusCreated,
			fmt.Sprint("wrong login information"), nil,
		)
		return
	}
	// If token exists, return token:
	token, err := app.services.AuthenticationService.RetrieveToken(input.NetId)
	if err != errors.New("record not found") {
		membership, err2 := app.services.UserService.GetMembership(input.NetId)
		if err2 != nil {
			app.serverError(w, r, err2)
			return
		}
		wrapped := jsonWrap{"authentication_token": token, "permissions": membership}
		err = app.writeJSON(
			w, http.StatusCreated,
			wrapped, nil,
		)
		if err != nil {
			app.serverError(w, r, err)
			return
		}
		return
	}
	// Otherwise, generate new token
	token, err = app.services.AuthenticationService.NewToken(input.NetId)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	membership, err2 := app.services.UserService.GetMembership(input.NetId)
	if err2 != nil {
		app.serverError(w, r, err2)
		return
	}
	wrapped := jsonWrap{"authentication_token": token, "permissions": membership}
	err = app.writeJSON(
		w, http.StatusCreated,
		wrapped, nil,
	)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

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
	fmt.Printf("AssignmentCreateHandler")
	var input struct {
		Title       string   `json:"title"`
		Token       string   `json:"token"`
		Description string   `json:"description"`
		Media       []string `json:"media"`
		DueDate     string   `json:"duedate"`
		CourseId  string `json:"courseid"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	netid, err := app.services.AuthenticationService.GetNetIdFromToken(input.Token)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	fmt.Printf("Assignment Authenticating token")

	post := models.Post{
		Title:       input.Title,
		Description: input.Description,
		Owner:       netid,
		Media:       input.Media,
		Course: input.CourseId,
	}

	dueDate, err := time.Parse("2006-01-02", input.DueDate)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	assignment := &models.Assignment{
		Post:    post,
		DueDate: dueDate,
	}

	assignment, err = app.services.AssignmentService.CreateAssignment(assignment)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	res := jsonWrap{"assignment": assignment}

	err = app.writeJSON(w, http.StatusOK, res, nil)
	if err != nil {
		app.serverError(w, r, err)
		return
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

	courseid := r.PathValue("id")
	fmt.Printf("courseid: %s", courseid)
	assignmentids, err := app.services.AssignmentService.RetrieveAssignments(courseid)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	fmt.Printf("Assignmentids: %v", assignmentids)
	var assignments []models.Assignment
	for _, id := range assignmentids {
		assignment, err := app.services.AssignmentService.ReadAssignment(id)
		if err != nil {
			app.serverError(w, r, err)
			return
		}
		assignments = append(assignments, *assignment)
	}
	fmt.Printf("Assignments: %v", assignments)

	res := jsonWrap{"assignment": assignments}

	err = app.writeJSON(w, http.StatusOK, res, nil)
	if err != nil {
		app.serverError(w, r, err)
		return
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
		return
	}

	assignment, err := app.services.AssignmentService.UpdateAssignment(
		input.Uuid,
		input.UpdatedField,
		input.Action,
	)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	res := jsonWrap{"assignment": assignment}

	err = app.writeJSON(w, http.StatusOK, res, nil)
	if err != nil {
		app.serverError(w, r, err)
		return
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
		return
	}

	err = app.services.AssignmentService.DeleteAssignment(input.Uuid)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, nil, nil)
	if err != nil {
		app.serverError(w, r, err)
		return
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
		return
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
func (app *application) mediaCreateHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

}
func (app *application) mediaDeleteHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

}

// Comment handlers
//
// REQUEST: discussion/announcement uuid + comment + author netid
// RESPONSE: comment
func (app *application) commentCreateHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	var input struct {
		Uuid    string `json:"uuid"`
		Comment string `json:"comment"`
		Netid   string `json:"netid"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

}
func (app *application) commentDeleteHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	var input struct {
		Uuid  string `json:"uuid"`
		Netid string `json:"netid"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

// Submission handlers
//
// REQUEST: assignmentid + userid + filetype
// RESPONSE: submission
func (app *application) submissionCreateHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	r.ParseMultipartForm(10 << 20)
	file, header, err := r.FormFile("file")
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	fileTypeStr := r.FormValue("filetype") // file type of string to type FileType
	fileTypeInt, err := strconv.Atoi(fileTypeStr)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	filetype := models.FileType(fileTypeInt)
	defer file.Close()

	metadata := &models.Media{
		FileName:           header.Filename,
		AttributionsByType: make(map[string]string),
		FileType:           filetype,
	}

	metadata.AttributionsByType["assignment"] = r.FormValue("assignmentid") // Set media attributions
	metadata.AttributionsByType["user"] = r.FormValue("userid")

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	submission := &models.Submission{
		AssignmentId: r.FormValue("assignmentid"),
		User: models.User{
			Entity: models.Entity{
				ID: r.FormValue("userid"),
			},
		},
	}
	submission, err = app.services.SubmissionService.CreateSubmission(submission) // Add submission into database and return model with ID
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	metadata.AttributionsByType["submission"] = submission.ID // Set media submission attribution with new submission ID

	metadata, err = app.services.MediaService.UploadMedia(file, metadata) // implement cloud storage of file and add reference to submission ID, return media struct (metadata)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	_, err = app.services.AssignmentService.UpdateAssignment( // Submit assignment
		submission.AssignmentId,
		true,
		"submit",
	) // assignment is now completed
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	res := jsonWrap{"submission": submission} // Return submission
	err = app.writeJSON(w, http.StatusOK, res, nil)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

// SubmissionUpdateHandler handles multiple submisisons
// REQUEST: submission id + action()
func (app *application) submissionUpdateHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

}

func (app *application) addStudentHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	var input struct {
		NetId    string `json:"netid"`
		CourseId string `json:"courseid"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	user, err := app.services.UserService.GetByID(input.NetId)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	for _, course := range user.Courses {
		if course == input.CourseId {
			res := jsonWrap{"response": "User is already enrolled"}
			err = app.writeJSON(w, http.StatusOK, res, nil)
			if err != nil {
				app.serverError(w, r, err)
				return
			}
		}
	}
	// User is not enrolled in the course
	_, err = app.services.CourseService.AddToRoster(input.CourseId, input.NetId)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	res := jsonWrap{"response": "user successfully added to course"}
	err = app.writeJSON(w, http.StatusOK, res, nil)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

// func (app *application) courseImageHandler(
// 	w http.ResponseWriter,
// 	r *http.Request,
// ) {
// 	f := app.services.MediaService.
// 	buf, err := os.ReadFile("sid.png")

// 	if err != nil {

// 		log.Fatal(err)
// 	}

// 	w.Header().Set("Content-Type", "image/png")
// 	w.Header().Set("Content-Disposition", `attachment;filename="sid.png"`)

// 	w.Write(buf)
// }
