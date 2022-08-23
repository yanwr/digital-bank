package exceptions

import (
	"net/http"
	"time"
)

type StandardError struct {
	Time          time.Time `json:"time"`
	Status        int       `json:"status"`
	Error_Message string    `json:"error_message"`
}

func ThrowInternalServerError(message string) *StandardError {
	return &StandardError{
		Time:          time.Now(),
		Status:        http.StatusInternalServerError,
		Error_Message: message,
	}
}

func ThrowBadRequestError(message string) *StandardError {
	return &StandardError{
		Time:          time.Now(),
		Status:        http.StatusBadRequest,
		Error_Message: message,
	}
}

func ThrowNotFoundError(message string) *StandardError {
	return &StandardError{
		Time:          time.Now(),
		Status:        http.StatusNotFound,
		Error_Message: message,
	}
}

func ThrowUnauthorizedError(message string) *StandardError {
	return &StandardError{
		Time:          time.Now(),
		Status:        http.StatusUnauthorized,
		Error_Message: message,
	}
}
