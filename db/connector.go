package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type DBConnector struct {}

var db *sql.DB = nil

func Connect() {
	dataSource := "host=localhost port=5432 user=postgres password=mysecretpassword dbname=mydb sslmode=disable"
	var err error
	db, err = sql.Open("postgres", dataSource)
	if (err != nil) {
		fmt.Println("DB Connection ERROR")
		fmt.Println(err.Error())
		return
	}
	fmt.Println("DB Connected")
}

func DB() *sql.DB {
	return db
}