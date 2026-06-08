-- name: CreateProduct :execresult
INSERT INTO products (
    api_id,
    type_id,
    release_date,
    title,
    artist_id,
    price,
    cover
)
VALUES (
    sqlc.arg(api_id),
    sqlc.arg(type_id),
    sqlc.arg(release_date),
    sqlc.arg(title),
    sqlc.arg(artist_id),
    sqlc.arg(price),
    sqlc.arg(cover)
);
