package models

import (
	"database/sql"
	"fmt"
	"restApi/internal/database"
)

var dbFilm *sql.DB

type Film struct {
	Id int64 `json:"id_film"`
	GenreId int64  `json:"id_genre"`
	Name string `json:"name"`
}

func init()  {
	d := database.New()
	dbFilm = d.GetDB()
}

func (film Film) GetAll(idGenre int64) []Film{
	var(
		f Film
		films[] Film
	)
	sqlStatement :=  "select films.id_film, list.genre_id, films.name from list " +
		"left join films on list.film_id = films.id_film  where genre_id=$1"
	rows, err := dbFilm.Query(sqlStatement, idGenre)
	if err != nil{
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		rows.Scan(&f.Id, &f.GenreId, &f.Name)
		films = append(films, f)
	}
	defer rows.Close()
	return  films
}

func (film Film) GetById(idFilm, idGenre int64) (*Film, error) {
	f := &Film{}
	sqlStatement :=  "select films.id_film, list.genre_id, films.name from list " +
		"left join films on list.film_id = films.id_film  where film_id=$1 and genre_id=$2"
	err := dbFilm.QueryRow(sqlStatement, idFilm, idGenre,
	).Scan(
		&f.Id,
		&f.GenreId,
		&f.Name,
	)
	if err != nil{
		return nil , err
	}
	return f, nil
}

func (film Film) Create(f *Film) (int64, error) {
	sqlInFilms :=  "insert into films (name) values ($1) returning id_film"
	id := 0
	err := dbFilm.QueryRow(sqlInFilms, f.Name).Scan(&id)
	if err != nil {
		return 0, err
	}
	sqlInList :=  "insert into list (film_id, genre_id) values ($1,$2)"
	_, err = dbFilm.Exec(sqlInList, id, f.GenreId)
	if err != nil {
		return 0, err
	}
	return int64(id), nil
}

func (film Film) Update(f *Film, id int64) error {
	sqlInFilms :=  "update films set name=$2  where id_film=$1"
	_, err :=  dbGenre.Exec(sqlInFilms, id, f.Name)
	if err != nil {
		return err
	}
	sqlInList :=  "update list set genre_id=$2  where id_film=$1"
	_, err =  dbGenre.Exec(sqlInList, id, f.GenreId)
	if err != nil {
		return err
	}
	return nil
}

func (film Film) Delete(idFilm, idGenre int64) error {
	sqlInList :=  "delete from list where film_id=$1 and genre_id=$2"
	_, err :=  dbGenre.Exec(sqlInList, idFilm, idGenre)
	if err != nil {
		return err
	}
	sqlInFilms :=  "delete from films where id_film=$1"
	_, err =  dbGenre.Exec(sqlInFilms, idFilm, idGenre)
	if err != nil {
		return err
	}
	return nil
}
