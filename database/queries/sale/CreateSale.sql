-- name: CreateSale :execresult
INSERT INTO sales (date_time)
VALUES (CURRENT_TIMESTAMP);
