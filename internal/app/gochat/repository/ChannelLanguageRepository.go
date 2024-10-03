package repository

import (
	"lecter/goserver/internal/app/gochat/db"
	"lecter/goserver/internal/app/gochat/model"

	"github.com/google/uuid"
)

type ChannelLanguageRepository struct{}

func (ChannelLanguageRepository) Index(channelId uuid.UUID) ([]model.ChannelLanguageModel, error) {
	var models []model.ChannelLanguageModel
	err := db.Database().Where("channel_id = ?", channelId[:]).Find(&models).Error
	if err != nil {
		return nil, err
	}
	return models, nil
}

func (ChannelLanguageRepository) Delete(channelId uuid.UUID) error {
	err := db.Database().Where("channel_id = ?", channelId[:]).Delete(&model.ChannelLanguageModel{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (ChannelLanguageRepository) InsertAll(models []model.ChannelLanguageModel) ([]model.ChannelLanguageModel, error) {
	err := db.Database().Create(models).Error
	if err != nil {
		return nil, err
	}
	return models, nil
}
