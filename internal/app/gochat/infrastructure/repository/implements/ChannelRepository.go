package implemenst

import (
	"lecter/goserver/internal/app/gochat/domain/entity"
	"lecter/goserver/internal/app/gochat/infrastructure/db"
	"lecter/goserver/internal/app/gochat/infrastructure/model"

	"github.com/google/uuid"
)

type ChannelRepositoryImpl struct{}

func (r ChannelRepositoryImpl) Index() ([]entity.ChannelEntity, error) {
	var models []model.ChannelModel
	err := db.Database().Where("deleted = FALSE").Find(&models).Error
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
	err := db.Database().Where("id = ? AND deleted = FALSE", id[:]).First(&model).Error
	if err != nil {
		return nil, err
	}
	entity := r.toEntity(model)
	return &entity, nil
}

func (r ChannelRepositoryImpl) Create(entity entity.ChannelEntity) (*entity.ChannelEntity, error) {
	model := r.toModel(entity)
	err := db.Database().Create(&model).Error
	if err != nil {
		return nil, err
	}
	entity = r.toEntity(model)
	return &entity, nil
}

func (r ChannelRepositoryImpl) Update(entity entity.ChannelEntity) (*entity.ChannelEntity, error) {
	model := r.toModel(entity)
	err := db.Database().Where("deleted = FALSE").Save(&model).Error
	if err != nil {
		return nil, err
	}
	entity = r.toEntity(model)
	return &entity, nil
}

func (ChannelRepositoryImpl) toEntity(model model.ChannelModel) (entity.ChannelEntity) {
	return entity.ChannelEntity{
		Id: model.Id,
		OwnerId: model.OwnerId,
		Name: model.Name,
		Permission: model.Permission,
		Deleted: model.Deleted,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func (ChannelRepositoryImpl) toModel(entity entity.ChannelEntity) (model.ChannelModel) {
	return model.ChannelModel{
		Id: entity.Id,
		OwnerId: entity.OwnerId,
		Name: entity.Name,
		Permission: entity.Permission,
		Deleted: entity.Deleted,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}