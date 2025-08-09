package models

import "time"

type Task struct {
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	Tags        []string   `json:"tags,omitempty"`
}
