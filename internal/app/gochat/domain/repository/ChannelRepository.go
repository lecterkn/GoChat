package repository

import (
	"lecter/goserver/internal/app/gochat/domain/entity"

	"github.com/google/uuid"
)

type ChannelRepository interface {
	Index() ([]entity.ChannelEntity, error)
	Select(id uuid.UUID) (*entity.ChannelEntity, error)
	Create(entity entity.ChannelEntity) (*entity.ChannelEntity, error)
	Update(entity entity.ChannelEntity) (*entity.ChannelEntity, error)
}
