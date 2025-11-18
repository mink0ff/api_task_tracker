package main

import (
	"log"
	"net/http"
	_ "os/user"

	"github.com/go-chi/chi/v5"

	"github.com/mink0ff/api_task_tracker/internal/config"
	"github.com/mink0ff/api_task_tracker/internal/database"
	"github.com/mink0ff/api_task_tracker/internal/handler"
	"github.com/mink0ff/api_task_tracker/internal/repository"
	"github.com/mink0ff/api_task_tracker/internal/service"
)

func main() {
	cfg := config.LoadConfig()

	database.RunMigrations(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := database.Connect(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	taskRepo := repository.NewTaskRepository(db.DB)
	userRepo := repository.NewUserRepository(db.DB)

	taskService := service.NewTaskService(taskRepo)
	userService := service.NewUserService(userRepo)

	taskHandler := handler.NewTaskHandler(taskService)
	userHandler := handler.NewUserHandler(userService)

	r := chi.NewRouter()

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	r.Route("/tasks", func(r chi.Router) {
		r.Post("/", taskHandler.CreateTask)
		r.Get("/", taskHandler.GetTasksByAssigneeID)
		r.Put("/{id}", taskHandler.UpdateTask)
		r.Delete("/{id}", taskHandler.DeleteTask)
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.CreateUser)
		r.Get("/{id}", userHandler.GetUserByID)
	})

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", r)
}
