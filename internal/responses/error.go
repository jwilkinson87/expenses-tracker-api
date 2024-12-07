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
	case http.StatusInternalServerError:
		errorResponse = NewErrorResponse("An error occurred", map[string]string{})
	case http.StatusBadRequest:
		errs, ok := errors.(validator.ValidationErrors)
		if ok {
			formattedErrors := util.FormatValidationMessages(obj, errs)
			errorResponse = NewErrorResponse("Invalid request", formattedErrors)
		} else {
			errorResponse = NewErrorResponse("Invalid request", map[string]string{})
		}
	case http.StatusUnauthorized:
		errorResponse = NewErrorResponse("Unauthorised", map[string]string{})
	}

	return errorResponse
}
