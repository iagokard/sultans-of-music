-- name: GetProductTypeByTypeName :one
SELECT * FROM product_types
WHERE type_name = sqlc.arg(type_name)
  AND deleted_at IS NULL;
