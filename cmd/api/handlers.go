package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) getGenres(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	js := `{"name": "Sheila", "dogs": 3}`

	msgJSON, err := json.MarshalIndent(js, "", "\t")
	app.check(err)
	w.Write(msgJSON)
}
