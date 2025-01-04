package response

import "github.com/google/uuid"

type ChannelLanguageResponse struct {
	ChannelId uuid.UUID `json:"channelId"`
	Languages []string  `json:"languages"`
}
