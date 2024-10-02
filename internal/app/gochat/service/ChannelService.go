package service

import (
	"lecter/goserver/internal/app/gochat/controller/response"
	"lecter/goserver/internal/app/gochat/model"
	"lecter/goserver/internal/app/gochat/repository"
	"time"

	"github.com/google/uuid"
)

type ChannelService struct{}

var channelRepository = repository.ChannelRepository{}

/*
 * チャンネル一覧を取得する
 */
func (ChannelService) GetChannels() ([]model.ChannelModel, *response.ErrorResponse) {
	models, err := channelRepository.Index()
	if err != nil {
		return nil, response.InternalError("database error")
	}
	return models, nil
}

/*
 * チャンネルIDからチャンネルを取得する
 */
func (ChannelService) GetChannel(id uuid.UUID) (*model.ChannelModel, *response.ErrorResponse) {
	model, err := channelRepository.Select(id)
	if err != nil {
		return nil, response.NotFoundError("channel not found")
	}
	return model, nil
}

/*
 * チャンネルを作成する
 */
func (ChannelService) CreateChannel(userId uuid.UUID, name string, permission int16) (*model.ChannelModel, *response.ErrorResponse) {
	// permissionのバリデーションチェック
	if _, exists := model.PermissionMap[permission]; !exists {
		return nil, response.ValidationError("invalid permission")
	}

	// ID生成
	id, err := uuid.NewV7()
	if err != nil {
		return nil, response.InternalError("failed to genearate uuid")
	}

	// チャンネル作成
	model := &model.ChannelModel{
		Id:         id,
		OwnerId:    userId,
		Name:       name,
		Permission: permission,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	model, err = channelRepository.Create(*model)
	if err != nil {
		return nil, response.InternalError("failed to create channel")
	}
	return model, nil
}

/*
 * チャンネルを更新する
 */
func (ChannelService) UpdateChannel(channelId, userId uuid.UUID, name string, permission int16) (*model.ChannelModel, *response.ErrorResponse) {
	// permisssionのバリデーションチェック
	if _, exists := model.PermissionMap[permission]; !exists {
		return nil, response.ValidationError("invalid permission")
	}

	// 更新対象のチャンネルを取得
	if _, exists := model.PermissionMap[permission]; !exists {
		return nil, response.ValidationError("invalid permission")
	}
	model, err := channelRepository.Select(channelId)
	if err != nil {
		return nil, response.NotFoundError("channel not found")
	}
	if model.OwnerId != userId {
		return nil, response.ForbiddenError("permission error")
	}

	// チャンネルを更新
	model.Name = name
	model.Permission = permission
	model.UpdatedAt = time.Now()

	model, error := channelRepository.Update(*model)
	if error != nil {
		return nil, response.InternalError("failed to update channel")
	}
	return model, nil
}

/*
 * チャンネルを削除する
 */
func (ChannelService) DeleteChannel(userId, channelId uuid.UUID) *response.ErrorResponse {
	model, err := channelRepository.Select(channelId)
	if err != nil {
		return response.NotFoundError("channel not found")
	}
	if model.OwnerId != userId {
		return response.ForbiddenError("permission error")
	}
	model.Deleted = true
	_, err = channelRepository.Update(*model)
	if err != nil {
		return response.InternalError("failed to delete channel")
	}
	return nil
}
