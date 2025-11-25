package pkgerrors

import (
	"net/http"
)

func BadRequest(msg string, err error) *AppError {
	return &AppError{
		Code:    http.StatusBadRequest,
		Message: msg,
		Detail:  err.Error(),
	}
}

func Unauthorized(msg string, err error) *AppError {
	return &AppError{
		Code:    http.StatusUnauthorized,
		Message: msg,
		Detail:  err.Error(),
	}
}

func InternalServerError(msg string, err error) *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: msg,
		Detail:  err.Error(),
	}
}
