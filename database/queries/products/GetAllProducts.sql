-- name: GetAllProducts :many
SELECT
    p.*,
    COALESCE(i.stock, 0) AS stock
FROM products p
LEFT JOIN inventory i ON p.id = i.product_id AND i.deleted_at IS NULL
WHERE p.deleted_at IS NULL
ORDER BY p.id;
