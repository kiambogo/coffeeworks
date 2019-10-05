package main

import (
	"github.com/gorilla/mux"
)

// BuildRoutes composes the various routes used in the application with their respective handlers
func BuildRoutes() *mux.Router {
	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/cafes", ListCafes).Methods("GET")

	return router
}
