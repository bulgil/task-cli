package storage

import (
	"time"
)

type Task struct {
	ID          int        `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type TaskStatus string

const (
	TaskStatusTODO       TaskStatus = "todo"
	TaskStatusDone       TaskStatus = "done"
	TaskStatusInProgress TaskStatus = "in_progress"
)
