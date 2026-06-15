-- name: GetAllSales :many
SELECT * FROM sales
WHERE deleted_at IS NULL
ORDER BY date_time DESC;
