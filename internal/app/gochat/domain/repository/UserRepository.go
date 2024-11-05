package repository

import (
	"github.com/google/uuid"
	"lecter/goserver/internal/app/gochat/domain/entity"
)

type UserRepository interface {
	Select(id uuid.UUID) (*entity.UserEntity, error)
	SelectByName(name string) (*entity.UserEntity, error)
	Insert(entity entity.UserEntity) (*entity.UserEntity, error)
	Update(entity entity.UserEntity) (*entity.UserEntity, error)
}
