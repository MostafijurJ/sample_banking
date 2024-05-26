package utils

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "password"
	hash, err := HashPassword(password)
	if err != nil {
		t.Error(err)
	}
	if hash == "" {
		t.Error("hash is empty")
	}

	require.NoError(t, err)
	require.NotEmpty(t, hash)
	err = CheckPassword(password, hash)
	require.NoError(t, err)
}
