-- name: CreateArtist :execresult
INSERT INTO artists (
    name,
    api_id
)
VALUES (
    sqlc.arg(name),
    sqlc.arg(api_id)
);
