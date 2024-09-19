package repository

import (
	"fmt"
	"lecter/goserver/db"
	"lecter/goserver/model"

	"github.com/google/uuid"
)

type UserRepository struct{}

func (ur UserRepository) Insert(model model.UserModel) (*model.UserModel, error) {
	table := model.ToTable()
	connector := db.Database()
	err := connector.QueryRow("INSERT INTO users VALUES($1, $2, $3, $4, $5) RETURNING *", table.Id, table.Name, table.Password, table.CreatedAt, table.UpdatedAt).Scan(&table.Id, &table.Name, &table.Password, &table.CreatedAt, &table.UpdatedAt)
	if err != nil {
		return nil, err
	}
	model = *table.ToModel()
	return &model, nil
}

func (ur UserRepository) Select(id uuid.UUID) (*model.UserModel, error) {
	var table model.UserTable
	connector := db.Database()
	err := connector.QueryRow("SELECT * FROM users WHERE id=$1", id[:]).Scan(&table.Id, &table.Name, &table.Password, &table.CreatedAt, &table.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return table.ToModel(), nil
}

func (ur UserRepository) SelectByName(name string) (*model.UserModel, error) {
	var table model.UserTable
	connector := db.Database()
	err := connector.QueryRow("SELECT * FROM users WHERE name=$1", name).Scan(&table.Id, &table.Name, &table.Password, &table.CreatedAt, &table.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return table.ToModel(), nil
}

func (ur UserRepository) Index() ([]model.UserModel, error) {
	// モデルとテーブル
	tables := []model.UserTable{}	
	users := []model.UserModel{}

	// db接続
	connector := db.Database()
	rows, err := connector.Query("SELECT * FROM users")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var table model.UserTable
		err := rows.Scan(&table.Id, &table.Name, &table.Password, &table.CreatedAt, &table.UpdatedAt)
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
	return users, nil
}

func (ur UserRepository) Update(model model.UserModel) (*model.UserModel, error) {
	table := model.ToTable()
	connector := db.Database()
	err := connector.QueryRow("UPDATE users SET name=$1, password=$2, created_at=$3, updated_at=$4 WHERE id=$5 RETURNING *", table.Name, table.Password, table.CreatedAt, table.UpdatedAt, table.Id).Scan(&table.Id, &table.Name, &table.Password, &table.CreatedAt, &table.UpdatedAt)
	if err != nil {
		return nil, err
	}
	model = *table.ToModel()
	return &model, nil
}