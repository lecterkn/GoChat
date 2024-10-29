package implemenst

import (
	"lecter/goserver/internal/app/gochat/domain/entity"
	"lecter/goserver/internal/app/gochat/infrastructure/db"
	"lecter/goserver/internal/app/gochat/infrastructure/model"

	"github.com/google/uuid"
)

type UserRepositoryImpl struct{}

func (ur UserRepositoryImpl) Insert(entity entity.UserEntity) (*entity.UserEntity, error) {
	model := ur.toModel(entity)
	err := db.Database().Create(&model).Error
	if err != nil {
		return nil, err
	}
	entity = ur.toEntity(model)
	return &entity, nil
}

func (ur UserRepositoryImpl) Select(id uuid.UUID) (*entity.UserEntity, error) {
	var model model.UserModel
	err := db.Database().Where("id = ?", id[:]).First(&model).Error
	if err != nil {
		return nil, err
	}
	entity := ur.toEntity(model)
	return &entity, nil
}

func (ur UserRepositoryImpl) SelectByName(name string) (*entity.UserEntity, error) {
	var model model.UserModel
	err := db.Database().Where("name = ?", name).First(&model).Error
	if err != nil {
		return nil, err
	}
	entity := ur.toEntity(model)
	return &entity, nil
}

func (ur UserRepositoryImpl) Update(entity entity.UserEntity) (*entity.UserEntity, error) {
	model := ur.toModel(entity)
	err := db.Database().Save(&model).Error
	if err != nil {
		return nil, err
	}
	entity = ur.toEntity(model)
	return &entity, nil
}

func (UserRepositoryImpl) toModel(entity entity.UserEntity) (model.UserModel) {
	return model.UserModel{
		Id: entity.Id,
		Name: entity.Name,
		Password: entity.Password,
		Language: entity.Language,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}

func (UserRepositoryImpl) toEntity(model model.UserModel) (entity.UserEntity) {
	return entity.UserEntity{
		Id: model.Id,
		Name: model.Name,
		Password: model.Password,
		Language: model.Language,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}