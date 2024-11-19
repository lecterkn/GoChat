package response

import (
	"time"

	"github.com/google/uuid"
)

type UserResponse struct {
	Id        uuid.UUID `gorm:"type:uuid" json:"id"`
	Name      string    `json:"name"`
	Language  string    `json:"language"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
}
