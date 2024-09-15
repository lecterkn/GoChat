package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var database *sql.DB = nil

func Connect() error {
	dataSource := "host=localhost port=5432 user=postgres password=mysecretpassword dbname=mydb sslmode=disable"
	var err error
	database, err = sql.Open("postgres", dataSource)
	if (err != nil) {
		return err
	}
	return nil
}

func Database() *sql.DB {
	return database
}