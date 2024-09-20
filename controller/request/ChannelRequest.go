package request

type ChannelCreateRequest struct {
	Name string `json:"name" binding:"required,min=1,max=32"`
	Permission *int16 `json:"permission" binding:"required"`
}

type ChannelUpdateRequest struct {
	Name string `json:"name" binding:"required,min=1,max=32"`
	Permission *int16 `json:"permission" binding:"required"`
}