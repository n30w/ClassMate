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

func (app *application) errorResponse(
	w http.ResponseWriter, r *http.Request,
	status int, message any,
) {
	wrap := jsonWrap{"error": message}

	err := app.writeJSON(w, status, wrap, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

// serverError returns a 400 Bad Request. This is called when JSON is messed up.
func (app *application) serverError(
	w http.ResponseWriter,
	r *http.Request,
	err error,
) {
	app.logError(r, err)
	http.Error(w, err.Error(), http.StatusBadRequest)
}

// rateLimitExceededResponse returns a 429 Too Many Requests response.
func (app *application) rateLimitExceededResponse(
	w http.ResponseWriter,
	r *http.Request,
) {
	message := "rate limit exceeded"
	app.errorResponse(w, r, http.StatusTooManyRequests, message)
}

func (app *application) invalidCredentialsResponse(
	w http.ResponseWriter,
	r *http.Request,
) {
	message := "invalid authentication creentials"
	app.errorResponse(w, r, http.StatusUnauthorized, message)
}
