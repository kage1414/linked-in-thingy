package database

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// JobService handles job-related database operations
type JobService struct {
	db *gorm.DB
}

// NewJobService creates a new JobService
func NewJobService(db *gorm.DB) *JobService {
	return &JobService{db: db}
}

// GetAllJobs retrieves all jobs from the database with optimized queries
func (s *JobService) GetAllJobs() ([]Job, error) {
	var jobs []Job
	err := s.db.Preload("Videos").Find(&jobs).Error
	return jobs, err
}

// GetJobsWithPagination retrieves jobs with pagination
func (s *JobService) GetJobsWithPagination(page, pageSize int) ([]Job, int64, error) {
	var jobs []Job
	var total int64

	// Count total records
	if err := s.db.Model(&Job{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * pageSize

	// Get jobs with pagination
	err := s.db.Preload("Videos").
		Offset(offset).
		Limit(pageSize).
		Order("created_at DESC").
		Find(&jobs).Error

	return jobs, total, err
}

// GetJobsByCompany retrieves jobs by company name
func (s *JobService) GetJobsByCompany(company string) ([]Job, error) {
	var jobs []Job
	err := s.db.Where("company ILIKE ?", "%"+company+"%").
		Preload("Videos").
		Find(&jobs).Error
	return jobs, err
}

// GetJobsByLocation retrieves jobs by location
func (s *JobService) GetJobsByLocation(location string) ([]Job, error) {
	var jobs []Job
	err := s.db.Where("location ILIKE ?", "%"+location+"%").
		Preload("Videos").
		Find(&jobs).Error
	return jobs, err
}

// SearchJobs performs a full-text search on jobs
func (s *JobService) SearchJobs(query string) ([]Job, error) {
	var jobs []Job
	searchQuery := "%" + query + "%"
	err := s.db.Where("title ILIKE ? OR description ILIKE ? OR company ILIKE ?",
		searchQuery, searchQuery, searchQuery).
		Preload("Videos").
		Find(&jobs).Error
	return jobs, err
}

// GetJobByID retrieves a job by its ID with related videos
func (s *JobService) GetJobByID(id uint) (*Job, error) {
	var job Job
	err := s.db.Preload("Videos").First(&job, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("job with ID %d not found", id)
		}
		return nil, fmt.Errorf("failed to retrieve job: %w", err)
	}
	return &job, nil
}

// CreateJob creates a new job in the database
func (s *JobService) CreateJob(job *Job) error {
	job.PostedAt = time.Now()
	if err := s.db.Create(job).Error; err != nil {
		return fmt.Errorf("failed to create job: %w", err)
	}
	return nil
}

// UpdateJob updates an existing job
func (s *JobService) UpdateJob(id uint, job *Job) error {
	// Get existing job to preserve PostedAt
	var existingJob Job
	if err := s.db.First(&existingJob, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("job with ID %d not found", id)
		}
		return fmt.Errorf("failed to find job for update: %w", err)
	}

	// Preserve the original PostedAt and set the ID
	job.PostedAt = existingJob.PostedAt
	job.ID = id

	if err := s.db.Save(job).Error; err != nil {
		return fmt.Errorf("failed to update job: %w", err)
	}
	return nil
}

// DeleteJob soft deletes a job
func (s *JobService) DeleteJob(id uint) error {
	result := s.db.Delete(&Job{}, id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete job: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("job with ID %d not found", id)
	}
	return nil
}

// VideoService handles video-related database operations
type VideoService struct {
	db *gorm.DB
}

// NewVideoService creates a new VideoService
func NewVideoService(db *gorm.DB) *VideoService {
	return &VideoService{db: db}
}

// GetAllVideos retrieves all videos from the database
func (s *VideoService) GetAllVideos() ([]Video, error) {
	var videos []Video
	err := s.db.Preload("Job").Find(&videos).Error
	return videos, err
}

// GetVideosWithPagination retrieves videos with pagination
func (s *VideoService) GetVideosWithPagination(page, pageSize int) ([]Video, int64, error) {
	var videos []Video
	var total int64

	// Count total records
	if err := s.db.Model(&Video{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * pageSize

	// Get videos with pagination
	err := s.db.Preload("Job").
		Offset(offset).
		Limit(pageSize).
		Order("created_at DESC").
		Find(&videos).Error

	return videos, total, err
}

// GetVideoByID retrieves a video by its ID
func (s *VideoService) GetVideoByID(id uint) (*Video, error) {
	var video Video
	err := s.db.Preload("Job").First(&video, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("video with ID %d not found", id)
		}
		return nil, fmt.Errorf("failed to retrieve video: %w", err)
	}
	return &video, nil
}

// GetVideosByJobID retrieves all videos for a specific job
func (s *VideoService) GetVideosByJobID(jobID uint) ([]Video, error) {
	var videos []Video
	err := s.db.Where("job_id = ?", jobID).Find(&videos).Error
	return videos, err
}

// CreateVideo creates a new video in the database
func (s *VideoService) CreateVideo(video *Video) error {
	if err := s.db.Create(video).Error; err != nil {
		return fmt.Errorf("failed to create video: %w", err)
	}
	return nil
}

// UpdateVideo updates an existing video
func (s *VideoService) UpdateVideo(id uint, video *Video) error {
	video.ID = id
	if err := s.db.Save(video).Error; err != nil {
		return fmt.Errorf("failed to update video: %w", err)
	}
	return nil
}

// DeleteVideo soft deletes a video
func (s *VideoService) DeleteVideo(id uint) error {
	result := s.db.Delete(&Video{}, id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete video: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("video with ID %d not found", id)
	}
	return nil
}
