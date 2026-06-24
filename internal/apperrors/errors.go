package apperrors

import "net/http"

type AppError struct {
	Code    int
	Message string
	Err     error
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(code int, msg string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: msg,
		Err:     err,
	}
}

func NotFound(msg string, err error) *AppError {
	return NewAppError(http.StatusNotFound, msg, err)
}

func BadRequest(msg string, err error) *AppError {
	return NewAppError(http.StatusBadRequest, msg, err)
}

func InternalServerError(msg string, err error) *AppError {
	return NewAppError(http.StatusInternalServerError, msg, err)
}
