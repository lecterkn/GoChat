package request

// ユーザー作成用リクエスト
type UserCreateRequest struct {
	Name string `json:"name" binding:"required,min=1,max=20"`
	Url string `json:"url" binding:"required"`
}

// ユーザー更新用リクエスト
type UserUpdateRequest struct {
	Name string `json:"name" binding:"required,min=1,max=20"`
	Url string `json:"url" binding:"required"`
}