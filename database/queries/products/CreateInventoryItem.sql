-- name: CreateInventoryItem :execresult
INSERT INTO inventory (
    product_id,
    stock
) VALUES (
    sqlc.arg(product_id),
    sqlc.arg(stock)
);
