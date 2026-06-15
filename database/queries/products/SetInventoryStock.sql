-- name: SetInventoryStock :execresult
UPDATE inventory
SET stock = sqlc.arg(new_stock)
WHERE product_id = sqlc.arg(product_id)
  AND deleted_at IS NULL;
