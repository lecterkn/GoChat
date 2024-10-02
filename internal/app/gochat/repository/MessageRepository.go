package repository

import (
	"lecter/goserver/internal/app/gochat/db"
	"lecter/goserver/internal/app/gochat/model"
	"time"

	"github.com/google/uuid"
)

type MessageRepository struct{}

func (MessageRepository) Index(
	channelId uuid.UUID,
	lastId *uuid.UUID,
	lastCreatedAt *time.Time,
	limit int) (*[]model.MessageModel, error) {
	var models []model.MessageModel
	var err error
	if lastId == nil || lastCreatedAt == nil {
		err = db.Database().
			Where("channel_id = ?", channelId[:]).
			Where("deleted = FALSE").
			Order("created_at DESC, id").
			Limit(limit).
			Find(&models).Error
	} else {
		err = db.Database().
			Where("channel_id = ?", channelId[:]).
			Where("(created_at < ? OR created_at = ? AND id > ?)", lastCreatedAt, lastCreatedAt, lastId[:]).
			Where("deleted = FALSE").
			Order("created_at DESC, id").
			Limit(limit).
			Find(&models).Error
	}
	if err != nil {
		return nil, err
	}
	return &models, nil
}

func (MessageRepository) Select(id uuid.UUID) (*model.MessageModel, error) {
	var model model.MessageModel
	err := db.Database().Where("id = ? AND deleted = FALSE", id[:]).First(&model).Error
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (MessageRepository) Create(model model.MessageModel) (*model.MessageModel, error) {
	err := db.Database().Create(&model).Error
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (MessageRepository) Update(model model.MessageModel) (*model.MessageModel, error) {
	err := db.Database().Where("deleted = FALSE").Save(&model).Error
	if err != nil {
		return nil, err
	}
	return &model, nil
}
