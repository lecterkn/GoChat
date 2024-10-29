package service

import (
	"encoding/json"
	"fmt"
	"lecter/goserver/internal/app/gochat/application/service/output"
	"lecter/goserver/internal/app/gochat/domain/entity"
	"lecter/goserver/internal/app/gochat/domain/enum/language"
	"lecter/goserver/internal/app/gochat/domain/repository"
	"lecter/goserver/internal/app/gochat/presentation/controller/response"
	"time"

	"github.com/google/uuid"
)

type MessageService struct{
	MessageRepository repository.MessageRepository
  	ChannelRepository repository.ChannelRepository
	MessageDomainService MessageDomainService
	RedisService RedisService
}

func (ms MessageService) GetMessages(userId, channelId uuid.UUID, lastId *uuid.UUID, limit int, langCode *string) (*output.MessageOutput, *response.ErrorResponse) {
	// 言語を取得
	var lang *language.Language = nil
	if langCode != nil {
		langu, err := language.GetLanguageFromCode(*langCode)
		if err != nil {
			return nil, response.ValidationError("invalid language code")
		}
		lang = &langu
	}
	// チャンネルを取得
	channel, err := ms.ChannelRepository.Select(channelId)
	if err != nil {
		return nil, response.NotFoundError("the channel does not exists")
	}
	// 権限確認
	if !isChannelReadable(*channel, userId) {
		return nil, response.ForbiddenError("permission error")
	}
	// 原文を取得
	if lang == nil {
		// メッセージを一覧取得
		models, error := ms.MessageDomainService.GetOriginalMessage(channelId, lastId, limit)
		if error != nil {
			return nil, error
		}
		messageOutput := output.MessageOutput{
			Messages: []output.MessageItem{},
		}
		for _, message := range models {
			messageOutput.Messages = append(messageOutput.Messages, output.MessageItem{
				Id:        message.Id,
				ChannelId: message.ChannelId,
				UserId:    message.UserId,
				Message:   message.Message,
				CreatedAt: message.CreatedAt,
				UpdatedAt: message.UpdatedAt,
			})
		}
		return &messageOutput, nil
		// 翻訳されたメッセージを取得
	} else {
		// メッセージを一覧取得
		models, error := ms.MessageDomainService.GetTranslatedMessage(channelId, lastId, limit, *lang)
		if error != nil {
			return nil, error
		}
		messageOutput := output.MessageOutput{
			Messages: []output.MessageItem{},
		}
		for _, message := range models {
			messageOutput.Messages = append(messageOutput.Messages, output.MessageItem{
				Id:        message.Id,
				ChannelId: message.ChannelId,
				UserId:    message.UserId,
				Message:   message.MessageContent,
				CreatedAt: message.CreatedAt,
				UpdatedAt: message.UpdatedAt,
			})
		}
		return &messageOutput, nil
	}
}

func (ms MessageService) CreateMessage(userId, channelId uuid.UUID, message string) (*entity.MessageEntity, *response.ErrorResponse) {
	channel, err := ms.ChannelRepository.Select(channelId)
	if err != nil {
		return nil, response.NotFoundError("the channel does not exist")
	}
	if !isChannelWritable(*channel, userId) {
		return nil, response.ForbiddenError("permission error")
	}
	id, err := uuid.NewV7()
	if err != nil {
		return nil, response.InternalError("failed to generate uuidms ")
	}
	model := &entity.MessageEntity{
		Id:        id,
		ChannelId: channel.Id,
		UserId:    userId,
		Message:   message,
		Deleted:   false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	model, err = ms.MessageRepository.Create(*model)
	if err != nil {
		return nil, response.InternalError("failed to create message")
	}
	// jsonに変換
	messageJson, err := json.Marshal(model)
	if err != nil {
		fmt.Println("failed to unmarshal message")
	}
	// redisにパブリッシュ
	_, err = ms.RedisService.Publish(Broadcast, RedisMessage{
		SrcUser: model.UserId,
		Event:   AddMessage,
		Message: string(messageJson),
	})
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("failed to publish message")
	}
	return model, nil
}

func (ms MessageService) UpdateMessage(userId, channelId, messageId uuid.UUID, message string) (*entity.MessageEntity, *response.ErrorResponse) {
	channel, err := ms.ChannelRepository.Select(channelId)
	if err != nil {
		return nil, response.NotFoundError("the channel does not exist")
	}
	if !isChannelWritable(*channel, userId) {
		return nil, response.ForbiddenError("permission error")
	}
	model, err := ms.MessageRepository.Select(messageId)
	if err != nil {
		return nil, response.NotFoundError("the message does not exist")
	}
	model.Message = message
	model.UpdatedAt = time.Now()
	model, err = ms.MessageRepository.Update(*model)
	if err != nil {
		return nil, response.InternalError("failed to update message")
	}
	// jsonに変換
	messageJson, err := json.Marshal(model)
	if err != nil {
		fmt.Println("failed to unmarshal message")
	}
	// redisにパブリッシュ
	_, err = ms.RedisService.Publish(Broadcast, RedisMessage{
		SrcUser: model.UserId,
		Event:   UpdateMessage,
		Message: string(messageJson),
	})
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("failed to publish message")
	}
	return model, nil
}

func (ms MessageService) DeleteMessage(userId, channelId, messageId uuid.UUID) *response.ErrorResponse {
	channel, err := ms.ChannelRepository.Select(channelId)
	if err != nil {
		return response.NotFoundError("the channel does not exist")
	}
	model, err := ms.MessageRepository.Select(messageId)
	if err != nil {
		return response.NotFoundError("the message does not exist")
	}
	if !isMessageDeletable(*channel, *model, userId) {
		return response.ForbiddenError("permission error")
	}
	model.Deleted = true
	model.UpdatedAt = time.Now()

	_, err = ms.MessageRepository.Update(*model)
	if err != nil {
		return response.InternalError("failed to delete message")
	}
	// jsonに変換
	messageJson, err := json.Marshal(model)
	if err != nil {
		fmt.Println("failed to unmarshal message")
	}
	// redisにパブリッシュ
	_, err = ms.RedisService.Publish(Broadcast, RedisMessage{
		SrcUser: model.UserId,
		Event:   DeleteMessage,
		Message: string(messageJson),
	})
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("failed to publish message")
	}
	return nil
}

func isChannelReadable(channel entity.ChannelEntity, userId uuid.UUID) bool {
	// チャンネルがms プライベートの場合
	if channel.Permission == 2 && channel.OwnerId != userId {
		return false
	}
	return true
}

func isChannelWritable(channel entity.ChannelEntity, userId uuid.UUID) bool {
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

func isMessageDeletable(channel entity.ChannelEntity, message entity.MessageEntity, userId uuid.UUID) bool {
	if message.UserId != userId && channel.OwnerId != userId {
		return false
	}
	return true
}
