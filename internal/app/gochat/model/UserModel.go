package model

import (
	"lecter/goserver/internal/app/gochat/controller/response"
	"lecter/goserver/internal/app/gochat/enum/language"
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	Id        uuid.UUID         `gorm:"type:uuid" json:"id"`
	Name      string            `json:"name"`
	Password  []byte            `json:"-"`
	Language  language.Language `json:"language"`
	CreatedAt time.Time         `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time         `json:"updatedAt" gorm:"column:updated_at"`
}

func (UserModel) TableName() string {
	return "users"
}

func (model UserModel) ToResponse() response.UserResponse {
	return response.UserResponse{
		Id:        model.Id,
		Name:      model.Name,
		Language:  model.Language.ToCode(),
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}
