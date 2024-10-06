package response

import (
	"time"

	"github.com/google/uuid"
)

type ChannelResponse struct {
	Id         uuid.UUID `json:"id"`
	OwnerId    uuid.UUID `json:"owner_id"`
	Name       string    `json:"name"`
	Permission string    `json:"permission"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ChannelListResponse struct {
	List []ChannelResponse `json:"list"`
}