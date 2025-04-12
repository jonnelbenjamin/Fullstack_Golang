package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jonnelbenjamin/Fullstack_Golang/backend/models"
	"github.com/jonnelbenjamin/Fullstack_Golang/backend/templates"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	tasks, err := models.GetAllTasks(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	templates.Render(w, "index.html", tasks)
}

func HandleGetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := models.GetAllTasks(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Header.Get("HX-Request") == "true" {
		templates.Render(w, "task-list.html", tasks)
	} else {
		json.NewEncoder(w).Encode(tasks)
	}
}

func HandleCreateTask(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	if title == "" {
		var payload struct {
			Title string `json:"title"`
		}
		if err := json.NewDecoder(r.Body).Decode(&payload); err == nil {
			title = payload.Title
		}
	}

	task, err := models.CreateTask(r.Context(), title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Trigger", "taskUpdate")
	if r.Header.Get("HX-Request") == "true" {
		templates.Render(w, "task.html", task)
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(task)
	}
}

func HandleUpdateTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var payload struct {
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := models.UpdateTask(r.Context(), id, payload.Title, payload.Completed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Trigger", "taskUpdate")
	templates.Render(w, "task.html", task)
}

func HandleDeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	if err := models.DeleteTask(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Trigger", "taskUpdate")
	w.WriteHeader(http.StatusNoContent)
}
