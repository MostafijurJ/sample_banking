package db

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func CreateNewEntry(t *testing.T, account Account) Entry {
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    account.Balance,
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	fmt.Println("Entry Created : ", entry)
	return entry
}

func TestCreateEntry(t *testing.T) {
	account := CreateNewAccount(t)
	CreateNewEntry(t, account)
}

func TestGetEntry(t *testing.T) {
	account := CreateNewAccount(t)
	entry := CreateNewEntry(t, account)

	entries, err := testQueries.GetEntries(context.Background(), account.ID)

	if err != nil {
		log.Fatal("Error in GetEntryTest")
	}
	require.NoError(t, err)
	require.NotEmpty(t, entries)

	require.Len(t, entries, 1)

	e := entries[0]
	require.Equal(t, entry.ID, e.ID)
	require.Equal(t, entry.AccountID, e.AccountID)
	require.Equal(t, entry.Amount, e.Amount)
	require.Equal(t, entry.CreatedAt, e.CreatedAt)

}
