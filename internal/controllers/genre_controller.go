package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"restApi/internal/models"
	"strconv"
)

func GetGenres(w http.ResponseWriter, r *http.Request)  {
	genres := models.Genre{}.GetAll()
	res, _ := json.Marshal(genres)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetGenreById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["genre_id"]
	genre_id, err:= strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	genre, _:= models.Genre{}.GetById(genre_id)
	res, _ := json.Marshal(genre)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func  CreateGenre(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var genre models.Genre
	json.Unmarshal(reqBody, &genre)
	id, _ := models.Genre{}.Create(&genre)
	genre.Id = id
	res,_ := json.Marshal(genre)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateGenre(w http.ResponseWriter, r *http.Request) {
	var updateGenre models.Genre
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &updateGenre)
	vars := mux.Vars(r)
	id := vars["genre_id"]
	upd_id, err:= strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	g, _ := models.Genre{}.GetById(upd_id)
	if g != nil {
		fmt.Println("cxgsfhdfhsdh", upd_id)
		models.Genre{}.Update(&updateGenre, upd_id)
		res, _ := json.Marshal(&updateGenre)
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func DeleteGenre(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["genre_id"]
	genre_id, err:= strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	models.Genre{}.Delete(genre_id)
}