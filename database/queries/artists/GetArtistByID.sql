-- name: GetArtistByID :one
SELECT * FROM artists
WHERE id = sqlc.arg(artist_id)
  AND deleted_at IS NULL;
