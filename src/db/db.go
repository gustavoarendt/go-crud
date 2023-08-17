package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func DbConnection() *sql.DB {
	connectionString := "user=admin dbname=go-crud password=admin host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err.Error())
	}
	return db
}
