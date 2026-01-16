
-- name: CreateTask :one
INSERT INTO tasks (
  user_id,
  title
) VALUES (
  $1, $2
)
RETURNING id, user_id, title, completed, created_at;



-- name: ListTasksByUser :many
SELECT
  id,
  title,
  completed,
  created_at
FROM tasks
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: GetTaskByID :one
SELECT
  id,
  user_id,
  title,
  completed,
  created_at
FROM tasks
WHERE id = $1;


-- name: UpdateTaskStatus :one
UPDATE tasks
SET completed = $2
WHERE id = $1
RETURNING id, completed;


-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = $1;

