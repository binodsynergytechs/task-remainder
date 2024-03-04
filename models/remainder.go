package models

import "time"

type Reminder struct {
	ID         uint       `gorm:"primaryKey"`
	TaskID     uint       `json:"task_id"`
	Contact    string     `json:"contact"`
	ReminderAt *time.Time `json:"reminder_at"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}
