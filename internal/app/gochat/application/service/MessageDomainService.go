package service

import (
	"lecter/goserver/internal/app/gochat/domain/entity"
	"lecter/goserver/internal/app/gochat/domain/enum/language"
	"lecter/goserver/internal/app/gochat/domain/repository"
	"lecter/goserver/internal/app/gochat/presentation/controller/response"

	"github.com/google/uuid"
)

type MessageDomainService struct {
	MessageRepository repository.MessageRepository
}

func NewMessageDomainService(messageRepository repository.MessageRepository) MessageDomainService {
	return MessageDomainService{
		MessageRepository: messageRepository,
	}
}

func (mds MessageDomainService) GetOriginalMessage(channelId uuid.UUID, lastMessageId *uuid.UUID, limit int) ([]entity.MessageEntity, *response.ErrorResponse) {
	// 1ページ目
	if lastMessageId == nil {
		entity, err := mds.MessageRepository.Index(channelId, nil, limit)
		if err != nil {
			return nil, response.InternalError("failed to get messages")
		}
		return entity, nil
		// 2ページ目以降
	} else {
		lastMessage, err := mds.MessageRepository.Select(*lastMessageId)
		if err != nil {
			return nil, response.ForbiddenError("last message does not exist")
		}
		entity, err := mds.MessageRepository.Index(channelId, lastMessage, limit)
		if err != nil {
			return nil, response.InternalError("failed to get messages")
		}
		return entity, nil
	}
}

func (mds MessageDomainService) GetTranslatedMessage(channelId uuid.UUID, lastMessageId *uuid.UUID, limit int, lang language.Language) ([]entity.TranslatedMessageEntity, *response.ErrorResponse) {
	// 1ページ目
	if lastMessageId == nil {
		entity, err := mds.MessageRepository.TranslatedMessageIndex(channelId, nil, limit, lang)
		if err != nil {
			return nil, response.InternalError("failed to get messages")
		}
		return entity, nil
		// 2ページ目以降
	} else {
		lastMessage, err := mds.MessageRepository.Select(*lastMessageId)
		if err != nil {
			return nil, response.ForbiddenError("last message does not exist")
		}
		entity, err := mds.MessageRepository.TranslatedMessageIndex(channelId, lastMessage, limit, lang)
		if err != nil {
			return nil, response.InternalError("failed to get messages")
		}
		return entity, nil
	}
}
