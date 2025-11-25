package service

import (
	"errors"

	"github.com/mink0ff/api_task_tracker/internal/auth"
	"github.com/mink0ff/api_task_tracker/internal/models"
	"github.com/mink0ff/api_task_tracker/internal/repository"
	"github.com/mink0ff/api_task_tracker/internal/utils"
)

type AuthService struct {
	userRepo *repository.UserRepository
	jwtMaker *auth.JWTManager
}

func NewAuthService(userRepo *repository.UserRepository, jwtMaker *auth.JWTManager) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		jwtMaker: jwtMaker,
	}
}

var ErrInvalidCredentials = errors.New("invalid email or password")

func (s *AuthService) Login(req *models.LoginRequest) (string, error) {
	user, err := s.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		return "", err
	}

	if user == nil || !utils.CheckPassword(user.PasswordHash, req.Password) {
		return "", ErrInvalidCredentials
	}

	if !utils.CheckPassword(user.PasswordHash, req.Password) {
		return "", ErrInvalidCredentials
	}

	return s.jwtMaker.Generate(user.ID)
}
