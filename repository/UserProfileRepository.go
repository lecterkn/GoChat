package repository

import (
	"lecter/goserver/db"
	"lecter/goserver/model"

	"github.com/google/uuid"
)

type UserProfileRepository struct{}

func (upr UserProfileRepository) SelectByUserId(userId uuid.UUID) (*model.UserProfileModel, error) {
	var table model.UserProfileTable
	err := db.Database().QueryRow("SELECT * FROM user_profiles WHERE user_id=$1", userId[:]).Scan(&table.Id, &table.UserId, &table.DisplayName, &table.Url, &table.Description)
	if err != nil {
		return nil, err
	}
	return table.ToModel(), nil
}

func (upr UserProfileRepository) Create(model model.UserProfileModel) (*model.UserProfileModel, error) {
	table := model.ToTable()
	err := db.Database().QueryRow("INSERT INTO user_profiles VALUES($1, $2, $3, $4, $5) RETURNING *", table.Id, table.UserId, table.DisplayName, table.Url, table.Description).Scan(table.ToArray())
	if err != nil {
		return nil, err
	}
	return table.ToModel(), nil
}

func (upt UserProfileRepository) Update(model model.UserProfileModel) (*model.UserProfileModel, error) {
	table := model.ToTable()
	err := db.Database().QueryRow("UPDATE user_profiles SET user_id=$1,display_name=$2,url=$3,description=$4 WHERE id=$5 RETURNING *", table.UserId, table.DisplayName, table.Url, table.Description, table.Id).Scan(table.ToArray())
	if err != nil {
		return nil, err
	}
	return table.ToModel(), nil
}