package handler

import (
	"encoding/json"
	"net/http"

	"github.com/mink0ff/tasktracker/internal/models"
	"github.com/mink0ff/tasktracker/internal/service"
	"github.com/mink0ff/tasktracker/internal/utils"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteJSON(w, map[string]string{"error": "invalid request body"}, http.StatusBadRequest)
		return
	}

	token, err := h.authService.Login(&req)

	if err != nil {
		if err == service.ErrInvalidCredentials {
			utils.WriteJSON(w, map[string]string{"error": "invalid email or password"}, http.StatusUnauthorized)
		} else {
			utils.WriteJSON(w, map[string]string{"error": "internal server error"}, http.StatusInternalServerError)
		}
		return
	}

	utils.WriteJSON(w, map[string]string{"access_token": token}, http.StatusOK)
}
