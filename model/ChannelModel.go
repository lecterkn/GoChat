package model

import (
	"time"

	"github.com/google/uuid"
)

type ChannelModel struct {
	Id        uuid.UUID `json:"id" gorm:"type:uuid"`
	OwnerId uuid.UUID `json:"owner_id" gorm:"type:uuid;column:owner_id"`
	Name      string `json:"name"`
	Permission int16 `json:"permission"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (ChannelModel) TableName() string {
	return "channels"
}

var PermissionMap = map[int16]string {
	0: "readOnly",
	1: "writable",
}