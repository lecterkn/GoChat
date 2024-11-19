package implements

import (
	"gorm.io/gorm"
	"lecter/goserver/internal/app/gochat/domain/entity"
	"lecter/goserver/internal/app/gochat/infrastructure/model"

	"github.com/google/uuid"
)

type ChannelRepositoryImpl struct {
	Database *gorm.DB
}

func NewChannelRepositoryImpl(database *gorm.DB) ChannelRepositoryImpl {
	return ChannelRepositoryImpl{
		Database: database,
	}
}

func (r ChannelRepositoryImpl) Index() ([]entity.ChannelEntity, error) {
	var models []model.ChannelModel
	err := r.Database.Where("deleted = FALSE").Find(&models).Error
	if err != nil {
		return nil, err
	}
	var entity []entity.ChannelEntity
	for _, model := range models {
		entity = append(entity, r.toEntity(model))
	}
	return entity, nil
}

func (r ChannelRepositoryImpl) Select(id uuid.UUID) (*entity.ChannelEntity, error) {
	var model model.ChannelModel
	err := r.Database.Where("id = ? AND deleted = FALSE", id[:]).First(&model).Error
	if err != nil {
		return nil, err
	}
	entity := r.toEntity(model)
	return &entity, nil
}

func (r ChannelRepositoryImpl) Create(entity entity.ChannelEntity) (*entity.ChannelEntity, error) {
	model := r.toModel(entity)
	err := r.Database.Create(&model).Error
	if err != nil {
		return nil, err
	}
	entity = r.toEntity(model)
	return &entity, nil
}

func (r ChannelRepositoryImpl) Update(entity entity.ChannelEntity) (*entity.ChannelEntity, error) {
	model := r.toModel(entity)
	err := r.Database.Where("deleted = FALSE").Save(&model).Error
	if err != nil {
		return nil, err
	}
	entity = r.toEntity(model)
	return &entity, nil
}

func (ChannelRepositoryImpl) toEntity(model model.ChannelModel) entity.ChannelEntity {
	return entity.ChannelEntity(model)
}

func (ChannelRepositoryImpl) toModel(entity entity.ChannelEntity) model.ChannelModel {
	return model.ChannelModel(entity)
}
