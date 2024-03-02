package main

import (
	"encoding/json"
	"net/http"
)

// writeJSON ingests a map of map[string]any and writes it to a
// http.ResponseWriter stream. It returns an error in case
// there was one writing to the stream.
func (app *application) writeJSON(w http.ResponseWriter, status int, data any, headers http.Header) error {

	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Convenience \n for terminal view
	js = append(js, '\n')

	// Add headers to writer stream.
	for k, v := range headers {
		w.Header()[k] = v
	}

	// Add headers, write them, then write to the output stream.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
