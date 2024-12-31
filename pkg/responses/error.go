package responses

import "net/http"

type ErrorDetails struct {
	Message string            `json:"message"`
	Details map[string]string `json:"details,omitempty"`
}

type ErrorResponse struct {
	Status string `json:"status"`
	ErrorDetails
}

func NewErrorResponse(message string, details map[string]string) *ErrorResponse {
	return &ErrorResponse{
		Status: "error",
		ErrorDetails: ErrorDetails{
			Message: message,
			Details: details,
		},
	}
}

func NewErrorJsonHttpResponse(statusCode int, errors *map[string]string) *ErrorResponse {
	var errorResponse *ErrorResponse

	switch statusCode {
	case http.StatusBadRequest:
		errorResponse = NewErrorResponse("invalid request", *errors)
	case http.StatusUnauthorized:
		errorResponse = NewErrorResponse("unauthorised", nil)
	default:
		errorResponse = NewErrorResponse("an error occurred", nil)
	}

	return errorResponse
}
