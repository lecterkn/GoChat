package db

import (
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func Connect() error {
	dataSource := "host=localhost port=5432 user=postgres password=mysecretpassword dbname=mydb sslmode=disable"
	var err error
	database, err = gorm.Open(postgres.Open(dataSource), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func Close() error {
	if database == nil {
		return fmt.Errorf("database is not defined")
	}
	db, err := database.DB()
	if err != nil {
		return err
	}
	err = db.Close()
	if err != nil {
		return err
	}
	return nil
}

func Database() *gorm.DB {
	if database == nil {
		Connect()
	}
	return database
}
