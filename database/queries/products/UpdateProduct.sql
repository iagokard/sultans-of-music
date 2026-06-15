-- name: UpdateProduct :execresult
UPDATE products
SET
    api_id       = sqlc.arg(api_id),
    type_id      = sqlc.arg(type_id),
    release_date = sqlc.arg(release_date),
    title        = sqlc.arg(title),
    artist_id    = sqlc.arg(artist_id),
    price        = sqlc.arg(price),
    cover        = sqlc.arg(cover)
WHERE id = sqlc.arg(product_id);
