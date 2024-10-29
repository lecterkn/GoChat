package model

import (
	"lecter/goserver/internal/app/gochat/domain/enum/channel_permission"
	"time"

	"github.com/google/uuid"
)

type ChannelModel struct {
	Id         uuid.UUID                            `gorm:"type:uuid"`
	OwnerId    uuid.UUID                            `gorm:"type:uuid;column:owner_id"`
	Name       string                               
	Permission channel_permission.ChannelPermission 
	Deleted    bool                                 
	CreatedAt  time.Time                            `gorm:"column:created_at"`
	UpdatedAt  time.Time                            `gorm:"column:updated_at"`
}

func (ChannelModel) TableName() string {
	return "channels"
}

