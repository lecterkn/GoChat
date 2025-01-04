package model

import (
	"time"

	"github.com/google/uuid"
)

type UserProfileModel struct {
	Id          uuid.UUID `gorm:"type:uuid"`
	UserId      uuid.UUID `gorm:"type:uuid;column:user_id"`
	DisplayName string    `gorm:"column:display_name"`
	Url         string
	Description string
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (UserProfileModel) TableName() string {
	return "user_profiles"
}
