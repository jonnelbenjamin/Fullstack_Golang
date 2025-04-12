package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/jonnelbenjamin/Fullstack_Golang/backend/db"
	"github.com/jonnelbenjamin/Fullstack_Golang/backend/handlers"
	"github.com/jonnelbenjamin/Fullstack_Golang/backend/templates"
)

func main() {
	// Load .env file if it exists (for local development)
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file:", err)
		}
	}

	// Initialize database
	db.InitDB()
	defer db.CloseDB()

	// Initialize templates
	if err := templates.InitTemplates(); err != nil {
		log.Fatal("Failed to initialize templates:", err)
	}

	// Setup router
	r := chi.NewRouter()

	// Static files
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("../frontend/static"))))

	// Routes
	r.Get("/", handlers.HandleIndex)
	r.Get("/tasks", handlers.HandleGetTasks)
	r.Post("/tasks", handlers.HandleCreateTask)
	r.Put("/tasks/{id}", handlers.HandleUpdateTask)
	r.Delete("/tasks/{id}", handlers.HandleDeleteTask)

	// Start server
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("Server running on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server error:", err)
		}
	}()

	<-quit
	log.Println("Shutting down server...")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatal("Server shutdown error:", err)
	}
	log.Println("Server stopped")
}
