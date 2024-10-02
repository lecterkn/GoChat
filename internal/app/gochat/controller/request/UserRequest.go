package request

// ユーザー作成用リクエスト
type UserCreateRequest struct {
	Name     string `json:"name" binding:"required,min=1,max=20"`
	Password string `json:"password" binding:"required,min=6,max=32"`
}

// ユーザー更新用リクエスト
type UserUpdateRequest struct {
	Name     string `json:"name" binding:"required,min=1,max=20"`
	Password string `json:"password" binding:"required,min=6,max=32"`
}
