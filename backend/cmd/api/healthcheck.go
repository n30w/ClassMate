package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	res := map[string]any{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	err := app.writeJSON(w, http.StatusOK, res, nil)

	if err != nil {
		app.logger.Print(err)
		http.Error(w, "The server could not process your request", http.StatusInternalServerError)
	}
}
