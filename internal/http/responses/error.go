package responses

import (
	"net/http"

	httpResponse "example.com/expenses-tracker/pkg/responses"
	"github.com/go-playground/validator/v10"
)

func NewErrorJsonHttpResponse(statusCode int, obj any, errors any) *httpResponse.ErrorResponse {
	var errorResponse *httpResponse.ErrorResponse

	switch statusCode {
	case http.StatusBadRequest:
		errs, ok := errors.(validator.ValidationErrors)
		if ok {
			formattedErrors := FormatValidationMessages(obj, errs)
			errorResponse = httpResponse.NewErrorResponse("invalid request", formattedErrors)
		} else {
			errorResponse = httpResponse.NewErrorResponse("invalid request", nil)
		}
	case http.StatusUnauthorized:
		errorResponse = httpResponse.NewErrorResponse("unauthorised", nil)
	default:
		errorResponse = httpResponse.NewErrorResponse("an error occurred", nil)
	}

	return errorResponse
}
