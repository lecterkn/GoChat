package repository

import (
	"lecter/goserver/internal/app/gochat/db"
	"lecter/goserver/internal/app/gochat/model"

	"github.com/google/uuid"
)

type ChannelRepository struct{}

func (ChannelRepository) Index() ([]model.ChannelModel, error) {
	var models []model.ChannelModel
	err := db.Database().Where("deleted = FALSE").Find(&models).Error
	if err != nil {
		return nil, err
	}
	return models, nil
}

func (ChannelRepository) Select(id uuid.UUID) (*model.ChannelModel, error) {
	var model model.ChannelModel
	err := db.Database().Where("id = ? AND deleted = FALSE", id[:]).First(&model).Error
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (ChannelRepository) Create(model model.ChannelModel) (*model.ChannelModel, error) {
	err := db.Database().Create(&model).Error
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (ChannelRepository) Update(model model.ChannelModel) (*model.ChannelModel, error) {
	err := db.Database().Where("deleted = FALSE").Save(&model).Error
	if err != nil {
		return nil, err
	}
	return &model, nil
}
