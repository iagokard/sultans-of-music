package util

type ErrorResponse struct {
	Error   string `json:"error" example:"sample error message"`
	Details string `json:"details,omitempty" example:"error details"`
}

type MessageResponse struct {
	Message string `json:"message" example:"Sample status message"`
}
