package model

import (
	"time"

	"github.com/google/uuid"
)

type UserProfileModel struct {
	Id          uuid.UUID `json:"id" gorm:"type:uuid"`
	UserId      uuid.UUID  `json:"userId" gorm:"type:uuid;column:user_id"`
	DisplayName string  `json:"displayName" gorm:"column:display_name"`
	Url         string  `json:"url"`
	Description string  `json:"description"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

func (UserProfileModel) TableName() string {
	return "user_profiles"
}