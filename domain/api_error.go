package domain

import (
	"errors"

	"github.com/surasithit/gin-basic-api/internal/helper/constants"
)

type ApiError struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("internal Server Error")
	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("your requested item is not found")
	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("your item already exist")
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadRequestInput = errors.New("given params or request body is not valid")
)

func ApplicationError(error string, message string) *ApiError {
	return &ApiError{
		Error:   error,
		Message: message,
	}
}

func BadRequestError(message string) *ApiError {
	if message == "" {
		message = ErrBadRequestInput.Error()
	}
	return &ApiError{
		Error:   constants.BadRequestError,
		Message: message,
	}
}

func ConflictError(message string) *ApiError {
	if message == "" {
		message = ErrConflict.Error()
	}
	return &ApiError{
		Error:   constants.BadRequestError,
		Message: message,
	}
}

func NotFoundError(message string) *ApiError {
	if message == "" {
		message = ErrNotFound.Error()
	}
	return &ApiError{
		Error:   constants.NotFoundError,
		Message: message,
	}
}

func ServerdError(message string) *ApiError {
	if message == "" {
		message = ErrInternalServerError.Error()
	}
	return &ApiError{
		Error:   constants.InternalServerError,
		Message: message,
	}
}
