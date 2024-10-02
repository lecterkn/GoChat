package model

import (
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	Id        uuid.UUID `gorm:"type:uuid" json:"id"`
	Name      string    `json:"name"`
	Password  []byte    `json:"-"`
	Language  int16     `json:"language"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

func (UserModel) TableName() string {
	return "users"
}
