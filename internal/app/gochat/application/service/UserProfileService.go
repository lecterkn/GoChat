package service

import (
	"lecter/goserver/internal/app/gochat/domain/entity"
	"lecter/goserver/internal/app/gochat/domain/repository"
	"lecter/goserver/internal/app/gochat/presentation/controller/response"
	"time"

	"github.com/google/uuid"
)

type UserProfileService struct {
	UserProfileRepository repository.UserProfileRepository
}

func NewUserProfileService(userProfileRepository repository.UserProfileRepository) UserProfileService {
	return UserProfileService{
		UserProfileRepository: userProfileRepository,
	}
}

/*
 * ユーザーのプロフィールがない場合は新規作成、ある場合は更新
 */
func (ups UserProfileService) UpdateUserProfile(userId uuid.UUID, displayName, url, description string) (*entity.UserProfileEntity, *response.ErrorResponse) {
	model, err := ups.UserProfileRepository.SelectByUserId(userId)
	if err != nil {
		return ups.createUserProfile(userId, displayName, url, description)
	}
	return ups.updateUserProfile(model, displayName, url, description)
}

/*
 * ユーザーIDからプロフィールを取得
 */
func (ups UserProfileService) SelectUserProfile(userId uuid.UUID) (*entity.UserProfileEntity, *response.ErrorResponse) {
	model, err := ups.UserProfileRepository.SelectByUserId(userId)
	if err != nil {
		return nil, response.NotFoundError("userProfiles not found")
	}
	return model, nil
}

func (ups UserProfileService) createUserProfile(userId uuid.UUID, displayName, url, description string) (*entity.UserProfileEntity, *response.ErrorResponse) {
	uuid, err := uuid.NewV7()
	if err != nil {
		return nil, response.InternalError("failed to generate id")
	}
	model := &entity.UserProfileEntity{
		Id:          uuid,
		UserId:      userId,
		DisplayName: displayName,
		Url:         url,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	model, err = ups.UserProfileRepository.Create(*model)
	if err != nil {
		return nil, response.InternalError("failed to insert userProfiles")
	}
	return model, nil
}

func (ups UserProfileService) updateUserProfile(model *entity.UserProfileEntity, displayName, url, description string) (*entity.UserProfileEntity, *response.ErrorResponse) {
	var err error
	model.DisplayName = displayName
	model.Url = url
	model.Description = description
	model.UpdatedAt = time.Now()
	model, err = ups.UserProfileRepository.Update(*model)
	if err != nil {
		return nil, response.InternalError("failed to update userProfiles")
	}
	return model, nil
}
