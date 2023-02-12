package data

import "database/sql"

func ConnectionDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://postgres:docker@localhost:5432/postgres?sslmode=disable")
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
