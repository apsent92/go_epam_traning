package apiserver

import (

	"github.com/gorilla/mux"
	"log"
	"net/http"
	"restApi/internal/controllers"
)

type APIServer struct{
	router *mux.Router
}

func (s *APIServer) Start() {
	s.router = mux.NewRouter()
	s.initRoutes()
	log.Fatal(http.ListenAndServe(":8080", s.router))
}

func (s *APIServer) initRoutes()  {
	s.router.HandleFunc("/genres", controllers.GetGenres).Methods("GET")
	s.router.HandleFunc("/genre/", controllers.CreateGenre).Methods("POST")
	s.router.HandleFunc("/genre/{genre_id}", controllers.GetGenreById).Methods("GET")
	s.router.HandleFunc("/genre/{genre_id}", controllers.UpdateGenre).Methods("PUT")
	s.router.HandleFunc("/genre/{genre_id}", controllers.DeleteGenre).Methods("DELETE")
	s.router.HandleFunc("/genre/{genre_id}/films", controllers.GetFilms).Methods("GET")
	s.router.HandleFunc("/genre/{genre_id}/film/", controllers.CreateFilm).Methods("POST")
	s.router.HandleFunc("/genre/{genre_id}/film/{film_id}", controllers.GetFilmById).Methods("GET")
	s.router.HandleFunc("/genre/{genre_id}/film/{film_id}", controllers.UpdateFilm).Methods("PUT")
	s.router.HandleFunc("/genre/{genre_id}/film/{film_id}", controllers.DeleteFilm).Methods("DELETE")
}
