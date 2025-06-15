package dtos

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func CustomErrorResponse(status int, message string) ErrorResponse {
	return ErrorResponse{
		Status:  status,
		Message: message,
	}
}
