package main

import (
	"errors"
	"fmt"
	"mime"
	"net/http"

	"path/filepath"
	"time"

	"github.com/n30w/Darkspace/internal/dal"
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
	var input struct {
		Token string `json:"token"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.logger.Printf("token received: %s", input.Token)

	netId, err := app.services.AuthenticationService.GetNetIdFromToken(input.Token)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.logger.Printf("retrieved Net ID: %s", netId)

	courses, err := app.services.UserService.GetUserCourses(netId)
	if err != nil {
		app.logger.Printf("ERROR: %v", err)
		app.serverError(w, r, err)
		return
	}

	app.logger.Printf("courses retrieved: %d", len(courses))
	app.logger.Printf("courses retrieved: %v", courses)

	res := jsonWrap{"courses": courses}

	err = app.writeJSON(w, http.StatusOK, res, nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}

// courseHomepageHandler returns data related to the homepage of a course.
//
// REQUEST: course id
// RESPONSE: course + banner image
func (app *application) courseHomepageHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	id := r.PathValue("id")

	course, err := app.services.CourseService.RetrieveCourse(id)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Get the first couple people in the roster.
	roster, err := app.services.CourseService.RetrieveRoster(id)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	res := jsonWrap{"course": course, "roster": roster}

	app.logger.Printf("sending response: %+v", res)

	err = app.writeJSON(w, http.StatusOK, res, nil)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

// createCourseHandler creates a course.
func (app *application) courseCreateHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	app.logger.Printf("Creating course...")
	var input struct {
		Title string `json:"title"`
		Token string `json:"token"`
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
	app.logger.Printf("Getting teacher id from token %s", teacherid)

	teachers := []string{teacherid}

	course := &models.Course{
		Title:    input.Title,
		Teachers: teachers,
	}

	course, err = app.services.CourseService.CreateCourse(course, teacherid)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	app.logger.Printf("Created course: %v", course)

	// Return success.
	res := jsonWrap{"course": course}
	err = app.writeJSON(w, http.StatusOK, res, nil)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

// courseReadHandler relays information back to the requester
// about a certain course. This is hit when the frontend requests
// the homepage of a specific course via ID.
//
// REQUEST: course ID
// RESPONSE: course data
func (app *application) courseReadHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	app.logger.Printf("Reading course...")

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

	app.logger.Printf("sending response: %#v", res)

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
	courses, err := app.services.UserService.GetUserCourses(input.UserId)
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

// REQUEST: courseid + image file
// RESPONSE: status
func (app *application) bannerCreateHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	app.logger.Printf("Creating banner...")

	courseid := r.PathValue("mediaId")
	// Limit upload size to 10MB
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	f, handler, err := r.FormFile("file")
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	defer f.Close()

	ft := GetFileType(handler.Filename)

	if ft == models.NULL {
		app.serverError(w, r, fmt.Errorf("invalid file type"))
		return
	}

	fileName := courseid + "_banner." + ft.String()

	// Save the file to disk
	path, err := app.services.FileService.Save(fileName, f)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	app.logger.Printf("Saved banner: %s to disk...", fileName)

	// Create metadata and add to database
	metadata := &models.Media{
		FileName:           handler.Filename,
		AttributionsByType: make(map[string]string),
		FileType:           ft,
		FilePath:           path,
	}

	metadata.AttributionsByType["course"] = courseid

	_, err = app.services.MediaService.AddBanner(metadata)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.logger.Printf("Created banner: %v", metadata)
	err = app.writeJSON(w, http.StatusOK, nil, nil)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

// REQUEST: banner id
// RESPONSE: banner image
func (app *application) bannerReadHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	bannerId := r.PathValue("mediaId")

	app.logger.Printf("received Banner ID: %s", bannerId)

	banner, err := app.services.MediaService.GetMedia(bannerId)
	if err != nil {
		app.serverError(w, r, err)
	}

	app.logger.Printf("retrieved metadata: %v", banner)

	// Set Content-Type header based on file extension
	contentType := mime.TypeByExtension("." + banner.FileType.String())
	if contentType == "" {
		contentType = "application/octet-stream" // Default content type
	}

	app.logger.Printf("Setting content type to: %s", contentType)

	contentDispositionValue := "inline"

	filePath := banner.FilePath
	if filePath == "" {
		banner.FilePath = app.services.FileService.Path() + "/defaults/" + banner.FileName + "." + banner.FileType.String()
	}

	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Disposition", contentDispositionValue)

	app.logger.Printf("File path: %s", banner.FilePath)

	// Serve the file's content
	http.ServeFile(w, r, banner.FilePath)
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

	netId, err := app.services.AuthenticationService.GetNetIdFromToken(input.Token)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	//msg := &models.Message{
	//	Post: models.Post{
	//		Title:       input.Title,
	//		Description: input.Description,
	//		Owner:       netid,
	//		Media:       input.Media,
	//	},
	//	Type: true,
	//}
	//
	//msg, err = app.services.MessageService.CreateMessage(msg, cId)
	msg, err := app.services.MessageService.CreateAnnouncement(
		input.Title,
		input.Description,
		netId,
		cId,
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

	app.logger.Printf("retrieved announcements: %v", msgs)

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
		CourseId     string `json:"courseid"`
		TeacherId    string `json:"teacherid"`
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
// RESPONSE: status
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
	user, err = app.services.UserService.GetByID(id)
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
	var noToken bool
	var input struct {
		NetId    string `json:"netid"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// 	Validate credentials.
	err = app.services.UserService.ValidateUser(input.NetId, input.Password)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.logger.Printf("user validated")

	// If token exists, return token:
	token, err := app.services.AuthenticationService.RetrieveToken(input.NetId)
	if err != nil {
		switch {
		case errors.Is(err, dal.ERR_RECORD_NOT_FOUND):
			app.logger.Printf("user token not found for %s", input.NetId)
			noToken = true
		default:
			app.serverError(w, r, err)
			return
		}
	}

	if noToken {
		app.logger.Printf("generating user token...")
		token, err = app.services.AuthenticationService.NewToken(input.NetId)
	}

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.logger.Printf("Token: %v", token)

	membership, err := app.services.UserService.GetMembership(input.NetId)
	if err != nil {
		app.serverError(w, r, err)
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
	app.logger.Printf("Creating assignment...")
	var input struct {
		Title       string   `json:"title"`
		Token       string   `json:"token"`
		Description string   `json:"description"`
		Media       []string `json:"media"`
		DueDate     string   `json:"duedate"`
		CourseId    string   `json:"courseid"`
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

	post := models.Post{
		Title:       input.Title,
		Description: input.Description,
		Owner:       netid,
		Media:       input.Media,
		Course:      input.CourseId,
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
// RESPONSE: assignments
func (app *application) assignmentReadHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	var input struct {
		AssignmentId string `json:"assignmentid"`
		Token        string `json:"token"`
	}

	courseId := r.PathValue("id")

	switch r.Method {
	// Retrieve multiple assignments if its a single GET request.
	case http.MethodGet:
		assignmentIds, err := app.services.AssignmentService.RetrieveAssignments(courseId)
		if err != nil {
			app.serverError(w, r, err)
			return
		}

		var assignments []models.Assignment

		for _, id := range assignmentIds {
			assignment, err := app.services.AssignmentService.ReadAssignment(id)
			if err != nil {
				app.serverError(w, r, err)
				return
			}
			assignments = append(assignments, *assignment)
		}

		res := jsonWrap{"assignments": assignments}

		err = app.writeJSON(w, http.StatusOK, res, nil)
		if err != nil {
			app.serverError(w, r, err)
			return
		}

	// If it's a post request, we can expect an input and body. Therefore,
	// only retrieve a single Assignment.
	case http.MethodPost:
		err := app.readJSON(w, r, &input)
		if err != nil {
			app.serverError(w, r, err)
			return
		}

		assignment, err := app.services.AssignmentService.ReadAssignment(
			input.
				AssignmentId,
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
		return

	// Bad dog.
	default:
		app.serverError(w, r, fmt.Errorf("method %s not allowed", r.Method))
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

	assignment, err := app.services.AssignmentService.UpdateAssignment(input.Uuid, input.UpdatedField, input.Action)
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

func (app *application) assignmentMediaUploadHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	assignmentid := r.PathValue("id")
	// Parse the multipart form
	err := r.ParseMultipartForm(10 << 20) // 10 MB maximum form size
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Retrieve the file(s) from the form
	files := r.MultipartForm.File["files"]

	for _, fileHeader := range files {
		// Open the uploaded file
		file, err := fileHeader.Open()
		if err != nil {
			app.serverError(w, r, err)
			return
		}
		defer file.Close()
		path, err := app.services.FileService.Save(fileHeader.Filename, file)
		if err != nil {
			app.serverError(w, r, err)
			return
		}
		media := &models.Media{
			FileName:           fileHeader.Filename,
			AttributionsByType: make(map[string]string),
			FileType:           GetFileType(fileHeader.Filename),
			FilePath:           path,
		}
		media.AttributionsByType["assignment"] = assignmentid
		media, err = app.services.MediaService.AddAssignmentMedia(media)
		if err != nil {
			app.serverError(w, r, err)
		}
	}
	err = app.writeJSON(w, http.StatusOK, nil, nil)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

// REQUEST: media id
// RESPONSE: assignment media
func (app *application) mediaDownloadHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	mediaid := r.PathValue("mediaId")
	app.logger.Printf("Downloading media...")
	media, err := app.services.MediaService.GetMedia(mediaid)
	if err != nil {
		app.serverError(w, r, err)
	}
	file, err := app.services.FileService.GetFile(media.FilePath)
	if err != nil {
		app.serverError(w, r, err)
	}
	defer file.Close()

	// Get file information (size and name)
	fileInfo, err := file.Stat()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Set Content-Type header based on file extension
	contentType := mime.TypeByExtension(filepath.Ext(fileInfo.Name()))
	if contentType == "" {
		contentType = "application/octet-stream" // Default content type
	}
	contentDisposition := fmt.Sprintf(`attachment; filename="%s"`, fileInfo.Name())
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Disposition", contentDisposition)
	app.logger.Printf("Content-Type: %s", contentType)

	app.logger.Printf("Content-Disposition: %s", contentDisposition)
	// Serve the file's content
	// http.ServeContent(w, r, fileInfo.Name(), fileInfo.ModTime(), file)
	app.logger.Printf("file path: %s", media.FilePath)
	http.ServeFile(w, r, media.FilePath)
}

// Submission handlers
//
// REQUEST: assignmentid + userid
// RESPONSE: submission
func (app *application) submissionCreateHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	assignmentid := r.PathValue("assignmentId")

	var input struct {
		Token string `json:"token"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	userId, err := app.services.AuthenticationService.GetNetIdFromToken(input.Token)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	submission := &models.Submission{
		AssignmentId: assignmentid,
		User: models.User{
			Entity: models.Entity{
				ID: userId,
			},
		},
	}

	assignment, err := app.services.AssignmentService.ReadAssignment(assignmentid)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	submission.OnTime = submission.IsOnTime(assignment.DueDate)
	app.logger.Printf("Creating submission...")
	// Add submission into database and return submission with ID
	submission, err = app.services.SubmissionService.CreateSubmission(submission)
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

// teachersubmissionReadHandler reads a submission from teacher view
// REQUEST: netid + assignmentid
// RESPONSE: submission
func (app *application) teachersubmissionReadHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	assignmentId := r.PathValue("assignmentId")
	userId := r.PathValue("userId")
	app.logger.Printf("Reading student (%s) submission as teacher", userId)

	submission, err := app.services.SubmissionService.GetUserSubmission(userId, assignmentId)
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

// StudentsubmissionReadHandler reads a submission from student view
// REQUEST: assignmentid + token
// RESPONSE: submission
func (app *application) studentsubmissionReadHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	assignmentId := r.PathValue("assignmentId")

	var input struct {
		Token string `json:"token"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	userId, err := app.services.AuthenticationService.GetNetIdFromToken(input.Token)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	submission, err := app.services.SubmissionService.GetUserSubmission(userId, assignmentId)
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
// REQUEST: submission id + feedback + grade
func (app *application) submissionUpdateHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	submissionid := r.PathValue("id")
	app.logger.Printf("Updating submission...")
	var input struct {
		Grade    int    `json:"grade"`
		Feedback string `json:"feedback"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	submission, err := app.services.SubmissionService.GradeSubmission(input.Grade, input.Feedback, submissionid)
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

// SubmissionDeleteHandler deletes a submission
// REQUEST: netid + assignmentid
// RESPONSE: 200 or 404
func (app *application) submissionDeleteHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	panic("implement me")
}

func (app *application) submissionMediaUploadHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	app.logger.Printf("Uploading submission media...")
	submissionid := r.PathValue("id")
	// Parse the multipart form
	err := r.ParseMultipartForm(10 << 20) // 10 MB maximum form size
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Retrieve the file(s) from the form
	files := r.MultipartForm.File["files"]
	app.logger.Printf("Saving submission media...")
	for _, fileHeader := range files {
		// Open the uploaded file

		file, err := fileHeader.Open()
		if err != nil {
			app.serverError(w, r, err)
			return
		}
		defer file.Close()
		fileName := fileHeader.Filename
		path, err := app.services.FileService.Save(fileName, file)
		if err != nil {
			app.serverError(w, r, err)
			return
		}
		media := &models.Media{
			FileName:           fileHeader.Filename,
			AttributionsByType: make(map[string]string),
			FileType:           GetFileType(fileHeader.Filename),
			FilePath:           path,
		}
		media.AttributionsByType["submission"] = submissionid
		media, err = app.services.MediaService.AddSubmissionMedia(media)
		if err != nil {
			app.serverError(w, r, err)
		}
	}
	err = app.writeJSON(w, http.StatusOK, nil, nil)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
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

// sendOfflineTemplate receives a request from a user about
// an offline grading template. It then prepares the template and sends
// it back to the client who requested it.
func (app *application) sendOfflineTemplate(
	w http.ResponseWriter,
	r *http.Request,
) {
	var path string

	courseId := r.PathValue("id")
	assignmentId := r.PathValue("post")

	app.logger.Printf("Retrieving Submissions with assignment id: %s and course id: %s", assignmentId, courseId)

	// Get submissions of this assignment from database.
	submissions, err := app.services.SubmissionService.GetSubmissions(assignmentId)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	app.logger.Printf("Retrieved submissions: %v", submissions)

	// Generate file name for Excel.
	fileName := fmt.Sprintf("submissions_%s_%s", courseId, assignmentId)

	// Prepare the Excel file for transit. WriteSubmissions
	// saves the file to where it needs to be saved.
	path, err = app.services.ExcelService.WriteSubmissions(
		path,
		fileName,
		submissions,
	)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	app.logger.Printf("Saved excel file to %s", path)

	// With this path, send over the file.
	w.Header().Set(
		"Content-Type",
		"application/vnd.openxmlformats-officedocument.spreadsheet",
	)

	headerValue := fmt.Sprintf(`attachment; filename="%s.xlsx"`, fileName)

	w.Header().Set(
		"Content-Disposition",
		headerValue,
	)

	err = app.services.ExcelService.SendFile(path, w)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

// addOfflineGrading will receive an incoming template and will
// sort the itemized template submissions and input them into
// the database.
func (app *application) receiveOfflineGrades(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Limits the upload size to 10MB.
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	f, handler, err := r.FormFile("file")
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	defer f.Close()

	// Check file type.

	// Save the file to disk.
	path, err := app.services.FileService.Save(handler.Filename, f)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Get the submissions from the Excel file via path.
	submissions, err := app.services.ExcelService.ReadSubmissions(path)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Update the submission records in the database.
	err = app.services.SubmissionService.UpdateSubmissions(submissions)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// All is well.
	res := jsonWrap{"response": "successfully updated submissions"}
	err = app.writeJSON(w, http.StatusOK, res, nil)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}
