package main

import (
	"errors"
	"fmt"
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
	message := "invalid authentication credentials"
	app.errorResponse(w, r, http.StatusUnauthorized, message)
}

func (app *application) notFoundResponse(
	w http.ResponseWriter,
	r *http.Request,
) {
	message := "the requested resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

func (app *application) methodNotAllowedResponse(
	w http.ResponseWriter,
	r *http.Request,
) {
	message := fmt.Sprintf(
		"the %s method is not supported for this resource",
		r.Method,
	)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}

func (app *application) badRequestResponse(
	w http.ResponseWriter,
	r *http.Request,
	err error,
) {
	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func (app *application) failedValidationResponse(
	w http.ResponseWriter,
	r *http.Request,
	errors map[string]string,
) {
	app.errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}

func (app *application) editConflictResponse(
	w http.ResponseWriter,
	r *http.Request,
) {
	message := "unable to update the record due to an edit conflict, please try again"
	app.errorResponse(w, r, http.StatusConflict, message)
}

func (app *application) invalidAuthenticationTokenResponse(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("WWW-Authenticate", "Bearer")

	message := "invalid or missing authentication token"
	app.errorResponse(w, r, http.StatusUnauthorized, message)
}

func (app *application) authenticationRequiredResponse(
	w http.ResponseWriter,
	r *http.Request,
) {
	message := "you must be authenticated to access this resource"
	app.errorResponse(w, r, http.StatusUnauthorized, message)
}

func (app *application) inactiveAccountResponse(
	w http.ResponseWriter,
	r *http.Request,
) {
	message := "your user account must be activated to access this resource"
	app.errorResponse(w, r, http.StatusForbidden, message)
}

func (app *application) notPermittedResponse(
	w http.ResponseWriter,
	r *http.Request,
) {
	message := "your user account doesn't have the necessary permissions to access this resource"
	app.errorResponse(w, r, http.StatusForbidden, message)
}
