package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/S-Axhwin/prac-02/internal/db/sqlc"
	"github.com/S-Axhwin/prac-02/internal/middleware"
	"github.com/google/uuid"
)

// TODO: full handllers for task
//

func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	userIDRaw := r.Context().Value(middleware.UserIdKey)
	userIDStr, ok := userIDRaw.(string)
	if !ok {
		http.Error(w, "Invalid user id type", http.StatusUnauthorized)
		return
	}

	userUUID, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user id format", http.StatusUnauthorized)
		return
	}

	tasks, err := h.q.ListTasksByUser(r.Context(), userUUID)

	json.NewEncoder(w).Encode(tasks)

}

func (h *Handler) CreateTasks(w http.ResponseWriter, r *http.Request) {
	userIDRaw := r.Context().Value(middleware.UserIdKey)
	userIDStr, ok := userIDRaw.(string)
	if !ok {
		http.Error(w, "Invalid user id type", http.StatusUnauthorized)
		return
	}

	userUUID, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user id format", http.StatusUnauthorized)
		return
	}
	title := "idk"
	newTask, err := h.q.CreateTask(r.Context(), sqlc.CreateTaskParams{
		UserID: userUUID,
		Title:  title,
	})

	if err != nil {
		http.Error(w, "Error while creating task", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTask)
}
