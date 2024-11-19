package implements

import (
	"gorm.io/gorm"
	"lecter/goserver/internal/app/gochat/domain/entity"
	"lecter/goserver/internal/app/gochat/domain/enum/language"
	"lecter/goserver/internal/app/gochat/infrastructure/model"

	"github.com/google/uuid"
)

type MessageRepositoryImpl struct {
	Database *gorm.DB
}

func NewMessageRepositoryImpl(database *gorm.DB) MessageRepositoryImpl {
	return MessageRepositoryImpl{
		Database: database,
	}
}

func (mr MessageRepositoryImpl) Index(
	channelId uuid.UUID,
	lastMessage *entity.MessageEntity,
	limit int) ([]entity.MessageEntity, error) {
	var models []model.MessageModel
	var err error
	if lastMessage == nil {
		err = mr.Database.
			Where("channel_id = ?", channelId[:]).
			Where("deleted = FALSE").
			Order("created_at DESC, id").
			Limit(limit).
			Find(&models).Error
	} else {
		err = mr.Database.
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
	var entity []entity.MessageEntity
	for _, model := range models {
		entity = append(entity, mr.toEntity(model))
	}
	return entity, nil
}

func (mr MessageRepositoryImpl) TranslatedMessageIndex(
	channelId uuid.UUID,
	lastMessage *entity.MessageEntity,
	limit int,
	lang language.Language) ([]entity.TranslatedMessageEntity, error) {
	var models []model.TranslatedMessageModel
	var err error
	if lastMessage == nil {
		err = mr.Database.
			Select("messages.*, COALESCE("+lang.TableName()+".content, messages.message) AS message_content").
			Joins("LEFT JOIN "+lang.TableName()+" ON "+lang.TableName()+".channel_id = messages.channel_id AND "+lang.TableName()+".message_id = messages.id").
			Where("messages.channel_id = ?", channelId[:]).
			Where("messages.deleted = FALSE").
			Order("created_at DESC, id").
			Limit(limit).
			Find(&models).Error
	} else {
		err = mr.Database.
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
	var entity []entity.TranslatedMessageEntity
	for _, model := range models {
		entity = append(entity, mr.toTranslatedEntity(model))
	}
	return entity, nil
}

func (r MessageRepositoryImpl) Select(id uuid.UUID) (*entity.MessageEntity, error) {
	var model model.MessageModel
	err := r.Database.Where("id = ? AND deleted = FALSE", id[:]).First(&model).Error
	if err != nil {
		return nil, err
	}
	entity := r.toEntity(model)
	return &entity, nil
}

func (r MessageRepositoryImpl) Create(entity entity.MessageEntity) (*entity.MessageEntity, error) {
	model := r.toModel(entity)
	err := r.Database.Create(&model).Error
	if err != nil {
		return nil, err
	}
	entity = r.toEntity(model)
	return &entity, nil
}

func (r MessageRepositoryImpl) Update(entity entity.MessageEntity) (*entity.MessageEntity, error) {
	model := r.toModel(entity)
	err := r.Database.Where("deleted = FALSE").Save(&model).Error
	if err != nil {
		return nil, err
	}
	entity = r.toEntity(model)
	return &entity, nil
}

func (MessageRepositoryImpl) toModel(entity entity.MessageEntity) model.MessageModel {
	return model.MessageModel(entity)
}

func (MessageRepositoryImpl) toEntity(model model.MessageModel) entity.MessageEntity {
	return entity.MessageEntity(model)
}

func (MessageRepositoryImpl) toTranslatedEntity(model model.TranslatedMessageModel) entity.TranslatedMessageEntity {
	return entity.TranslatedMessageEntity(model)
}
