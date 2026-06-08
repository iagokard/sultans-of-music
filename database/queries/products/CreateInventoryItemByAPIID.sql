-- name: CreateInventoryItemByAPIID :execresult
INSERT INTO inventory (product_id, stock)
SELECT id, sqlc.arg(stock)
FROM products
WHERE api_id = sqlc.arg(api_id)
  AND deleted_at IS NULL;
