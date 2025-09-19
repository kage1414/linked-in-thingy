package model

type Job struct {
	ID           string   `json:"id"`
	Title        string   `json:"title"`
	Company      string   `json:"company"`
	Description  string   `json:"description"`
	Location     string   `json:"location"`
	Salary       *string  `json:"salary"`
	Requirements []string `json:"requirements"`
	Benefits     []string `json:"benefits"`
	PostedAt     string   `json:"postedAt"`
	VideoURL     *string  `json:"videoUrl"`
}

type Video struct {
	ID        string  `json:"id"`
	JobID     string  `json:"jobId"`
	Title     string  `json:"title"`
	URL       string  `json:"url"`
	Duration  *int    `json:"duration"`
	Thumbnail *string `json:"thumbnail"`
}

type JobInput struct {
	Title        string   `json:"title"`
	Company      string   `json:"company"`
	Description  string   `json:"description"`
	Location     string   `json:"location"`
	Salary       *string  `json:"salary"`
	Requirements []string `json:"requirements"`
	Benefits     []string `json:"benefits"`
	VideoURL     *string  `json:"videoUrl"`
}

type VideoInput struct {
	JobID     string  `json:"jobId"`
	Title     string  `json:"title"`
	URL       string  `json:"url"`
	Duration  *int    `json:"duration"`
	Thumbnail *string `json:"thumbnail"`
}
