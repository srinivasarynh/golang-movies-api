package router

import (
	controller "mongoapi/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PuT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteAMovie).Methods("DELETE")

	return router
}
