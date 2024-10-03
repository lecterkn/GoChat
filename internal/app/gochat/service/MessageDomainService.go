package service

import (
	"lecter/goserver/internal/app/gochat/controller/response"
	"lecter/goserver/internal/app/gochat/enum/language"
	"lecter/goserver/internal/app/gochat/model"

	"github.com/google/uuid"
)

type MessageDomainService struct{}

func (MessageDomainService) GetOriginalMessage(channelId uuid.UUID, lastMessageId *uuid.UUID, limit int) ([]model.MessageModel, *response.ErrorResponse) {
	// 1ページ目
	if lastMessageId == nil {
		models, err := messageRepository.Index(channelId, nil, limit)
		if err != nil {
			return nil, response.InternalError("failed to get messages")
		}
		return models, nil
		// 2ページ目以降
	} else {
		lastMessage, err := messageRepository.Select(*lastMessageId)
		if err != nil {
			return nil, response.ForbiddenError("last message does not exist")
		}
		models, err := messageRepository.Index(channelId, lastMessage, limit)
		if err != nil {
			return nil, response.InternalError("failed to get messages")
		}
		return models, nil
	}
}

func (MessageDomainService) GetTranslatedMessage(channelId uuid.UUID, lastMessageId *uuid.UUID, limit int, lang language.Language) ([]model.TranslatedMessageModel, *response.ErrorResponse) {
	// 1ページ目
	if lastMessageId == nil {
		models, err := messageRepository.TranslatedMessageIndex(channelId, nil, limit, lang)
		if err != nil {
			return nil, response.InternalError("failed to get messages")
		}
		return models, nil
		// 2ページ目以降
	} else {
		lastMessage, err := messageRepository.Select(*lastMessageId)
		if err != nil {
			return nil, response.ForbiddenError("last message does not exist")
		}
		models, err := messageRepository.TranslatedMessageIndex(channelId, lastMessage, limit, lang)
		if err != nil {
			return nil, response.InternalError("failed to get messages")
		}
		return models, nil
	}
}
