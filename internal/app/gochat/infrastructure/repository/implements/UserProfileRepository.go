package implemenst

import (
	"lecter/goserver/internal/app/gochat/domain/entity"
	"lecter/goserver/internal/app/gochat/infrastructure/db"
	"lecter/goserver/internal/app/gochat/infrastructure/model"

	"github.com/google/uuid"
)

type UserProfileRepositoryImpl struct{}

func (upr UserProfileRepositoryImpl) SelectByUserId(userId uuid.UUID) (*entity.UserProfileEntity, error) {
	var model model.UserProfileModel
	err := db.Database().Where("user_id = ?", userId[:]).First(&model).Error
	if err != nil {
		return nil, err
	}
	entity := upr.toEntity(model)
	return &entity, nil
}

func (upr UserProfileRepositoryImpl) Create(entity entity.UserProfileEntity) (*entity.UserProfileEntity, error) {
	model := upr.toModel(entity)
	err := db.Database().Create(&model).Error
	if err != nil {
		return nil, err
	}
	entity = upr.toEntity(model)
	return &entity, nil
}

func (upr UserProfileRepositoryImpl) Update(entity entity.UserProfileEntity) (*entity.UserProfileEntity, error) {
	model := upr.toModel(entity)
	err := db.Database().Save(&model).Error
	if err != nil {
		return nil, err
	}
	entity = upr.toEntity(model)
	return &entity, nil
}

func (upt UserProfileRepositoryImpl) toModel(entity entity.UserProfileEntity) (model.UserProfileModel) {
	return model.UserProfileModel{
		Id: entity.Id,
		UserId: entity.UserId,
		DisplayName: entity.DisplayName,
		Url: entity.Url,
		Description: entity.Description,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}

func (upt UserProfileRepositoryImpl) toEntity(model model.UserProfileModel) (entity.UserProfileEntity) {
	return entity.UserProfileEntity{
		Id: model.Id,
		UserId: model.UserId,
		DisplayName: model.DisplayName,
		Url: model.Url,
		Description: model.Description,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}