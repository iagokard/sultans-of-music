-- name: GetUserByID :one
SELECT * FROM users
WHERE id = sqlc.arg(user_id)
  AND deleted_at IS NULL;
