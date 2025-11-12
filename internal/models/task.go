package models

import (
	"time"
)

type TaskStatus string

const (
	StatusInactive   TaskStatus = "inactive"
	StatusInProgress TaskStatus = "in_progress"
	StatusDone       TaskStatus = "Done"
)

type Task struct {
	ID          int64      `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatorID   int64      `json:"creator_id"`
	AssigneeID  int64      `json:"assignee_id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type CreateTaskRequest struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	AssigneeID  int64      `json:"assignee_id"`
}

type UpdateTaskRequest struct {
	Title       *string     `json:"title,omitempty"`
	Description *string     `json:"description,omitempty"`
	Status      *TaskStatus `json:"status,omitempty"`
	AssigneeID  *int64      `json:"assignee_id,omitempty"`
}
