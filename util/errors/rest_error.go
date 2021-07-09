package errors

import (
	"errors"
	"net/http"
)


type RestError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewError(msg string) error {
	return errors.New(msg)
}

func NewBadRequestError(message, err string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error: err,
	}
}

func NewNotFoundError(message, err string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusNotFound,
		Error: err,
	}
}

func NewInternalServerError(message, err string) *RestError {
	return &RestError{
		Message: message,
		Status: http.StatusInternalServerError,
		Error: err,
	}
}