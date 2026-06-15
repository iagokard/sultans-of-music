-- name: GetArtistByAPIID :one
SELECT * FROM artists
WHERE api_id = sqlc.arg(api_id)
  AND deleted_at IS NULL;
