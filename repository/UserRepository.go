package repository

import (
	"fmt"
	"lecter/goserver/db"
	"lecter/goserver/model"

	"github.com/google/uuid"
)

type UserRepository struct{}

func (ur UserRepository) Insert(model model.UserModel) (*model.UserModel) {
	table := model.ToTable()
	connector := db.Database()
	err := connector.QueryRow("INSERT INTO users (id, name, url) VALUES($1, $2, $3) RETURNING *", table.Id, table.Name, table.Url).Scan(&table.Id, &table.Name, &table.Url)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	model = *table.ToModel()
	return &model
}

func (ur UserRepository) Select(id uuid.UUID) (*model.UserModel) {
	var table model.UserTable
	connector := db.Database()
	err := connector.QueryRow("SELECT * FROM users WHERE id=$1", id[:]).Scan(&table.Id, &table.Name, &table.Url)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return table.ToModel()
}

func (ur UserRepository) Index() ([]model.UserModel) {
	// モデルとテーブル
	tables := []model.UserTable{}	
	users := []model.UserModel{}

	// db接続
	connector := db.Database()
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
			fmt.Println(err.Error())
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

func (ur UserRepository) Update(model model.UserModel) (*model.UserModel) {
	table := model.ToTable()
	connector := db.Database()
	err := connector.QueryRow("UPDATE users SET name=$1, url=$2 WHERE id=$3 RETURNING *", table.Name, table.Url, table.Id).Scan(&table.Id, &table.Name, &table.Url)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	model = *table.ToModel()
	return &model
}