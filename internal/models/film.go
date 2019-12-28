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
	sqlStatement :=  "select * from films where id_genre=$1"
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
	sqlStatement :=  "select * from films where id_film=$1 and id_genre=$2"
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
	sqlStatement :=  "insert into films (id_genre, name) values ($1,$2) returning id_film"
	id := 0
	err := dbFilm.QueryRow(sqlStatement, f.GenreId, f.Name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return int64(id), nil
}

func (film Film) Update(f *Film, id int64) error {
	sqlStatement :=  "update films set id_genre=$2, name=$3  where id_film=$1"
	_, err :=  dbGenre.Exec(sqlStatement, id, f.GenreId, f.Name)
	if err != nil {
		return err
	}
	return nil
}

func (film Film) Delete(idFilm, idGenre int64) error {
	sqlStatement :=  "delete from films where id_film=$1 and id_genre=$2"
	_, err :=  dbGenre.Exec(sqlStatement, idFilm, idGenre)
	if err != nil {
		return err
	}
	return nil
}
