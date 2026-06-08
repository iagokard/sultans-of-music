-- name: GetInventoryByProductID :one
SELECT * FROM inventory
WHERE product_id = sqlc.arg(product_id)
  AND deleted_at IS NULL;
