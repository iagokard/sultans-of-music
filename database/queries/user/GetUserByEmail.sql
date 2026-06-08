-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = sqlc.arg(email)
  AND deleted_at IS NULL;
