package response

import "net/http"

// エラー時のレスポンスボディ
type ErrorResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

func (errorResponse ErrorResponse) ToResponse() (int, ErrorResponse) {
	return errorResponse.Code, errorResponse
}

func ValidationError(message string) (*ErrorResponse) {
	return &ErrorResponse{
		http.StatusBadRequest,
		message,
	}
}

func InternalError(message string) (*ErrorResponse) {
	return &ErrorResponse{
		http.StatusInternalServerError,
		message,
	}
}

func NotFoundError(message string) (*ErrorResponse) {
	return &ErrorResponse{
		http.StatusNotFound,
		message,
	}
}