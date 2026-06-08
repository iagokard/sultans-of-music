-- name: GetProductsListPage :many
SELECT
    p.id,
    p.api_id,
    p.type_id,
    pt.type_name,
    p.release_date,
    p.title,
    p.artist_id,
    a.name AS artist_name,
    p.price,
    p.cover,
    COALESCE(i.stock, 0) AS stock
FROM products p
JOIN artists a ON p.artist_id = a.id
JOIN product_types pt ON p.type_id = pt.id
LEFT JOIN inventory i ON p.id = i.product_id AND i.deleted_at IS NULL
WHERE p.deleted_at IS NULL
  AND a.deleted_at IS NULL
  AND pt.deleted_at IS NULL
ORDER BY a.name
LIMIT ?
OFFSET ?;
