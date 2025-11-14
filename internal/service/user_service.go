package service

import (
	"errors"
	"time"

	"github.com/mink0ff/api_task_tracker/internal/models"
	"github.com/mink0ff/api_task_tracker/internal/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func hashPassword(password string) string {
	// не забыдь сделать реальную хеш-функцию
	return "hashed_" + password
}

func (s *UserService) CreateUser(req *models.CreateUserRequest) (*models.User, error) {
	user := &models.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: hashPassword(req.Password),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := s.userRepo.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
	return s.userRepo.GetUserByID(id)
}

func (s *UserService) UpdateUser(id int, req *models.UpdateUserRequest) error {
	user, err := s.userRepo.GetUserByID(id)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}
	if req.Username != nil {
		user.Username = *req.Username
	}
	if req.Email != nil {
		user.Email = *req.Email
	}
	if req.Password != nil {
		user.PasswordHash = hashPassword(*req.Password)
	}
	user.UpdatedAt = time.Now()

	return s.userRepo.UpdateUser(user)
}

func (s *UserService) AuthenticateUser(email, password string) (*models.User, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil || user.PasswordHash != hashPassword(password) {
		return nil, errors.New("invalid email or password")
	}
	return user, nil
}
