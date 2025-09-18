package graph

import (
	"context"
	"fmt"
	"time"

	"job-board/graph/model"
)

// This file will not be regenerated automatically.
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

// Jobs returns all jobs
func (r *Resolver) Jobs(ctx context.Context) ([]*model.Job, error) {
	// Mock data - in a real app, this would come from a database
	jobs := []*model.Job{
		{
			ID:          "1",
			Title:       "Senior Software Engineer",
			Company:     "Tech Corp",
			Description: "We are looking for a senior software engineer to join our team...",
			Location:    "San Francisco, CA",
			Salary:      "$120,000 - $150,000",
			Requirements: []string{
				"5+ years of experience",
				"Proficiency in Go and React",
				"Experience with GraphQL",
			},
			Benefits: []string{
				"Health insurance",
				"401k matching",
				"Flexible work hours",
			},
			PostedAt: time.Now().Add(-24 * time.Hour).Format(time.RFC3339),
			VideoURL: "/video/1",
		},
		{
			ID:          "2",
			Title:       "Frontend Developer",
			Company:     "StartupXYZ",
			Description: "Join our fast-growing startup as a frontend developer...",
			Location:    "Remote",
			Salary:      "$80,000 - $100,000",
			Requirements: []string{
				"3+ years of React experience",
				"TypeScript proficiency",
				"Experience with modern CSS",
			},
			Benefits: []string{
				"Stock options",
				"Unlimited PTO",
				"Learning budget",
			},
			PostedAt: time.Now().Add(-48 * time.Hour).Format(time.RFC3339),
			VideoURL: "/video/2",
		},
	}
	return jobs, nil
}

// Job returns a specific job by ID
func (r *Resolver) Job(ctx context.Context, id string) (*model.Job, error) {
	jobs, err := r.Jobs(ctx)
	if err != nil {
		return nil, err
	}

	for _, job := range jobs {
		if job.ID == id {
			return job, nil
		}
	}
	return nil, fmt.Errorf("job not found")
}

// Videos returns all videos
func (r *Resolver) Videos(ctx context.Context) ([]*model.Video, error) {
	videos := []*model.Video{
		{
			ID:        "1",
			JobID:     "1",
			Title:     "Company Culture Video",
			URL:       "/video/1",
			Duration:  120,
			Thumbnail: "/thumbnails/1.jpg",
		},
		{
			ID:        "2",
			JobID:     "2",
			Title:     "Team Introduction",
			URL:       "/video/2",
			Duration:  90,
			Thumbnail: "/thumbnails/2.jpg",
		},
	}
	return videos, nil
}

// Video returns a specific video by ID
func (r *Resolver) Video(ctx context.Context, id string) (*model.Video, error) {
	videos, err := r.Videos(ctx)
	if err != nil {
		return nil, err
	}

	for _, video := range videos {
		if video.ID == id {
			return video, nil
		}
	}
	return nil, fmt.Errorf("video not found")
}

// CreateJob creates a new job
func (r *Resolver) CreateJob(ctx context.Context, input model.JobInput) (*model.Job, error) {
	// In a real app, this would save to database
	job := &model.Job{
		ID:           fmt.Sprintf("%d", time.Now().Unix()),
		Title:        input.Title,
		Company:      input.Company,
		Description:  input.Description,
		Location:     input.Location,
		Salary:       input.Salary,
		Requirements: input.Requirements,
		Benefits:     input.Benefits,
		PostedAt:     time.Now().Format(time.RFC3339),
		VideoURL:     input.VideoURL,
	}
	return job, nil
}

// UpdateJob updates an existing job
func (r *Resolver) UpdateJob(ctx context.Context, id string, input model.JobInput) (*model.Job, error) {
	// In a real app, this would update in database
	job := &model.Job{
		ID:           id,
		Title:        input.Title,
		Company:      input.Company,
		Description:  input.Description,
		Location:     input.Location,
		Salary:       input.Salary,
		Requirements: input.Requirements,
		Benefits:     input.Benefits,
		PostedAt:     time.Now().Format(time.RFC3339),
		VideoURL:     input.VideoURL,
	}
	return job, nil
}

// DeleteJob deletes a job
func (r *Resolver) DeleteJob(ctx context.Context, id string) (bool, error) {
	// In a real app, this would delete from database
	return true, nil
}

// CreateVideo creates a new video
func (r *Resolver) CreateVideo(ctx context.Context, input model.VideoInput) (*model.Video, error) {
	// In a real app, this would save to database
	video := &model.Video{
		ID:        fmt.Sprintf("%d", time.Now().Unix()),
		JobID:     input.JobID,
		Title:     input.Title,
		URL:       input.URL,
		Duration:  input.Duration,
		Thumbnail: input.Thumbnail,
	}
	return video, nil
}
