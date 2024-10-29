package request

type UserProfileUpdateRequest struct {
	DisplayName string `json:"displayName" binding:"required,min=1,max=64"`
	Url         string `json:"url" binding:"required,min=1,max=1000"`
	Description string `json:"description" binding:"required,min=1,max=1000"`
}
