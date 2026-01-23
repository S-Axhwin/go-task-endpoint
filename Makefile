
build:
	go build -o bin/api ./cmd/api

run:
	go run ./cmd/api

test:
	go test ./...

sqlc:
	sqlc generate

migrate:
	psql $(DB_URL) -f internal/db/migrations/001_init.sql
