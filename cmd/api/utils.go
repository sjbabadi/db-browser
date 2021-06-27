package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// take in some data
// marshal it/ encode into json
// set res headers
// write status
// write response
func (app *application) writeJSON(w http.ResponseWriter, data interface{}, status int) error {
	js, err := json.MarshalIndent(data, "", "\t")
	app.check(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

func (app *application) check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
