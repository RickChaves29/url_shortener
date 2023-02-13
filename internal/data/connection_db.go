package data

import (
	"database/sql"
	"os"
)

func ConnectionDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("CONNECT_DB"))
	if err != nil {
		println("error when try open database")
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		println("error when try connection on database")
	}

	println("connection on database is good")

	return db, err
}
