package repository

import (
	"database/sql"

	"github.com/mink0ff/api_task_tracker/internal/models"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) CreateTask(task *models.Task) error {
	query := `INSERT INTO tasks (title, description, status, creator_id, assignee_id, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, NOW(), NOW()) RETURNING id created_at, updated_at`

	return r.db.QueryRow(query, task.Title, task.Description, task.Status, task.CreatorID, task.AssigneeID).
		Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt)
}

func (r *TaskRepository) GetTaskByID(id int) (*models.Task, error) {
	query := `SELECT id, title, description, status, creator_id, created_at, updated_at FROM tasks WHERE id = $1`
	task := &models.Task{}
	err := r.db.QueryRow(query, id).
		Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.AssigneeID, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (r *TaskRepository) UpdateTask(task *models.Task) error {
	query := `UPDATE tasks SET title = $1, description = $2, status = $3, assignee_id = $4 WHERE id = $5`
	_, err := r.db.Exec(query, task.Title, task.Description, task.Status, task.AssigneeID, task.ID)

	return err
}

func (r *TaskRepository) DeleteTask(id int) error {
	query := `DELETE FROM tasks WHERE id = $1`
	_, err := r.db.Exec(query, id)

	return err
}

func (r *TaskRepository) ListTasks() ([]*models.Task, error) {
	query := `SELECT id, title, description, status, creator_id, assignee_id, created_at, updated_atFROM tasks`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*models.Task
	for rows.Next() {
		task := &models.Task{}
		if err := rows.Scan(&task.ID, &task.Title, &task.Description,
			&task.Status, &task.CreatorID, &task.AssigneeID,
			&task.CreatedAt, &task.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
