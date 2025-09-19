package errors

import (
	"fmt"
	"net/http"
)

// AppError represents an application error
type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// Error implements the error interface
func (e *AppError) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("%s: %s", e.Message, e.Details)
	}
	return e.Message
}

// NewAppError creates a new application error
func NewAppError(code int, message string, details ...string) *AppError {
	err := &AppError{
		Code:    code,
		Message: message,
	}
	if len(details) > 0 {
		err.Details = details[0]
	}
	return err
}

// Predefined errors
var (
	// Database errors
	ErrDatabaseConnection = NewAppError(http.StatusInternalServerError, "Database connection failed")
	ErrDatabaseQuery      = NewAppError(http.StatusInternalServerError, "Database query failed")
	ErrRecordNotFound     = NewAppError(http.StatusNotFound, "Record not found")
	ErrRecordExists       = NewAppError(http.StatusConflict, "Record already exists")

	// Validation errors
	ErrInvalidInput  = NewAppError(http.StatusBadRequest, "Invalid input")
	ErrMissingField  = NewAppError(http.StatusBadRequest, "Required field is missing")
	ErrInvalidFormat = NewAppError(http.StatusBadRequest, "Invalid format")

	// Job errors
	ErrJobNotFound       = NewAppError(http.StatusNotFound, "Job not found")
	ErrJobCreationFailed = NewAppError(http.StatusInternalServerError, "Failed to create job")
	ErrJobUpdateFailed   = NewAppError(http.StatusInternalServerError, "Failed to update job")
	ErrJobDeleteFailed   = NewAppError(http.StatusInternalServerError, "Failed to delete job")

	// Video errors
	ErrVideoNotFound       = NewAppError(http.StatusNotFound, "Video not found")
	ErrVideoCreationFailed = NewAppError(http.StatusInternalServerError, "Failed to create video")
	ErrVideoStreamFailed   = NewAppError(http.StatusInternalServerError, "Failed to stream video")

	// Server errors
	ErrInternalServer     = NewAppError(http.StatusInternalServerError, "Internal server error")
	ErrServiceUnavailable = NewAppError(http.StatusServiceUnavailable, "Service unavailable")
)

// WrapError wraps an existing error with additional context
func WrapError(err error, appErr *AppError) *AppError {
	if err == nil {
		return appErr
	}
	return NewAppError(appErr.Code, appErr.Message, err.Error())
}
