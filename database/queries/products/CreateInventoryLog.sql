-- name: CreateInventoryLog :execresult
INSERT INTO inventory_log (product_id, amount, type_id)
VALUES (
    sqlc.arg(product_id),
    sqlc.arg(amount),
    (SELECT id FROM inventory_log_type WHERE type_name = sqlc.arg(type_name) AND deleted_at IS NULL)
);
