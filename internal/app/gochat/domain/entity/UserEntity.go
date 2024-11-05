package entity

import (
	"lecter/goserver/internal/app/gochat/domain/enum/language"
	"lecter/goserver/internal/app/gochat/presentation/controller/response"
	"time"

	"github.com/google/uuid"
)

type UserEntity struct {
	Id        uuid.UUID         `json:"id"`
	Name      string            `json:"name"`
	Password  []byte            `json:"-"`
	Language  language.Language `json:"language"`
	CreatedAt time.Time         `json:"createdAt"`
	UpdatedAt time.Time         `json:"updatedAt"`
}

func (model UserEntity) ToResponse() response.UserResponse {
	return response.UserResponse{
		Id:        model.Id,
		Name:      model.Name,
		Language:  model.Language.ToCode(),
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}
