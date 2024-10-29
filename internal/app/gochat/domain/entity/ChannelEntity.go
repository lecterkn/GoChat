package entity

import (
	"lecter/goserver/internal/app/gochat/domain/enum/channel_permission"
	"lecter/goserver/internal/app/gochat/domain/enum/language"
	"lecter/goserver/internal/app/gochat/presentation/controller/response"
	"time"

	"github.com/google/uuid"
)

type ChannelEntity struct {
	Id         uuid.UUID                            `json:"id"`
	OwnerId    uuid.UUID                            `json:"owner_id"`
	Name       string                               `json:"name"`
	Permission channel_permission.ChannelPermission `json:"permission"`
	Deleted    bool                                 `json:"-"`
	CreatedAt  time.Time                            `json:"created_at"`
	UpdatedAt  time.Time                            `json:"updated_at"`
}

type ChannelLanguageEntity struct {
	ChannelId uuid.UUID         `json:"channel_id"`
	Language  language.Language `json:"language"`
}

func (model ChannelEntity) ToResponse() response.ChannelResponse {
	return response.ChannelResponse{
		Id:         model.Id,
		OwnerId:    model.OwnerId,
		Name:       model.Name,
		Permission: model.Permission.ToCode(),
		CreatedAt:  model.CreatedAt,
		UpdatedAt:  model.UpdatedAt,
	}
}
