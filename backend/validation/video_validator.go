package validation

import (
	"job-board/backend/database"
)

// VideoValidator provides validation for Video entities
type VideoValidator struct {
	*Validator
}

// NewVideoValidator creates a new video validator
func NewVideoValidator() *VideoValidator {
	return &VideoValidator{
		Validator: NewValidator(),
	}
}

// ValidateVideo validates a video entity
func (vv *VideoValidator) ValidateVideo(video *database.Video) error {
	// Validate job ID
	if err := vv.ValidatePositiveInt(int(video.JobID), "jobId", true); err != nil {
		return err
	}

	// Validate title
	if err := vv.ValidateString(video.Title, "title", true, 200); err != nil {
		return err
	}

	// Validate URL
	if err := vv.ValidateURL(video.URL, "url", true); err != nil {
		return err
	}

	// Validate duration (optional)
	if video.Duration != nil {
		if err := vv.ValidatePositiveInt(*video.Duration, "duration", false); err != nil {
			return err
		}
	}

	// Validate thumbnail (optional)
	if video.Thumbnail != nil && *video.Thumbnail != "" {
		if err := vv.ValidateURL(*video.Thumbnail, "thumbnail", false); err != nil {
			return err
		}
	}

	return nil
}

// SanitizeVideo sanitizes a video entity
func (vv *VideoValidator) SanitizeVideo(video *database.Video) {
	video.Title = vv.SanitizeString(video.Title)
	video.URL = vv.SanitizeString(video.URL)

	if video.Thumbnail != nil {
		sanitized := vv.SanitizeString(*video.Thumbnail)
		video.Thumbnail = &sanitized
	}
}
