-- name: GetSaleItemsBySaleID :many
SELECT
    si.id          AS sale_item_id,
    si.product_id,
    si.product_amount,
    p.api_id       AS product_api_id,
    p.title        AS product_title,
    p.price        AS product_price,
    p.cover        AS product_cover,
    a.name         AS artist_name,
    a.id           AS artist_id,
    pt.type_name   AS product_type,
    pt.id          AS type_id,
    COALESCE(inv.stock, 0) AS stock
FROM sale_item si
JOIN products p ON si.product_id = p.id AND p.deleted_at IS NULL
JOIN artists a ON p.artist_id = a.id AND a.deleted_at IS NULL
JOIN product_types pt ON p.type_id = pt.id AND pt.deleted_at IS NULL
LEFT JOIN inventory inv ON p.id = inv.product_id AND inv.deleted_at IS NULL
WHERE si.sale_id = sqlc.arg(sale_id)
  AND si.deleted_at IS NULL;
