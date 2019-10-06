package main

import (
	"github.com/gorilla/mux"
)

// BuildRoutes composes the various routes used in the application with their respective handlers
func BuildRoutes() *mux.Router {
	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api").Subrouter()

	cafesRouter := apiRouter.PathPrefix("/cafes").Subrouter()
	cafesRouter.HandleFunc("", ListCafes).Methods("GET")
	cafesRouter.HandleFunc("/{id}", GetCafe).Methods("GET")

	reviewsRouter := apiRouter.PathPrefix("/reviews").Subrouter()
	reviewsRouter.HandleFunc("", CreateReview).Methods("POST")

	return router
}
