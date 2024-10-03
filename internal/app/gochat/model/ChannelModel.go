package model

import (
	"lecter/goserver/internal/app/gochat/controller/response"
	"lecter/goserver/internal/app/gochat/enum/channel_permission"
	"time"

	"github.com/google/uuid"
)

type ChannelModel struct {
	Id         uuid.UUID                            `json:"id" gorm:"type:uuid"`
	OwnerId    uuid.UUID                            `json:"owner_id" gorm:"type:uuid;column:owner_id"`
	Name       string                               `json:"name"`
	Permission channel_permission.ChannelPermission `json:"permission"`
	Deleted    bool                                 `json:"-"`
	CreatedAt  time.Time                            `json:"created_at" gorm:"column:created_at"`
	UpdatedAt  time.Time                            `json:"updated_at" gorm:"column:updated_at"`
}

func (ChannelModel) TableName() string {
	return "channels"
}

func (model ChannelModel) ToResponse() response.ChannelResponse {
	return response.ChannelResponse{
		Id:         model.Id,
		OwnerId:    model.OwnerId,
		Name:       model.Name,
		Permission: model.Permission.ToCode(),
		CreatedAt:  model.CreatedAt,
		UpdatedAt:  model.UpdatedAt,
	}
}
