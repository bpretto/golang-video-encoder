package domain

import "time"

type Job struct {
	ID               string
	OutputBucketPath string
	Status           string
	Video            *Video // This is a pointer to the video object
	Error            string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
