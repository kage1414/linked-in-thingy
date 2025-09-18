package database

import (
	"errors"
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

// GetAllJobs retrieves all jobs from the database
func (s *JobService) GetAllJobs() ([]Job, error) {
	var jobs []Job
	err := s.db.Find(&jobs).Error
	return jobs, err
}

// GetJobByID retrieves a job by its ID
func (s *JobService) GetJobByID(id uint) (*Job, error) {
	var job Job
	err := s.db.First(&job, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("job not found")
		}
		return nil, err
	}
	return &job, nil
}

// CreateJob creates a new job in the database
func (s *JobService) CreateJob(job *Job) error {
	job.PostedAt = time.Now()
	return s.db.Create(job).Error
}

// UpdateJob updates an existing job
func (s *JobService) UpdateJob(id uint, job *Job) error {
	// Get existing job to preserve PostedAt
	var existingJob Job
	if err := s.db.First(&existingJob, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("job not found")
		}
		return err
	}

	// Preserve the original PostedAt
	job.PostedAt = existingJob.PostedAt
	job.ID = id

	return s.db.Save(job).Error
}

// DeleteJob soft deletes a job
func (s *JobService) DeleteJob(id uint) error {
	return s.db.Delete(&Job{}, id).Error
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

// GetVideoByID retrieves a video by its ID
func (s *VideoService) GetVideoByID(id uint) (*Video, error) {
	var video Video
	err := s.db.Preload("Job").First(&video, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("video not found")
		}
		return nil, err
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
	return s.db.Create(video).Error
}

// UpdateVideo updates an existing video
func (s *VideoService) UpdateVideo(id uint, video *Video) error {
	video.ID = id
	return s.db.Save(video).Error
}

// DeleteVideo soft deletes a video
func (s *VideoService) DeleteVideo(id uint) error {
	return s.db.Delete(&Video{}, id).Error
}
