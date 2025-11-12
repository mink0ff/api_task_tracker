package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mink0ff/api_task_tracker/internal/config"
	"github.com/mink0ff/api_task_tracker/internal/database"
)

func main() {
	cfg := config.LoadConfig()

	database.RunMigrations(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := database.Connect(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", r)
}
