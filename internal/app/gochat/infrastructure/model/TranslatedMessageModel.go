package model

import (
	"time"

	"github.com/google/uuid"
)

type TranslatedMessageModel struct {
	Id             uuid.UUID `gorm:"type:uuid"`
	ChannelId      uuid.UUID `gorm:"type:uuid;column:channel_id"`
	UserId         uuid.UUID `gorm:"type:uuid;column:user_id"`
	MessageContent string    `gorm:"column:message_content"`
	Deleted        bool
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
}

func (TranslatedMessageModel) TableName() string {
	return "messages"
}
