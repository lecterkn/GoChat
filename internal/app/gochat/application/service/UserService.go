package service

import (
	"lecter/goserver/internal/app/gochat/common"
	"lecter/goserver/internal/app/gochat/domain/entity"
	"lecter/goserver/internal/app/gochat/domain/enum/language"
	"lecter/goserver/internal/app/gochat/domain/repository"
	"lecter/goserver/internal/app/gochat/presentation/controller/response"
	"time"

	"github.com/google/uuid"
)

type UserService struct{
	UserRepository repository.UserRepository
}

/*
 * IDからユーザーを取得
 */
func (us UserService) GetUser(id uuid.UUID) (*entity.UserEntity, *response.ErrorResponse) {
	entity, err := us.UserRepository.Select(id)
	if err != nil {
		return nil, response.NotFoundError("user not found")
	}
	return entity, nil
}

/*
 * ユーザー名からユーザー取得
 */
func (us UserService) GetUserByName(name string) (*entity.UserEntity, *response.ErrorResponse) {
	entity, err := us.UserRepository.SelectByName(name)
	if err != nil {
		return nil, response.NotFoundError("user not found")
	}
	return entity, nil
}

/*
 * ユーザーを作成
 */
func (us UserService) CreateUser(name, password string) (*entity.UserEntity, *response.ErrorResponse) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, response.InternalError("failed to generate uuid")
	}

	var hashedPassword []byte
	hashedPassword, err = common.HashPassword(password)

	if err != nil {
		return nil, response.InternalError("failed to hash password")
	}

	entity := &entity.UserEntity{
		Id:        id,
		Name:      name,
		Password:  hashedPassword,
		Language:  language.English,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	entity, err = us.UserRepository.Insert(*entity)
	if err != nil {
		return nil, response.InternalError("failed to create user")
	}

	return entity, nil
}

/*
 * ユーザーを更新
 */
func (us UserService) UpdateUser(userId uuid.UUID, name string, langCode language.Language) (*entity.UserEntity, *response.ErrorResponse) {
	entity, err := us.UserRepository.Select(userId)

	if err != nil {
		return nil, response.NotFoundError("user not found")
	}

	entity.Name = name
	entity.UpdatedAt = time.Now()

	entity, err = us.UserRepository.Update(*entity)
	if err != nil {
		return nil, response.InternalError("failed to update user")
	}
	return entity, nil
}
