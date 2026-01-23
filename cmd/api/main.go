package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/S-Axhwin/prac-02/internal/db/sqlc"
	"github.com/S-Axhwin/prac-02/internal/handlers"
	"github.com/S-Axhwin/prac-02/internal/middleware"
	"github.com/S-Axhwin/prac-02/internal/store"
)

func main() {

	db, err := store.NewPostgres(
		context.Background(),
		os.Getenv("DATABASE_URL"),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	queries := sqlc.New(db)
	h := handlers.NewHandler(queries)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", h.HealthCheck)

	mux.HandleFunc("POST /auth/register", h.Register)
	mux.HandleFunc("POST /auth/login", h.Login)

	//secure routes
	auth := func(hf http.HandlerFunc) http.Handler {
		return middleware.AuthMiddleware(hf)
	}

	mux.Handle("POST /auth/logout", auth(h.Logout))
	mux.Handle("GET /tasks", auth((h.GetTasks)))
	mux.Handle("POST /tasks", auth((h.CreateTasks)))

	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
