package main

import (
	"net/http"
)

func (app *application) getGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := app.genreModel.GetAll()
	app.check(err)

	err = app.writeJSON(w, genres, http.StatusOK)
	app.check(err)
}

func (app *application) getTracks(w http.ResponseWriter, r *http.Request) {
	tracks, err := app.trackModel.GetAll()
	app.check(err)

	err = app.writeJSON(w, tracks, http.StatusOK)
	app.check(err)
}
