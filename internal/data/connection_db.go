package data

import (
	"database/sql"
	"log"
	"os"
)

func ConnectionDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("CONNECT_DB"))
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Println("error when try connection on database")
	}

	log.Println("connection on database is good")

	return db, err
}
