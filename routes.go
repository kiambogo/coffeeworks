package main

import (
	"github.com/gorilla/mux"
)

func BuildRoutes() *mux.Router {
	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/cafes", GetCafes).Methods("GET")

	return router
}
