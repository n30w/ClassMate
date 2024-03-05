package main

import (
	"errors"
	"net/http"
)

// Methods in this file define error handling functions.
// logError wraps the app.logger.Print() method in order to
// add more sophistication later on to the error logging
// capabilities. Many of these functions contain
// side effects.

type ServerError error

var (
	MALFORMED_JSON_SYNTAX ServerError = ServerError(errors.New("malformed json syntax"))
)

// logError logs an error using the application's logger.
func (app *application) logError(r *http.Request, err error) {
	app.logger.Print(err)
}

// serverError returns a 400 Bad Request. This is called when JSON is messed up.
func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)
	http.Error(w, err.Error(), http.StatusBadRequest)
}
