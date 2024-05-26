package db

import (
	"github.com/stretchr/testify/require"
	"sample_banking/db/utils"
	"testing"
)

func TestQueries_CreateUser(t *testing.T) {

	arg := CreateUserParams{
		Username: utils.RandomUserName(),
		Balance:  utils.RandomAmount(),
		Currency: utils.RandomCurrency(),
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
