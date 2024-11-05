package implements

import (
	"gorm.io/gorm"
	"lecter/goserver/internal/app/gochat/domain/entity"
	"lecter/goserver/internal/app/gochat/infrastructure/model"

	"github.com/google/uuid"
)

type UserProfileRepositoryImpl struct {
	Database *gorm.DB
}

func NewUserProfileRepositoryImpl(database *gorm.DB) UserProfileRepositoryImpl {
	return UserProfileRepositoryImpl{
		Database: database,
	}
}

func (upr UserProfileRepositoryImpl) SelectByUserId(userId uuid.UUID) (*entity.UserProfileEntity, error) {
	var model model.UserProfileModel
	err := upr.Database.Where("user_id = ?", userId[:]).First(&model).Error
	if err != nil {
		return nil, err
	}
	entity := upr.toEntity(model)
	return &entity, nil
}

func (upr UserProfileRepositoryImpl) Create(entity entity.UserProfileEntity) (*entity.UserProfileEntity, error) {
	model := upr.toModel(entity)
	err := upr.Database.Create(&model).Error
	if err != nil {
		return nil, err
	}
	entity = upr.toEntity(model)
	return &entity, nil
}

func (upr UserProfileRepositoryImpl) Update(entity entity.UserProfileEntity) (*entity.UserProfileEntity, error) {
	model := upr.toModel(entity)
	err := upr.Database.Save(&model).Error
	if err != nil {
		return nil, err
	}
	entity = upr.toEntity(model)
	return &entity, nil
}

func (upt UserProfileRepositoryImpl) toModel(entity entity.UserProfileEntity) model.UserProfileModel {
	return model.UserProfileModel(entity)
}

func (upt UserProfileRepositoryImpl) toEntity(model model.UserProfileModel) entity.UserProfileEntity {
	return entity.UserProfileEntity(model)
}
