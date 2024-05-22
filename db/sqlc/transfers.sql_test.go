package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"log"
	"sample_banking/db/utils"
	"testing"
)

func MakeTransaction(t *testing.T, fromAccount Account, toAccount Account) Transfer {
	// Create a transfer
	transfer := CreateTransferParams{
		FromAccountID:     fromAccount.ID,
		ToAccountID:       toAccount.ID,
		Amount:            utils.RandomAmount(),
		FromAccountNumber: utils.RandomAccountNumber(),
		ToAccountNumber:   utils.RandomAccountNumber(),
		Status:            "SUCCESS",
	}
	txn, err := testQueries.CreateTransfer(context.Background(), transfer)
	if err != nil {
		log.Fatal("Error creating transfer: ", err)
	}

	require.Equal(t, transfer.FromAccountID, txn.FromAccountID)
	require.NotZero(t, txn.Amount)
	return txn
}

func TestQueries_CreateTransfer(t *testing.T) {
	fromAccount := CreateNewAccount(t)
	toAccount := CreateNewAccount(t)

	transaction := MakeTransaction(t, fromAccount, toAccount)
	log.Fatal("Transaction Created : ", transaction)

}
