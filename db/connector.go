package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type DBConnector struct {}

var db *sql.DB = nil

func Connect() error {
	dataSource := "host=localhost port=5432 user=postgres password=mysecretpassword dbname=mydb sslmode=disable"
	var err error
	db, err = sql.Open("postgres", dataSource)
	if (err != nil) {
		return err
	}
	return nil
}

func DB() *sql.DB {
	return db
}