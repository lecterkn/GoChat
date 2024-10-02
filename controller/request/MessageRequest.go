package request

import "github.com/google/uuid"

type MessageListRequest struct {
	LastId *uuid.UUID `json:"lastId"`
	Limit int `json:"limit" binding:"min=5,max=100"`
}

type MessageCreateRequest struct {
	Message string `json:"message" binding:"required,max=128"`
}

type MessageUpdateRequest struct {
	Message string `json:"message" binding:"required,max=128"`
}