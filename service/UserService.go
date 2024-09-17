package service

import (
	"lecter/goserver/common"
	"lecter/goserver/controller/response"
	"lecter/goserver/model"
	"lecter/goserver/repository"

	"github.com/google/uuid"
)

type UserService struct{}

var userRepository = repository.UserRepository{}
/*
 * ユーザー一覧を取得
 */
func (us UserService) GetUsers() ([]model.UserModel, *response.ErrorResponse) {
	models, err := userRepository.Index()
	if err != nil {
		return nil, response.InternalError("failed to get users")
	}
	return models, nil
}

/*
 * IDからユーザーを取得
 */
func (us UserService) GetUser(id uuid.UUID) (*model.UserModel, *response.ErrorResponse) {
	model, err := userRepository.Select(id)
	if err != nil {
		return nil, response.NotFoundError("user not found")
	}
	return model, nil
}

/*
 * ユーザーを作成
 */
func (us UserService) CreateUser(name, password string) (*model.UserModel, *response.ErrorResponse) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, response.InternalError("failed to generate uuid")
	}

	var hashedPassword []byte
	hashedPassword, err = common.HashPassword(password)

	if err != nil {
		return nil, response.InternalError("failed to hash password")
	}

	model := &model.UserModel{
		Id: id,
		Name: name,
		Password: hashedPassword,
	}

	model, err = userRepository.Insert(*model)
	if err != nil {
		return nil, response.InternalError("failed to create user")
	}

	return model, nil
}

/*
 * ユーザーを更新
 */
func (us UserService) UpdateUser(userId uuid.UUID, name string, password string) (*model.UserModel, *response.ErrorResponse) {
	model, err := userRepository.Select(userId)

	if err != nil {
		return nil, response.NotFoundError("user not found")
	}

	var hashedPassword []byte
	hashedPassword, err = common.HashPassword(password)

	if err != nil {
		return nil, response.InternalError("failed to hash password")
	}

	model.Name = name
	model.Password = hashedPassword
	
	model, err = userRepository.Update(*model)
	if err != nil {
		return nil, response.InternalError("failed to update user")
	}
	return model, nil
}