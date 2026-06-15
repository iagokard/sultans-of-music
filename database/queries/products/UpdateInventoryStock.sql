-- name: UpdateInventoryStock :execresult
UPDATE inventory
SET stock = stock - sqlc.arg(amount)
WHERE product_id = sqlc.arg(product_id)
  AND stock >= sqlc.arg(amount)
  AND deleted_at IS NULL;
