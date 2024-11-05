package implements

import (
	"gorm.io/gorm"
	"lecter/goserver/internal/app/gochat/domain/entity"
	"lecter/goserver/internal/app/gochat/infrastructure/model"

	"github.com/google/uuid"
)

type UserRepositoryImpl struct {
	Database *gorm.DB
}

func NewUserRepositoryImpl(database *gorm.DB) UserRepositoryImpl {
	return UserRepositoryImpl{
		Database: database,
	}
}

func (ur UserRepositoryImpl) Insert(entity entity.UserEntity) (*entity.UserEntity, error) {
	model := ur.toModel(entity)
	err := ur.Database.Create(&model).Error
	if err != nil {
		return nil, err
	}
	entity = ur.toEntity(model)
	return &entity, nil
}

func (ur UserRepositoryImpl) Select(id uuid.UUID) (*entity.UserEntity, error) {
	var model model.UserModel
	err := ur.Database.Where("id = ?", id[:]).First(&model).Error
	if err != nil {
		return nil, err
	}
	entity := ur.toEntity(model)
	return &entity, nil
}

func (ur UserRepositoryImpl) SelectByName(name string) (*entity.UserEntity, error) {
	var model model.UserModel
	err := ur.Database.Where("name = ?", name).First(&model).Error
	if err != nil {
		return nil, err
	}
	entity := ur.toEntity(model)
	return &entity, nil
}

func (ur UserRepositoryImpl) Update(entity entity.UserEntity) (*entity.UserEntity, error) {
	model := ur.toModel(entity)
	err := ur.Database.Save(&model).Error
	if err != nil {
		return nil, err
	}
	entity = ur.toEntity(model)
	return &entity, nil
}

func (UserRepositoryImpl) toModel(entity entity.UserEntity) model.UserModel {
	return model.UserModel(entity)
}

func (UserRepositoryImpl) toEntity(model model.UserModel) entity.UserEntity {
	return entity.UserEntity(model)
}
