package repository

import (
	"lecter/goserver/internal/app/gochat/db"
	"lecter/goserver/internal/app/gochat/enum/language"
	"lecter/goserver/internal/app/gochat/model"

	"github.com/google/uuid"
)

type MessageRepository struct{}

func (MessageRepository) Index(
	channelId uuid.UUID,
	lastMessage *model.MessageModel,
	limit int) ([]model.MessageModel, error) {
	var models []model.MessageModel
	var err error
	if lastMessage == nil {
		err = db.Database().
			Where("channel_id = ?", channelId[:]).
			Where("deleted = FALSE").
			Order("created_at DESC, id").
			Limit(limit).
			Find(&models).Error
	} else {
		err = db.Database().
			Where("channel_id = ?", channelId[:]).
			Where("(created_at < ? OR created_at = ? AND id > ?)", lastMessage.CreatedAt, lastMessage.CreatedAt, lastMessage.Id[:]).
			Where("deleted = FALSE").
			Order("created_at DESC, id").
			Limit(limit).
			Find(&models).Error
	}
	if err != nil {
		return nil, err
	}
	return models, nil
}

func (MessageRepository) TranslatedMessageIndex(
	channelId uuid.UUID,
	lastMessage *model.MessageModel,
	limit int,
	lang language.Language) ([]model.TranslatedMessageModel, error) {
	var models []model.TranslatedMessageModel
	var err error
	if lastMessage == nil {
		err = db.Database().
			Select("messages.*, COALESCE("+lang.TableName()+".content, messages.message) AS message_content").
			Joins("LEFT JOIN "+lang.TableName()+" ON "+lang.TableName()+".channel_id = messages.channel_id AND "+lang.TableName()+".message_id = messages.id").
			Where("messages.channel_id = ?", channelId[:]).
			Where("messages.deleted = FALSE").
			Order("created_at DESC, id").
			Limit(limit).
			Find(&models).Error
	} else {
		err = db.Database().
			Select("messages.*, COALESCE(message_?_contents.content, messages.message) AS message_content", lang.ToName()).
			Joins("LEFT JOIN message_?_contents ON message_?_contents.channel_id = messages.channel_id AND message_?_contents.message_id = messages.id", lang.ToName(), lang.ToName(), lang.ToName()).
			Where("channel_id = ?", channelId[:]).
			Where("(created_at < ? OR created_at = ? AND id > ?)", lastMessage.CreatedAt, lastMessage.CreatedAt, lastMessage.Id[:]).
			Where("deleted = FALSE").
			Order("created_at DESC, id").
			Limit(limit).
			Find(&models).Error
	}
	if err != nil {
		return nil, err
	}
	return models, nil
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
