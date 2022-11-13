package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateNewAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    "TOM",
		Balance:  1001,
		Currency: "BDT",
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	CreateNewAccount(t)
}

func TestGetAccount(t *testing.T) {
	CreateNewAccount(t)

}
