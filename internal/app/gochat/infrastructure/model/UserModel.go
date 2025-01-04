package model

import (
	"lecter/goserver/internal/app/gochat/domain/enum/language"
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	Id        uuid.UUID `gorm:"type:uuid"`
	Name      string
	Password  []byte
	Language  language.Language
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (UserModel) TableName() string {
	return "users"
}
