package output

import (
	"time"

	"github.com/google/uuid"
)

type MessageOutput struct {
	Messages []MessageItem `json:"messages"`
}

type MessageItem struct {
	Id        uuid.UUID `json:"id"`
	ChannelId uuid.UUID `json:"channelId"`
	UserId    uuid.UUID `json:"userId"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
