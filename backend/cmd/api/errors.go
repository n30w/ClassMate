package main

import "net/http"

// Methods in this file define error handling functions.
// logError wraps the app.logger.Print() method in order to
// add more sophistication later on to the error logging
// capabilities. Many of these functions contain
// side effects.

// logError logs an error using the application's logger.
func (app *application) logError(r *http.Request, err error) {
	app.logger.Print(err)
}

// internalServerError is a 500 status error.
func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)
	http.Error(w, "The server could not process your request", http.StatusInternalServerError)
}
