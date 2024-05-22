
-- name: CreateTransfer :one
INSERT INTO TRANSFERS (from_account_id, to_account_id, amount, from_account_number, to_account_number, status)
VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;


-- name: GetTransferById :one
SELECT * FROM TRANSFERS WHERE id = $1;
