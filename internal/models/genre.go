package models

import (
	"database/sql"
	"fmt"
	"restApi/internal/database"
)

var dbGenre *sql.DB

type Genre struct{
	Id	 int64 `json:"id_genre"`
	Title string `json:"title"`
}

func init()  {
	d := database.New()
	dbGenre = d.GetDB()
}

func (genre Genre) GetAll() []Genre {
	var (
		g Genre
		genres[] Genre
	)
	sqlStatement := "select * from genres"
	rows, err := dbGenre.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		rows.Scan(&g.Id, &g.Title)
		genres = append(genres, g)
	}
	defer rows.Close()
	return genres
}

func (g Genre) GetById(id int64)  (*Genre, error) {
	genre := &Genre{}
	sqlStatement := "select * from genres where id_genre=$1"
	err := dbGenre.QueryRow(sqlStatement, id,
	).Scan(
		&genre.Id,
		&genre.Title,
	)
	if err != nil {
		return nil, err
	}
	return genre , nil
}

func (g Genre) Update(genre *Genre, id int64)  error {
	sqlStatement := "update genres set title=$2  where id_genre=$1"
	_, err :=  dbGenre.Exec(sqlStatement, id, genre.Title)
	if err != nil {
		return err
	}
	return nil
}

func (g Genre) Create(genre *Genre) (int64, error) {
	sqlStatement := "insert into genres (title) values ($1) returning id_genre"
	id:= 0
	err := dbGenre.QueryRow(sqlStatement, genre.Title).Scan(&id)
	if err != nil {
		return 0, err
	}
	return int64(id), nil
}

func (g Genre) Delete(id int64) error {
	sqlFromList := "delete from list where genre_id=$1"
	_, err :=  dbGenre.Exec(sqlFromList, id)
	if err != nil {
		return err
	}
	sqlStatement := "delete from genres where id_genre=$1"
	_, err =  dbGenre.Exec(sqlStatement, id)
	if err != nil {
		return err
	}
	return nil
}