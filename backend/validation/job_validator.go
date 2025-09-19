package validation

import (
	"job-board/backend/database"
)

// JobValidator provides validation for Job entities
type JobValidator struct {
	*Validator
}

// NewJobValidator creates a new job validator
func NewJobValidator() *JobValidator {
	return &JobValidator{
		Validator: NewValidator(),
	}
}

// ValidateJob validates a job entity
func (jv *JobValidator) ValidateJob(job *database.Job) error {
	// Validate title
	if err := jv.ValidateString(job.Title, "title", true, 200); err != nil {
		return err
	}

	// Validate company
	if err := jv.ValidateString(job.Company, "company", true, 100); err != nil {
		return err
	}

	// Validate description
	if err := jv.ValidateString(job.Description, "description", true, 2000); err != nil {
		return err
	}

	// Validate location
	if err := jv.ValidateString(job.Location, "location", true, 100); err != nil {
		return err
	}

	// Validate salary (optional)
	if job.Salary != nil && *job.Salary != "" {
		if err := jv.ValidateString(*job.Salary, "salary", false, 50); err != nil {
			return err
		}
	}

	// Validate requirements
	if err := jv.ValidateStringSlice(job.Requirements, "requirements", true, 20); err != nil {
		return err
	}

	// Validate benefits
	if err := jv.ValidateStringSlice(job.Benefits, "benefits", true, 20); err != nil {
		return err
	}

	// Validate video URL (optional)
	if job.VideoURL != nil && *job.VideoURL != "" {
		if err := jv.ValidateURL(*job.VideoURL, "videoUrl", false); err != nil {
			return err
		}
	}

	return nil
}

// SanitizeJob sanitizes a job entity
func (jv *JobValidator) SanitizeJob(job *database.Job) {
	job.Title = jv.SanitizeString(job.Title)
	job.Company = jv.SanitizeString(job.Company)
	job.Description = jv.SanitizeHTML(job.Description)
	job.Location = jv.SanitizeString(job.Location)

	if job.Salary != nil {
		sanitized := jv.SanitizeString(*job.Salary)
		job.Salary = &sanitized
	}

	// Sanitize requirements
	for i, req := range job.Requirements {
		job.Requirements[i] = jv.SanitizeString(req)
	}

	// Sanitize benefits
	for i, benefit := range job.Benefits {
		job.Benefits[i] = jv.SanitizeString(benefit)
	}

	if job.VideoURL != nil {
		sanitized := jv.SanitizeString(*job.VideoURL)
		job.VideoURL = &sanitized
	}
}
