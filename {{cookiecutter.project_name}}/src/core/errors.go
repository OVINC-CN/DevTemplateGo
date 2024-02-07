package core

import (
	"net/http"
)

type APIError struct {
	Status  int
	Message string
	Detail  *map[string]any
}

func (err *APIError) Error() string {
	return err.Message
}

func NewError(status int, message string, detail *map[string]any) *APIError {
	return &APIError{
		Status:  status,
		Message: message,
		Detail:  detail,
	}
}

var (
	LoginRequired = NewError(http.StatusUnauthorized, "login required", nil)
)
