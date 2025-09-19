package handlers

import (
	"strconv"

	"job-board/backend/database"
	"job-board/backend/errors"
	"job-board/backend/logger"
	"job-board/backend/response"
	"job-board/backend/streaming"
	"job-board/backend/validation"

	"github.com/gin-gonic/gin"
)

// Handler struct holds all the services
type Handler struct {
	jobService     *database.JobService
	videoService   *database.VideoService
	jobValidator   *validation.JobValidator
	videoValidator *validation.VideoValidator
	videoStreamer  *streaming.VideoStreamer
}

// NewHandler creates a new handler instance
func NewHandler(jobService *database.JobService, videoService *database.VideoService, videoStreamer *streaming.VideoStreamer) *Handler {
	return &Handler{
		jobService:     jobService,
		videoService:   videoService,
		jobValidator:   validation.NewJobValidator(),
		videoValidator: validation.NewVideoValidator(),
		videoStreamer:  videoStreamer,
	}
}

// AppErrorResponse creates a standardized error response from AppError
func AppErrorResponse(c *gin.Context, appErr *errors.AppError) {
	logger.Error("API Error", "status", appErr.Code, "message", appErr.Message, "path", c.Request.URL.Path)
	response.ErrorResponse(c, appErr.Code, appErr.Message, "")
}

// SuccessResponse creates a standardized success response
func SuccessResponse(c *gin.Context, statusCode int, data interface{}) {
	logger.Info("API Success", "status", statusCode, "path", c.Request.URL.Path)
	response.SuccessResponse(c, statusCode, data)
}

// parseID parses a string ID parameter to uint
func parseID(idStr string) (uint, error) {
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}
