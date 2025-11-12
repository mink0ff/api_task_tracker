package service

import (
	"errors"
	"time"

	"github.com/mink0ff/api_task_tracker/internal/models"
	"github.com/mink0ff/api_task_tracker/internal/repository"
)

type TaskService struct {
	taskRepo *repository.TaskRepository
}

func NewTaskService(taskRepo *repository.TaskRepository) *TaskService {
	return &TaskService{taskRepo: taskRepo}
}

func (s *TaskService) CreateTask(req *models.CreateTaskRequest, creatorID int) (*models.Task, error) {
	task := &models.Task{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		CreatorID:   creatorID,
		AssigneeID:  req.AssigneeID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.taskRepo.CreateTask(task); err != nil {
		return nil, err
	}

	return task, nil
}

func (s *TaskService) GetTaskByID(id int) (*models.Task, error) {
	return s.taskRepo.GetTaskByID(id)
}

func (s *TaskService) GetTasksByAssigneeID(assignee_id int) ([]*models.Task, error) {
	return s.taskRepo.GetTasksByAssigneeID(assignee_id)
}

func (s *TaskService) UpdateTask(id int, req models.UpdateTaskRequest) error {
	task, err := s.taskRepo.GetTaskByID(id)
	if err != nil {
		return err
	}
	if task == nil {
		return errors.New("task not found")
	}
	if req.Title != nil {
		task.Title = *req.Title
	}
	if req.Description != nil {
		task.Description = *req.Description
	}
	if req.Status != nil {
		task.Status = *req.Status
	}
	if req.AssigneeID != nil {
		task.AssigneeID = *req.AssigneeID
	}

	task.UpdatedAt = time.Now()

	return s.taskRepo.UpdateTask(task)
}

func (s *TaskService) DeleteTask(id int) error {
	return s.taskRepo.DeleteTask(id)
}
