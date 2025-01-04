package repository

import (
	"lecter/goserver/internal/app/gochat/domain/entity"

	"github.com/google/uuid"
)

type UserProfileRepository interface {
	SelectByUserId(userId uuid.UUID) (*entity.UserProfileEntity, error)
	Create(entity entity.UserProfileEntity) (*entity.UserProfileEntity, error)
	Update(entity entity.UserProfileEntity) (*entity.UserProfileEntity, error)
}
