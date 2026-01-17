package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/S-Axhwin/prac-02/internal/db/sqlc"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	q *sqlc.Queries
}

func NewHandler(q *sqlc.Queries) *Handler {
	return &Handler{q: q}
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Server Running")
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invaild Input", http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}

	user, err := h.q.CreateUser(ctx, sqlc.CreateUserParams{
		Email:        req.Email,
		PasswordHash: string(hash),
	})

	if err != nil {
		http.Error(w, "Email already exisits", http.StatusConflict)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	//TODO: Impliment Login
	ctx := r.Context()

	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w, "Invalid Input", http.StatusBadRequest)
		return
	}

	// get the user first
	user, err := h.q.GetUserByEmail(ctx, req.Email)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w, "User didn exisits with mail", http.StatusBadRequest)
		return
	}
	bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))

}
