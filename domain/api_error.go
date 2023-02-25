package domain

import (
	"net/http"
)

type ApiErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message,omitempty"`
}

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = NewApiError(http.StatusInternalServerError, "internal Server Error")
	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = NewApiError(http.StatusNotFound, "your requested item is not found")
	// ErrConflict will throw if the current action already exists
	ErrConflict = NewApiError(http.StatusBadRequest, "your item already exist")
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadRequest = NewApiError(http.StatusBadRequest, "given params or request body is not valid")
)

func NewApiError(statusCode int, message string) *ApiErrorResponse {
	return &ApiErrorResponse{
		StatusCode: statusCode,
		Message:    message,
	}
}
