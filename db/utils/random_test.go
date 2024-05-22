package utils

// create a test case for RandomString
// Path: db/utils/random_test.go

import (
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestRandomAccountNumber(t *testing.T) {
	s := RandomAccountNumber()
	require.Len(t, s, 16)
	require.NotEmpty(t, s)
	log.Fatal("account number created successfully : ", s)
}
