package models

import "time"

type Task struct {
	ID          uint       `gorm:"primaryKey"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Priority    int        `json:"priority"`
	DueAt       *time.Time `json:"due_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
