package response

import "net/http"

// エラー時のレスポンスボディ
type ErrorResponse struct {
	Code string `json:"code"`
	Message string `json:"message"`
}

func ValidationErrorResponse(message string) (int, ErrorResponse) {
	return http.StatusBadRequest, 
		ErrorResponse{
			"400",
			message,
		}
}

func InternalErrorResponse(message string) (int, ErrorResponse) {
	return http.StatusInternalServerError,
		ErrorResponse{
			"500",
			message,
		}
}

func NotFoundErrorResponse(message string) (int, ErrorResponse) {
	return http.StatusNotFound,
		ErrorResponse{
			"404",
			message,
		}
}