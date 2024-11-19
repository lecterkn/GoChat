package repository

import (
	"lecter/goserver/internal/app/gochat/domain/entity"
	"lecter/goserver/internal/app/gochat/domain/enum/language"

	"github.com/google/uuid"
)

type MessageRepository interface {
	Index(id uuid.UUID, lastEntity *entity.MessageEntity, limit int) ([]entity.MessageEntity, error)
	TranslatedMessageIndex(
		channelId uuid.UUID,
		lastMessage *entity.MessageEntity,
		limit int,
		lang language.Language) ([]entity.TranslatedMessageEntity, error)
	Select(id uuid.UUID) (*entity.MessageEntity, error)
	Create(entity entity.MessageEntity) (*entity.MessageEntity, error)
	Update(entity entity.MessageEntity) (*entity.MessageEntity, error)
}
