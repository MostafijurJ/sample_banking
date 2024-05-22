

-- name: CreateEntry :one
INSERT INTO ENTRIES (account_id, amount)
values ($1, $2) RETURNING *;

-- name: GetEntries :many
SELECT * FROM ENTRIES WHERE account_id = $1;
