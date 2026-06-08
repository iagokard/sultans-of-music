-- name: GetArtistByName :one
SELECT * FROM artists
WHERE name = sqlc.arg(artist_name)
  AND deleted_at IS NULL;
