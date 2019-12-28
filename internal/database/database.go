package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type Database struct {
	database *sql.DB
}

func New() *Database {
	db, err := sql.Open("postgres", "host=db port=5432 dbname=mydb user=root password =root sslmode=disable")
	if err != nil{
		log.Fatal(err)
	}
	/*if err := db.Ping(); err != nil{
		log.Fatal(err)
	}*/
	return &Database{
		database: db,
	}
}

func (database *Database) GetDB() *sql.DB {
	return database.database
}