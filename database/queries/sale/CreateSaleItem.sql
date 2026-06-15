-- name: CreateSaleItem :execresult
INSERT INTO sale_item (sale_id, product_id, product_amount)
VALUES (
    sqlc.arg(sale_id),
    sqlc.arg(product_id),
    sqlc.arg(product_amount)
);
