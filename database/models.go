package database

import (
	"time"

	"gorm.io/gorm"
)

// Job represents a job posting in the database
type Job struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Title        string    `json:"title" gorm:"not null"`
	Company      string    `json:"company" gorm:"not null"`
	Description  string    `json:"description" gorm:"type:text"`
	Location     string    `json:"location" gorm:"not null"`
	Salary       *string   `json:"salary"`
	Requirements []string  `json:"requirements" gorm:"type:text[]"`
	Benefits     []string  `json:"benefits" gorm:"type:text[]"`
	PostedAt     time.Time `json:"postedAt" gorm:"default:CURRENT_TIMESTAMP"`
	VideoURL     *string   `json:"videoUrl"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

// Video represents a video associated with a job
type Video struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	JobID     uint      `json:"jobId" gorm:"not null"`
	Title     string    `json:"title" gorm:"not null"`
	URL       string    `json:"url" gorm:"not null"`
	Duration  *int      `json:"duration"`
	Thumbnail *string   `json:"thumbnail"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	
	// Relationship
	Job Job `json:"job" gorm:"foreignKey:JobID"`
}

// TableName specifies the table name for Job
func (Job) TableName() string {
	return "jobs"
}

// TableName specifies the table name for Video
func (Video) TableName() string {
	return "videos"
}
