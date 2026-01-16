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
	mux.HandleFunc("POST /auth/register", h.Register)
	mux.HandleFunc("GET /health", healthCheck)
	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Server Running")
}
