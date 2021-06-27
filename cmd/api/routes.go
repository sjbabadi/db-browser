package main

import "github.com/gorilla/mux"

func (app *application) routes() *mux.Router {
	router := mux.NewRouter()
	api := router.PathPrefix("/api/").Subrouter()

	api.HandleFunc("/genres", app.getGenres)
	api.HandleFunc("/tracks", app.getTracks)

	return router
}
