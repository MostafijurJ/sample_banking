-- name: CreateUser :one
INSERT INTO users (username,
                   email,
                   password,
                   name)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetUserByUsername :one
SELECT *
FROM users
WHERE username = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET email    = $2,
    password = $3,
    name     = $4
WHERE username = $1
RETURNING *;