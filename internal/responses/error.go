package responses

import (
	"net/http"

	"example.com/expenses-tracker/internal/util"
	"github.com/go-playground/validator/v10"
)

type ErrorDetails struct {
	Message string            `json:"message"`
	Details map[string]string `json:"details,omitempty"`
}

type ErrorResponse struct {
	Status string `json:"status"`
	ErrorDetails
}

func NewErrorResponse(message string, details map[string]string) *ErrorResponse {
	if details == nil {
		details = map[string]string{}
	}

	return &ErrorResponse{
		Status: "error",
		ErrorDetails: ErrorDetails{
			Message: message,
			Details: details,
		},
	}
}

func NewErrorJsonHttpResponse(statusCode int, obj any, errors any) *ErrorResponse {
	var errorResponse *ErrorResponse

	switch statusCode {
	case http.StatusBadRequest:
		errs, ok := errors.(validator.ValidationErrors)
		if ok {
			formattedErrors := util.FormatValidationMessages(obj, errs)
			errorResponse = NewErrorResponse("invalid request", formattedErrors)
		} else {
			errorResponse = NewErrorResponse("invalid request", nil)
		}
	case http.StatusUnauthorized:
		errorResponse = NewErrorResponse("unauthorised", nil)
	default:
		errorResponse = NewErrorResponse("an error occurred", nil)
	}

	return errorResponse
}
