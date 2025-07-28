package response

type SuccessResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"error"`
}

func NewSuccessResponse(data interface{}) SuccessResponse {
	return SuccessResponse{
		Status: "success",
		Data:   data,
	}
}

func NewErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		Status:  "error",
		Message: message,
	}
}