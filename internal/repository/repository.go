package repository

import (
	"database/sql"
)

type Repository struct {
	UserRepo *UserRepository
	TaskRepo *TaskRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		UserRepo: NewUserRepository(db),
		TaskRepo: NewTaskRepository(db),
	}
}
