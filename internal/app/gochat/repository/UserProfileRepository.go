package repository

import (
	"lecter/goserver/internal/app/gochat/db"
	"lecter/goserver/internal/app/gochat/model"

	"github.com/google/uuid"
)

type UserProfileRepository struct{}

func (upr UserProfileRepository) SelectByUserId(userId uuid.UUID) (*model.UserProfileModel, error) {
	var model model.UserProfileModel
	err := db.Database().Where("user_id = ?", userId[:]).First(&model).Error
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (upr UserProfileRepository) Create(model model.UserProfileModel) (*model.UserProfileModel, error) {
	err := db.Database().Create(&model).Error
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (upt UserProfileRepository) Update(model model.UserProfileModel) (*model.UserProfileModel, error) {
	err := db.Database().Save(&model).Error
	if err != nil {
		return nil, err
	}
	return &model, nil
}
