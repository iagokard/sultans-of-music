-- name: GetSaleByID :one
SELECT * FROM sales
WHERE id = sqlc.arg(sale_id)
  AND deleted_at IS NULL;
