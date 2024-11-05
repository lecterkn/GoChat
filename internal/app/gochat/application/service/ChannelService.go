package service

import (
	"lecter/goserver/internal/app/gochat/domain/entity"
	"lecter/goserver/internal/app/gochat/domain/enum/channel_permission"
	"lecter/goserver/internal/app/gochat/domain/repository"
	"lecter/goserver/internal/app/gochat/presentation/controller/response"
	"time"

	"github.com/google/uuid"
)

type ChannelService struct {
	ChannelRepository repository.ChannelRepository
}

func NewChannelService(channelRepository repository.ChannelRepository) ChannelService {
	return ChannelService{
		ChannelRepository: channelRepository,
	}
}

/*
 * チャンネル一覧を取得する
 */
func (cs ChannelService) GetChannels() ([]entity.ChannelEntity, *response.ErrorResponse) {
	entity, err := cs.ChannelRepository.Index()
	if err != nil {
		return nil, response.InternalError("database error")
	}
	return entity, nil
}

/*
 * チャンネルIDからチャンネルを取得する
 */
func (cs ChannelService) GetChannel(id uuid.UUID) (*entity.ChannelEntity, *response.ErrorResponse) {
	model, err := cs.ChannelRepository.Select(id)
	if err != nil {
		return nil, response.NotFoundError("channel not found")
	}
	return model, nil
}

/*
 * チャンネルを作成する
 */
func (cs ChannelService) CreateChannel(userId uuid.UUID, name string, permission channel_permission.ChannelPermission) (*entity.ChannelEntity, *response.ErrorResponse) {
	// ID生成
	id, err := uuid.NewV7()
	if err != nil {
		return nil, response.InternalError("failed to genearate uuid")
	}

	// チャンネル作成
	model := &entity.ChannelEntity{
		Id:         id,
		OwnerId:    userId,
		Name:       name,
		Permission: permission,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	model, err = cs.ChannelRepository.Create(*model)
	if err != nil {
		return nil, response.InternalError("failed to create channel")
	}
	return model, nil
}

/*
 * チャンネルを更新する
 */
func (cs ChannelService) UpdateChannel(channelId, userId uuid.UUID, name string, permission channel_permission.ChannelPermission) (*entity.ChannelEntity, *response.ErrorResponse) {
	// 更新対象のチャンネルを取得
	model, err := cs.ChannelRepository.Select(channelId)
	if err != nil {
		return nil, response.NotFoundError("channel not found")
	}
	if !IsChannelOwner(userId, *model) {
		return nil, response.ForbiddenError("permission error")
	}

	// チャンネルを更新
	model.Name = name
	model.Permission = permission
	model.UpdatedAt = time.Now()

	model, error := cs.ChannelRepository.Update(*model)
	if error != nil {
		return nil, response.InternalError("failed to update channel")
	}
	return model, nil
}

/*
 * チャンネルを削除する
 */
func (cs ChannelService) DeleteChannel(userId, channelId uuid.UUID) *response.ErrorResponse {
	model, err := cs.ChannelRepository.Select(channelId)
	if err != nil {
		return response.NotFoundError("channel not found")
	}
	if !IsChannelOwner(userId, *model) {
		return response.ForbiddenError("permission error")
	}
	model.Deleted = true
	_, err = cs.ChannelRepository.Update(*model)
	if err != nil {
		return response.InternalError("failed to delete channel")
	}
	return nil
}

func IsChannelOwner(userId uuid.UUID, channel entity.ChannelEntity) bool {
	return channel.OwnerId == userId
}
