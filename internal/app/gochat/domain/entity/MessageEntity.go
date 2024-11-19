package entity

import (
	"time"

	"github.com/google/uuid"
)

type MessageEntity struct {
	Id        uuid.UUID `json:"id"`
	ChannelId uuid.UUID `json:"channelId"`
	UserId    uuid.UUID `json:"userId"`
	Message   string    `json:"message"`
	Deleted   bool      `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type TranslatedMessageEntity struct {
	Id             uuid.UUID `json:"id"`
	ChannelId      uuid.UUID `json:"channelId"`
	UserId         uuid.UUID `json:"userId"`
	MessageContent string    `json:"messageContent"`
	Deleted        bool      `json:"-"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type MessageJapaneseEntity struct {
	ChannelId uuid.UUID `json:"channelId"`
	MessageId uuid.UUID `json:"messageId"`
	Content   string    `json:"content"`
}

type MessageChineseEntity struct {
	ChannelId uuid.UUID `json:"channelId"`
	MessageId uuid.UUID `json:"messageId"`
	Content   string    `json:"content"`
}

type MessageEnglishEntity struct {
	ChannelId uuid.UUID `json:"channelId"`
	MessageId uuid.UUID `json:"messageId"`
	Content   string    `json:"content"`
}
