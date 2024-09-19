package service

import (
	"lecter/goserver/controller/response"
	"lecter/goserver/model"
	"lecter/goserver/repository"

	"github.com/google/uuid"
)

type UserProfileService struct{}

var userProfileRepository = repository.UserProfileRepository{}

/*
 * ユーザーのプロフィールがない場合は新規作成、ある場合は更新
 */
func (ups UserProfileService) UpdateUserProfile(userId uuid.UUID, displayName, url, description string) (*model.UserProfileModel, *response.ErrorResponse) {
	model, err := userProfileRepository.SelectByUserId(userId)
	if err != nil {
		return createUserProfile(userId, displayName, url, description)
	}
	return updateUserProfile(model, displayName, url, description)
}

/*
 * ユーザーIDからプロフィールを取得
 */
func (ups UserProfileService) SelectUserProfile(userId uuid.UUID) (*model.UserProfileModel, *response.ErrorResponse) {
	model, err := userProfileRepository.SelectByUserId(userId)
	if err != nil {
		return nil, response.NotFoundError("userProfiles not found")
	}
	return model, nil
}

func createUserProfile(userId uuid.UUID, displayName, url, description string) (*model.UserProfileModel, *response.ErrorResponse) {
	uuid, err := uuid.NewV7()
	if err != nil {
		return nil, response.InternalError("failed to generate id")
	}
	model := &model.UserProfileModel {
		Id: uuid,
		UserId: userId,
		DisplayName: displayName,
		Url: url,
		Description: description,
	}
	model, err = userProfileRepository.Create(*model)
	if err != nil {
		return nil, response.Unauthorized("failed to insert userProfiles")
	}
	return model, nil
}

func updateUserProfile(model *model.UserProfileModel, displayName, url, description string) (*model.UserProfileModel, *response.ErrorResponse){
	var err error
	model.DisplayName = displayName
	model.Url = url
	model.Description = description
	model, err = userProfileRepository.Update(*model)
	if err != nil {
		return nil, response.InternalError("failed to update userProfiles")
	}
	return model, nil
}