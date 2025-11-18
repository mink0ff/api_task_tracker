package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mink0ff/api_task_tracker/internal/models"
	"github.com/mink0ff/api_task_tracker/internal/service"
	"github.com/mink0ff/api_task_tracker/internal/utils"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

// POST /users
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req models.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	user, err := h.service.CreateUser(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, user, http.StatusCreated)
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	userIDParam := chi.URLParam(r, "id")
	if userIDParam == "" {
		http.Error(w, "missing user ID parameter", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		http.Error(w, "invalid user ID parameter", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUserByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, user, http.StatusOK)
}

// func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
// 	var req models.UpdateUserRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		http.Error(w, "invalid request", http.StatusBadRequest)
// 		return
// 	}

// 	user, err := h.service.UpdateUser(&req)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	utils.WriteJSON(w, user, http.StatusOK)
// }

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Implementation goes here
}
