# Go Http Server
This is not just a http application but build for learning go language folder and how go creator(Ken Thompson) want the go to be used.

# Why this project?

I built this project to deeply understand how a production-style backend works
without hiding complexity behind frameworks.

The goal was to learn:
- How HTTP servers work at the net/http level
- How SQL is actually executed (no ORM magic)
- How authentication and middleware really work
- How data flows from request → DB → response


## Tech Stack & Decisions

- **Go (net/http)**  
  Used directly instead of frameworks to understand routing, middleware chaining, and request lifecycle.

- **PostgreSQL**  
  Chosen for strong relational and performance. 

- **pgx + pgxpool**  
  Used instead of database/sql for better control.

- **SQLC**  
  Used to generate type-safe queries while still writing raw SQL.
  This avoids ORMs while preventing runtime SQL errors.

- **JWT (cookie-based auth)**  
  Implemented manually to understand signing, verification, and middleware-based auth.


# Folder Structure
```
.
├── cmd
│   └── api
│       └── main.go
├── go.mod
├── go.sum
├── internal
│   ├── db
│   │   ├── migrations
│   │   │   ├── 001_init.sql
│   │   │   └── 002_tasks.sql
│   │   ├── queries
│   │   │   ├── auth.sql
│   │   │   └── task.sql
│   │   └── sqlc
│   │       ├── auth.sql.go
│   │       ├── db.go
│   │       ├── models.go
│   │       ├── querier.go
│   │       └── task.sql.go
│   ├── handlers
│   │   ├── auth.go
│   │   └── tasks.go
│   ├── middleware
│   │   └── userMiddle.go
│   └── store
│       └── postgres.go
├── Makefile
├── README.md
├── sqlc.yaml
├── tesh.sh
└── test.http

```


