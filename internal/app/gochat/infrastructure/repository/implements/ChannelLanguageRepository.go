package implements

import (
	"gorm.io/gorm"
	"lecter/goserver/internal/app/gochat/domain/entity"
	"lecter/goserver/internal/app/gochat/infrastructure/model"

	"github.com/google/uuid"
)

type ChannelLanguageRepositoryImpl struct {
	Database *gorm.DB
}

func NewChannelLanguageRepositoryImpl(database *gorm.DB) ChannelLanguageRepositoryImpl {
	return ChannelLanguageRepositoryImpl{
		Database: database,
	}
}

func (r ChannelLanguageRepositoryImpl) Index(channelId uuid.UUID) ([]entity.ChannelLanguageEntity, error) {
	var models []model.ChannelLanguageModel
	err := r.Database.Where("channel_id = ?", channelId[:]).Find(&models).Error
	if err != nil {
		return nil, err
	}
	var entity []entity.ChannelLanguageEntity
	for _, model := range models {
		entity = append(entity, r.toEntity(model))
	}
	return entity, nil
}

func (r ChannelLanguageRepositoryImpl) Delete(channelId uuid.UUID) error {
	err := r.Database.Where("channel_id = ?", channelId[:]).Delete(&model.ChannelLanguageModel{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r ChannelLanguageRepositoryImpl) InsertAll(entities []entity.ChannelLanguageEntity) ([]entity.ChannelLanguageEntity, error) {
	var models []model.ChannelLanguageModel
	for _, entity := range entities {
		models = append(models, r.toModel(entity))
	}
	err := r.Database.Create(models).Error
	if err != nil {
		return nil, err
	}
	entities = make([]entity.ChannelLanguageEntity, len(models))
	for _, model := range models {
		entities = append(entities, r.toEntity(model))
	}
	return entities, nil
}

func (ChannelLanguageRepositoryImpl) toModel(entity entity.ChannelLanguageEntity) model.ChannelLanguageModel {
	return model.ChannelLanguageModel(entity)
}

func (ChannelLanguageRepositoryImpl) toEntity(model model.ChannelLanguageModel) entity.ChannelLanguageEntity {
	return entity.ChannelLanguageEntity(model)
}
