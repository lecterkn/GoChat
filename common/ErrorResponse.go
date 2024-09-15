package common

type ErrorResponse struct {
	Code string `json:"code"`
	Message string `json:"message"`
}

func ValidationErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		"400",
		message,
	}
}

func InternalErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		"500",
		message,
	}
}

func NotFoundErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		"404",
		message,
	}
}