package repository

import (
	"fmt"
	"lecter/hello/db"
	"lecter/hello/model"

)

type UserRepository struct{}

func (ur UserRepository) Insert(model *model.UserModel) (*model.UserModel) {
	table := model.ToTable()
	connector := db.DB()
	err := connector.QueryRow("INSERT INTO users (id, name, url) VALUES($1, $2, $3) RETURNING *", table.Id, table.Name, table.Url).Scan(&table.Id, &table.Name, &table.Url)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	*model = *table.ToModel()
	return model
}

func (ur UserRepository) Index() ([]model.UserModel) {
	// モデルとテーブル
	tables := []model.UserTable{}	
	users := []model.UserModel{}

	// db接続
	connector := db.DB()
	rows, err := connector.Query("SELECT * FROM users")

	if err != nil {
		fmt.Println(err)
		return users
	}
	defer rows.Close()

	for rows.Next() {
		var table model.UserTable
		err := rows.Scan(&table.Id, &table.Name, &table.Url)
		if err != nil {
			fmt.Println("failed to convert UserTable")
			continue
		}
		tables = append(tables, table)
	}

	for _, table := range tables {
		model := table.ToModel()
		if model == nil {
			continue
		}
		users = append(users, *model)
	}
	return users
}