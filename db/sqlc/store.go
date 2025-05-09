package db

import (
	"context"
)

type Store interface {
	GetUserByUsername(ctx context.Context, username string) (User, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}
