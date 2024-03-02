package main

import (
	"net/http"
)

// healthCheckHandler handles healthcheck requests. It uses writeJSON()
// to write system health data to the http.ResponseWriter stream.
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	app.logger.Printf("received %s from %s", r.Method, r.RemoteAddr)

	res := jsonWrap{
		"status": "available",
		"system_info": map[string]any{
			"status":      "available",
			"environment": app.config.env,
			"version":     version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, res, nil)

	if err != nil {
		app.internalServerError(w, r, err)
	}
}
