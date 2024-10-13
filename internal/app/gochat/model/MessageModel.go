package model

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type MessageModel struct {
	Id        uuid.UUID `json:"id" gorm:"type:uuid"`
	ChannelId uuid.UUID `json:"channelId" gorm:"type:uuid;column:channel_id"`
	UserId    uuid.UUID `json:"userId" gorm:"type:uuid;column:user_id"`
	Message   string    `json:"message"`
	Deleted   bool      `json:"-"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

func (MessageModel) TableName() string {
	return "messages"
}

func (mm MessageModel) MarshalBinary() ([]byte, error) {
	return json.Marshal(mm)
}

type MessageJapaneseModel struct {
	ChannelId uuid.UUID `json:"channelId" gorm:"type:uuid;column:channel_id"`
	MessageId uuid.UUID `json:"messageId" gorm:"type:uuid;column:message_id"`
	Content   string    `json:"content"`
}

func (MessageJapaneseModel) TableName() string {
	return "message_japanese_contents"
}

type MessageChineseModel struct {
	ChannelId uuid.UUID `json:"channelId" gorm:"type:uuid;column:channel_id"`
	MessageId uuid.UUID `json:"messageId" gorm:"type:uuid;column:message_id"`
	Content   string    `json:"content"`
}

func (MessageChineseModel) TableName() string {
	return "message_chinese_contents"
}

type MessageEnglishModel struct {
	ChannelId uuid.UUID `json:"channelId" gorm:"type:uuid;column:channel_id"`
	MessageId uuid.UUID `json:"messageId" gorm:"type:uuid;column:message_id"`
	Content   string    `json:"content"`
}

func (MessageEnglishModel) TableName() string {
	return "message_english_contents"
}
