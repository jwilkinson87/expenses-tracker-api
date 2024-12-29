package responses

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
