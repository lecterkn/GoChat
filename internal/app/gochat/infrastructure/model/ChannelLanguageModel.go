package model

import (
	"lecter/goserver/internal/app/gochat/domain/enum/language"

	"github.com/google/uuid"
)

type ChannelLanguageModel struct {
	ChannelId uuid.UUID         `json:"channel_id" gorm:"type:uuid"`
	Language  language.Language `json:"language"`
}

func (ChannelLanguageModel) TableName() string {
	return "channel_languages"
}
