package main

import (
	"encoding/json"
	"net/http"
)

// jsonWrap wraps a json message response before it gets sent out.
// This makes it easier for the requester to read the JSON data
// they get back. You can imagine that this "envelops" the json
// data that will go into it.
type jsonWrap map[string]any

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

	// Add any headers
	for k, v := range headers {
		w.Header()[k] = v
	}

	// Add headers, then write to the output stream.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
