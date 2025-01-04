package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserProfileEntity struct {
	Id          uuid.UUID `json:"id"`
	UserId      uuid.UUID `json:"userId"`
	DisplayName string    `json:"displayName"`
	Url         string    `json:"url"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
