package service

import (
	"lecter/goserver/internal/app/gochat/controller/response"
	"lecter/goserver/internal/app/gochat/model"
	"lecter/goserver/internal/app/gochat/repository"
	"time"

	"github.com/google/uuid"
)

type MessageService struct{}

var messageRepository = repository.MessageRepository{}

func (MessageService) GetChannels(userId, channelId uuid.UUID, lastId *uuid.UUID, limit int) (*[]model.MessageModel, *response.ErrorResponse) {
	channel, err := channelRepository.Select(channelId)
	if err != nil {
		return nil, response.NotFoundError("the channel does not exists")
	}
	if !isChannelReadable(*channel, userId) {
		return nil, response.ForbiddenError("permission error")
	}
	var lastCreatedAt *time.Time = nil
	if lastId != nil {
		lastMessage, err := messageRepository.Select(*lastId)
		if err != nil {
			return nil, response.ForbiddenError("last message does not exist")
		}
		lastCreatedAt = &lastMessage.CreatedAt
	}
	models, err := messageRepository.Index(channelId, lastId, lastCreatedAt, limit)
	if err != nil {
		return nil, response.InternalError("failed to get messages")
	}
	return models, nil
}

func (MessageService) CreateMessage(userId, channelId uuid.UUID, message string) (*model.MessageModel, *response.ErrorResponse) {
	channel, err := channelRepository.Select(channelId)
	if err != nil {
		return nil, response.NotFoundError("the channel does not exist")
	}
	if !isChannelWritable(*channel, userId) {
		return nil, response.ForbiddenError("permission error")
	}
	id, err := uuid.NewV7()
	if err != nil {
		return nil, response.InternalError("failed to generate uuid")
	}
	model := &model.MessageModel{
		Id:        id,
		ChannelId: channel.Id,
		UserId:    userId,
		Message:   message,
		Deleted:   false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	model, err = messageRepository.Create(*model)
	if err != nil {
		return nil, response.InternalError("failed to create message")
	}
	return model, nil
}

func (MessageService) UpdateMessage(userId, channelId, messageId uuid.UUID, message string) (*model.MessageModel, *response.ErrorResponse) {
	channel, err := channelRepository.Select(channelId)
	if err != nil {
		return nil, response.NotFoundError("the channel does not exist")
	}
	if !isChannelWritable(*channel, userId) {
		return nil, response.ForbiddenError("permission error")
	}
	model, err := messageRepository.Select(messageId)
	if err != nil {
		return nil, response.NotFoundError("the message does not exist")
	}
	model.Message = message
	model.UpdatedAt = time.Now()
	model, err = messageRepository.Update(*model)
	if err != nil {
		return nil, response.InternalError("failed to update message")
	}
	return model, nil
}

func (MessageService) DeleteMessage(userId, channelId, messageId uuid.UUID) *response.ErrorResponse {
	channel, err := channelRepository.Select(channelId)
	if err != nil {
		return response.NotFoundError("the channel does not exist")
	}
	model, err := messageRepository.Select(messageId)
	if err != nil {
		return response.NotFoundError("the message does not exist")
	}
	if !isMessageDeletable(*channel, *model, userId) {
		return response.ForbiddenError("permission error")
	}
	model.Deleted = true
	model.UpdatedAt = time.Now()

	_, err = messageRepository.Update(*model)
	if err != nil {
		return response.InternalError("failed to delete message")
	}
	return nil
}

func isChannelReadable(channel model.ChannelModel, userId uuid.UUID) bool {
	// チャンネルがプライベートの場合
	if channel.Permission == 2 && channel.OwnerId != userId {
		return false
	}
	return true
}

func isChannelWritable(channel model.ChannelModel, userId uuid.UUID) bool {
	// チャンネルが読み取り専用の場合
	if channel.Permission == 0 && channel.OwnerId != userId {
		return false
	}
	// チャンネルがプライベートの場合
	if channel.Permission == 2 && channel.OwnerId != userId {
		return false
	}
	return true
}

func isMessageDeletable(channel model.ChannelModel, message model.MessageModel, userId uuid.UUID) bool {
	if message.UserId != userId && channel.OwnerId != userId {
		return false
	}
	return true
}
