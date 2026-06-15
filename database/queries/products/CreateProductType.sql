-- name: CreateProductType :execresult
INSERT INTO product_types (type_name)
VALUES (sqlc.arg(type_name));
