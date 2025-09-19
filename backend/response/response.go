package response

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Response represents a standard API response
type Response struct {
	Success   bool        `json:"success"`
	Data      interface{} `json:"data,omitempty"`
	Error     *ErrorInfo  `json:"error,omitempty"`
	Meta      *MetaInfo   `json:"meta,omitempty"`
	Timestamp string      `json:"timestamp"`
}

// ErrorInfo represents error information in the response
type ErrorInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// MetaInfo represents metadata in the response
type MetaInfo struct {
	Page       int   `json:"page,omitempty"`
	PageSize   int   `json:"pageSize,omitempty"`
	Total      int64 `json:"total,omitempty"`
	TotalPages int   `json:"totalPages,omitempty"`
}

// ResponseBuilder helps build consistent API responses
type ResponseBuilder struct {
	response *Response
}

// NewResponseBuilder creates a new response builder
func NewResponseBuilder() *ResponseBuilder {
	return &ResponseBuilder{
		response: &Response{
			Success:   true,
			Timestamp: time.Now().Format(time.RFC3339),
		},
	}
}

// WithData sets the data in the response
func (rb *ResponseBuilder) WithData(data interface{}) *ResponseBuilder {
	rb.response.Data = data
	return rb
}

// WithError sets the error in the response
func (rb *ResponseBuilder) WithError(code int, message, details string) *ResponseBuilder {
	rb.response.Success = false
	rb.response.Error = &ErrorInfo{
		Code:    code,
		Message: message,
		Details: details,
	}
	return rb
}

// WithMeta sets the metadata in the response
func (rb *ResponseBuilder) WithMeta(page, pageSize int, total int64) *ResponseBuilder {
	rb.response.Meta = &MetaInfo{
		Page:       page,
		PageSize:   pageSize,
		Total:      total,
		TotalPages: int((total + int64(pageSize) - 1) / int64(pageSize)),
	}
	return rb
}

// Build returns the built response
func (rb *ResponseBuilder) Build() *Response {
	return rb.response
}

// Send sends the response to the client
func (rb *ResponseBuilder) Send(c *gin.Context, statusCode int) {
	c.JSON(statusCode, rb.Build())
}

// SuccessResponse sends a success response
func SuccessResponse(c *gin.Context, statusCode int, data interface{}) {
	NewResponseBuilder().
		WithData(data).
		Send(c, statusCode)
}

// ErrorResponse sends an error response
func ErrorResponse(c *gin.Context, statusCode int, message, details string) {
	NewResponseBuilder().
		WithError(statusCode, message, details).
		Send(c, statusCode)
}

// PaginatedResponse sends a paginated response
func PaginatedResponse(c *gin.Context, statusCode int, data interface{}, page, pageSize int, total int64) {
	NewResponseBuilder().
		WithData(data).
		WithMeta(page, pageSize, total).
		Send(c, statusCode)
}

// ValidationErrorResponse sends a validation error response
func ValidationErrorResponse(c *gin.Context, message, details string) {
	ErrorResponse(c, http.StatusBadRequest, message, details)
}

// NotFoundResponse sends a not found response
func NotFoundResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusNotFound, message, "")
}

// InternalServerErrorResponse sends an internal server error response
func InternalServerErrorResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusInternalServerError, message, "")
}

// UnauthorizedResponse sends an unauthorized response
func UnauthorizedResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusUnauthorized, message, "")
}

// ForbiddenResponse sends a forbidden response
func ForbiddenResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusForbidden, message, "")
}
