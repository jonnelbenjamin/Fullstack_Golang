package models

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	Completed bool      `json:"completed"`
}

func GetAllTasks(ctx context.Context) ([]Task, error) {
	rows, err := db.DB.Query(ctx, "SELECT id, title, created_at, completed FROM tasks ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[Task])
}

func CreateTask(ctx context.Context, title string) (Task, error) {
	var task Task
	err := db.DB.QueryRow(ctx,
		"INSERT INTO tasks (title) VALUES ($1) RETURNING id, title, created_at, completed",
		title,
	).Scan(&task.ID, &task.Title, &task.CreatedAt, &task.Completed)
	return task, err
}

func UpdateTask(ctx context.Context, id int, title string, completed bool) (Task, error) {
	var task Task
	err := db.DB.QueryRow(ctx,
		"UPDATE tasks SET title = $1, completed = $2 WHERE id = $3 RETURNING id, title, created_at, completed",
		title, completed, id,
	).Scan(&task.ID, &task.Title, &task.CreatedAt, &task.Completed)
	return task, err
}

func DeleteTask(ctx context.Context, id int) error {
	_, err := db.DB.Exec(ctx, "DELETE FROM tasks WHERE id = $1", id)
	return err
}
