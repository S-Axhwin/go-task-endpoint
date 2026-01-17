package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/S-Axhwin/prac-02/internal/db/sqlc"
	"github.com/S-Axhwin/prac-02/internal/handlers"
	"github.com/S-Axhwin/prac-02/internal/store"
)

func main() {

	db, err := store.NewPostgres(context.Background(), os.Getenv("DATABASE_URL"))
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
	mux.HandleFunc("POST /auth/logout", h.Logout)

	// Secure Routes

	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
