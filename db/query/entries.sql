

-- name: CreateEntry :one
INSERT INTO ENTRIES (account_id, amount)
values ($1, $2) RETURNING *;

-- name: GetEntries :many
SELECT * FROM ENTRIES WHERE account_id = $1;

-- name: GetEntryById :one
SELECT * FROM ENTRIES WHERE id = $1;
