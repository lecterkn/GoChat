package model

import (
	"time"

	"github.com/google/uuid"
)

type MessageModel struct {
	Id uuid.UUID `json:"id" gorm:"type:uuid"`
	ChannelId uuid.UUID `json:"channelId" gorm:"type:uuid;column:channel_id"`
	UserId uuid.UUID `json:"userId" gorm:"type:uuid;column:user_id"`
	Message string `json:"message"`
	Deleted bool `json:"-"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

func (MessageModel) TableName() string {
	return "messages"
}