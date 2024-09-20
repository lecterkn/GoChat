package repository

import (
	"lecter/goserver/db"
	"lecter/goserver/model"

	"github.com/google/uuid"
)

type UserRepository struct{}

func (ur UserRepository) Insert(model model.UserModel) (*model.UserModel, error) {
	err := db.Database().Create(&model).Error
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (ur UserRepository) Select(id uuid.UUID) (*model.UserModel, error) {
	var model model.UserModel
	err := db.Database().Where("id = ?", id[:]).First(&model).Error
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (ur UserRepository) SelectByName(name string) (*model.UserModel, error) {
	var model model.UserModel
	err := db.Database().Where("name = ?", name).First(&model).Error
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (ur UserRepository) Update(model model.UserModel) (*model.UserModel, error) {
	err := db.Database().Save(&model).Error
	if err != nil {
		return nil, err
	}
	return &model, nil
}