// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (username,
                   email,
                   password,
                   name)
VALUES ($1, $2, $3, $4) RETURNING id, username, name, email, password, created_at
`

type CreateUserParams struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Username,
		arg.Email,
		arg.Password,
		arg.Name,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, username, name, email, password, created_at
FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET email    = $2,
    password = $3,
    name     = $4
WHERE username = $1
RETURNING id, username, name, email, password, created_at
`

type UpdateUserParams struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.Username,
		arg.Email,
		arg.Password,
		arg.Name,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
	)
	return i, err
}
