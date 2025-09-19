package handlers

import (
	"net/http"

	"job-board/backend/database"
	"job-board/backend/errors"
	"job-board/backend/logger"

	"github.com/gin-gonic/gin"
)

// GetVideos handles GET /api/videos
func (h *Handler) GetVideos(c *gin.Context) {
	videos, err := h.videoService.GetAllVideos()
	if err != nil {
		AppErrorResponse(c, errors.WrapError(err, errors.ErrDatabaseQuery))
		return
	}
	SuccessResponse(c, http.StatusOK, videos)
}

// GetVideo handles GET /api/videos/:id
func (h *Handler) GetVideo(c *gin.Context) {
	id, err := parseID(c.Param("id"))
	if err != nil {
		AppErrorResponse(c, errors.ErrInvalidInput)
		return
	}

	video, err := h.videoService.GetVideoByID(id)
	if err != nil {
		AppErrorResponse(c, errors.WrapError(err, errors.ErrVideoNotFound))
		return
	}
	SuccessResponse(c, http.StatusOK, video)
}

// CreateVideo handles POST /api/videos
func (h *Handler) CreateVideo(c *gin.Context) {
	var video database.Video
	if err := c.ShouldBindJSON(&video); err != nil {
		AppErrorResponse(c, errors.WrapError(err, errors.ErrInvalidInput))
		return
	}

	// Sanitize input
	h.videoValidator.SanitizeVideo(&video)

	// Validate required fields
	if err := h.videoValidator.ValidateVideo(&video); err != nil {
		AppErrorResponse(c, errors.WrapError(err, errors.ErrInvalidInput))
		return
	}

	if err := h.videoService.CreateVideo(&video); err != nil {
		AppErrorResponse(c, errors.WrapError(err, errors.ErrVideoCreationFailed))
		return
	}

	SuccessResponse(c, http.StatusCreated, video)
}

// StreamVideo handles GET /video/:id
func (h *Handler) StreamVideo(c *gin.Context) {
	videoID := c.Param("id")

	logger.Info("Streaming video request", "video_id", videoID)

	err := h.videoStreamer.StreamVideo(c.Writer, c.Request, videoID)
	if err != nil {
		logger.Error("Failed to stream video", "video_id", videoID, "error", err)
		AppErrorResponse(c, errors.ErrVideoNotFound)
		return
	}

	logger.Info("Successfully streamed video", "video_id", videoID)
}
