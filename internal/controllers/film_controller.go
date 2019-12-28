package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"restApi/internal/models"
	"strconv"
)

func GetFilms(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	idGenre := vars["genre_id"]
	genreId, _:= strconv.ParseInt(idGenre, 0, 0)
	films := models.Film{}.GetAll(genreId)
	res, _ := json.Marshal(films)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetFilmById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idFilm := vars["film_id"]
	idGenre := vars["genre_id"]
	filmId, _:= strconv.ParseInt(idFilm, 0, 0)
	genreId, _:= strconv.ParseInt(idGenre, 0, 0)
	film, _:= models.Film{}.GetById(filmId, genreId)
	res, _ := json.Marshal(film)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateFilm(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var film models.Film
	json.Unmarshal(reqBody, &film)
	vars := mux.Vars(r)
	idGenre := vars["genre_id"]
	genreId, _:= strconv.ParseInt(idGenre, 0, 0)
	genre, _ := models.Genre{}.GetById(genreId)
	if genre != nil {
		if genreId != film.GenreId {
			film.GenreId = genreId
		}
		id, _ := models.Film{}.Create(&film)
		film.Id = id
		res,_ := json.Marshal(film)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}

}

func UpdateFilm(w http.ResponseWriter, r *http.Request) {
	var updateFilm models.Film
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &updateFilm)
	vars := mux.Vars(r)
	idFilm := vars["film_id"]
	idGenre := vars["genre_id"]
	filmId, _:= strconv.ParseInt(idFilm, 0, 0)
	genreId, _:= strconv.ParseInt(idGenre, 0, 0)
	g, _ := models.Genre{}.GetById(genreId)
	if g != nil {
		if genreId != updateFilm.GenreId {
			updateFilm.GenreId = genreId
		}
		models.Film{}.Update(&updateFilm, filmId)
	}
}

func DeleteFilm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idFilm := vars["film_id"]
	idGenre := vars["genre_id"]
	filmId, _:= strconv.ParseInt(idFilm, 0, 0)
	genreId, _:= strconv.ParseInt(idGenre, 0, 0)
	film, _:= models.Film{}.GetById(filmId, genreId)
	if film != nil {
		models.Film{}.Delete(filmId, genreId)
	}
}
