package router

import (
	"github/Sourav-Sonkar/mongoApi/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/movies", controller.GetAllMoviesController).Methods("GET")
	r.HandleFunc("/api/movie",controller.CreateMovie).Methods("POST")
	r.HandleFunc("/api/movie/{id}",controller.MarkedAsWatched).Methods("PUT")
	r.HandleFunc("/api/movie/{id}",controller.DeleteOneMovieController).Methods("DELETE")
	r.HandleFunc("/api/deleteAllMovie",controller.DeleteAllMovieController).Methods("DELETE")

	return r
}
