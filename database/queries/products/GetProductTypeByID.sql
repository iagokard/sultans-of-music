-- name: GetProductTypeByID :one
SELECT * FROM product_types
WHERE id = sqlc.arg(type_id)
  AND deleted_at IS NULL;
