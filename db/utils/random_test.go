package utils

// create a test case for RandomString
// Path: db/helper/random_test.go

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRandomAccountNumber(t *testing.T) {
	s := RandomAccountNumber()
	require.Len(t, s, 16)
	require.NotEmpty(t, s)
}
