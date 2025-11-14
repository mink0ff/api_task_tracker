package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/mink0ff/api_task_tracker/internal/models"
	"github.com/mink0ff/api_task_tracker/internal/service"
	"github.com/mink0ff/api_task_tracker/internal/utils"
)

type TaskHandler struct {
	taskService *service.TaskService
}

func NewTaskHandler(taskService *service.TaskService) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req models.CreateTaskRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	userIDHeader := r.Header.Get("X-User-ID")
	if userIDHeader == "" {
		http.Error(w, "missing X-User-ID header", http.StatusBadRequest)
		return
	}

	creatorID, err := strconv.Atoi(userIDHeader)
	if err != nil {
		http.Error(w, "invalid X-User-ID header", http.StatusBadRequest)
		return
	}

	task, err := h.taskService.CreateTask(&req, creatorID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, task, http.StatusCreated)
}
