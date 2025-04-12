package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitDB() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://user:password@localhost:5432/taskmanager"
	}

	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		log.Fatal("Failed to parse DB URL:", err)
	}

	DB, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Create tasks table if not exists
	_, err = DB.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS tasks (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT NOW(),
			completed BOOLEAN DEFAULT FALSE
		)
	`)
	if err != nil {
		log.Fatal("Failed to create tasks table:", err)
	}

	log.Println("Connected to PostgreSQL database")
}

func CloseDB() {
	DB.Close()
}
package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitDB() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://user:password@localhost:5432/taskmanager"
	}

	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		log.Fatal("Failed to parse DB URL:", err)
	}

	DB, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Create tasks table if not exists
	_, err = DB.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS tasks (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT NOW(),
			completed BOOLEAN DEFAULT FALSE
		)
	`)
	if err != nil {
		log.Fatal("Failed to create tasks table:", err)
	}

	log.Println("Connected to PostgreSQL database")
}

func CloseDB() {
	DB.Close()
}