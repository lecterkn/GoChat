package repository

import (
	"lecter/goserver/internal/app/gochat/domain/entity"
	"github.com/google/uuid"
)

type UserRepository interface {
	Select(id uuid.UUID) (*entity.UserEntity, error)
	SelectByName(name string) (*entity.UserEntity, error)
	Insert(entity entity.UserEntity) (*entity.UserEntity, error)
	Update(entity entity.UserEntity) (*entity.UserEntity, error)
}