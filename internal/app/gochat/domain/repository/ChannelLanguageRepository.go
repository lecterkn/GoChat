package repository

import (
	"lecter/goserver/internal/app/gochat/domain/entity"

	"github.com/google/uuid"
)

type ChannelLanguageRepository interface {
	Index(channelId uuid.UUID) ([]entity.ChannelLanguageEntity, error)
	Delete(channelId uuid.UUID) error
	InsertAll(entity []entity.ChannelLanguageEntity) ([]entity.ChannelLanguageEntity, error)
}
