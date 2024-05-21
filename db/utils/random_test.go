package utils

// create a test case for RandomString
// Path: db/utils/random_test.go

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRandomString(t *testing.T) {
	s := RandomString(12)
	fmt.Println("Random String OutPut ->  " + s)
	require.Len(t, s, 12)
	require.NotEmpty(t, s)
}
