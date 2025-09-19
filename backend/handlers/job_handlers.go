package handlers

import (
	"net/http"

	"job-board/backend/database"
	"job-board/backend/errors"
	"job-board/backend/logger"

	"github.com/gin-gonic/gin"
)

// GetJobs handles GET /api/jobs
func (h *Handler) GetJobs(c *gin.Context) {
	logger.Info("Fetching all jobs")
	jobs, err := h.jobService.GetAllJobs()
	if err != nil {
		logger.Error("Failed to fetch jobs", "error", err)
		AppErrorResponse(c, errors.WrapError(err, errors.ErrDatabaseQuery))
		return
	}
	logger.Info("Successfully fetched jobs", "count", len(jobs))
	SuccessResponse(c, http.StatusOK, jobs)
}

// GetJob handles GET /api/jobs/:id
func (h *Handler) GetJob(c *gin.Context) {
	id, err := parseID(c.Param("id"))
	if err != nil {
		logger.Warn("Invalid job ID provided", "id", c.Param("id"), "error", err)
		AppErrorResponse(c, errors.ErrInvalidInput)
		return
	}

	logger.Info("Fetching job by ID", "id", id)
	job, err := h.jobService.GetJobByID(id)
	if err != nil {
		logger.Error("Failed to fetch job", "id", id, "error", err)
		AppErrorResponse(c, errors.WrapError(err, errors.ErrJobNotFound))
		return
	}
	logger.Info("Successfully fetched job", "id", id, "title", job.Title)
	SuccessResponse(c, http.StatusOK, job)
}

// CreateJob handles POST /api/jobs
func (h *Handler) CreateJob(c *gin.Context) {
	var job database.Job
	if err := c.ShouldBindJSON(&job); err != nil {
		logger.Error("Failed to bind JSON for job creation", "error", err)
		AppErrorResponse(c, errors.WrapError(err, errors.ErrInvalidInput))
		return
	}

	logger.Info("Creating new job", "title", job.Title, "company", job.Company)

	// Sanitize input
	h.jobValidator.SanitizeJob(&job)

	// Validate required fields
	if err := h.jobValidator.ValidateJob(&job); err != nil {
		logger.Warn("Job validation failed", "error", err)
		AppErrorResponse(c, errors.WrapError(err, errors.ErrInvalidInput))
		return
	}

	if err := h.jobService.CreateJob(&job); err != nil {
		logger.Error("Failed to create job", "error", err)
		AppErrorResponse(c, errors.WrapError(err, errors.ErrJobCreationFailed))
		return
	}

	logger.Info("Successfully created job", "id", job.ID, "title", job.Title)
	SuccessResponse(c, http.StatusCreated, job)
}

// UpdateJob handles PUT /api/jobs/:id
func (h *Handler) UpdateJob(c *gin.Context) {
	id, err := parseID(c.Param("id"))
	if err != nil {
		AppErrorResponse(c, errors.ErrInvalidInput)
		return
	}

	var job database.Job
	if err := c.ShouldBindJSON(&job); err != nil {
		AppErrorResponse(c, errors.WrapError(err, errors.ErrInvalidInput))
		return
	}

	// Sanitize input
	h.jobValidator.SanitizeJob(&job)

	// Validate required fields
	if err := h.jobValidator.ValidateJob(&job); err != nil {
		AppErrorResponse(c, errors.WrapError(err, errors.ErrInvalidInput))
		return
	}

	if err := h.jobService.UpdateJob(id, &job); err != nil {
		AppErrorResponse(c, errors.WrapError(err, errors.ErrJobUpdateFailed))
		return
	}

	SuccessResponse(c, http.StatusOK, job)
}

// DeleteJob handles DELETE /api/jobs/:id
func (h *Handler) DeleteJob(c *gin.Context) {
	id, err := parseID(c.Param("id"))
	if err != nil {
		AppErrorResponse(c, errors.ErrInvalidInput)
		return
	}

	if err := h.jobService.DeleteJob(id); err != nil {
		AppErrorResponse(c, errors.WrapError(err, errors.ErrJobDeleteFailed))
		return
	}

	SuccessResponse(c, http.StatusOK, true)
}
