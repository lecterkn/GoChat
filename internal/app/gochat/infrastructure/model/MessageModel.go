package model

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type MessageModel struct {
	Id        uuid.UUID `json:"id"`
	ChannelId uuid.UUID `json:"channelId"`
	UserId    uuid.UUID `json:"userId"`
	Message   string    `json:"message"`
	Deleted   bool      `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (MessageModel) TableName() string {
	return "messages"
}

func (mm MessageModel) MarshalBinary() ([]byte, error) {
	return json.Marshal(mm)
}

type MessageJapaneseModel struct {
	ChannelId uuid.UUID `json:"channelId"`
	MessageId uuid.UUID `json:"messageId"`
	Content   string    `json:"content"`
}

func (MessageJapaneseModel) TableName() string {
	return "message_japanese_contents"
}

type MessageChineseModel struct {
	ChannelId uuid.UUID `json:"channelId"`
	MessageId uuid.UUID `json:"messageId"`
	Content   string    `json:"content"`
}

func (MessageChineseModel) TableName() string {
	return "message_chinese_contents"
}

type MessageEnglishModel struct {
	ChannelId uuid.UUID `json:"channelId"`
	MessageId uuid.UUID `json:"messageId"`
	Content   string    `json:"content"`
}

func (MessageEnglishModel) TableName() string {
	return "message_english_contents"
}
