package request

import "github.com/google/uuid"

type MessageListRequestParam struct {
	Language *string    `form:"language"`
	LastId   *uuid.UUID `form:"lastId"`
	Limit    int        `form:"limit" binding:"min=5,max=100"`
}

type MessageCreateRequest struct {
	Message string `json:"message" binding:"required,max=128"`
}

type MessageUpdateRequest struct {
	Message string `json:"message" binding:"required,max=128"`
}
