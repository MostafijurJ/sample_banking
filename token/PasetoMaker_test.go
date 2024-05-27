package token

import (
	"github.com/stretchr/testify/require"
	"sample_banking/db/utils"
	"testing"
	"time"
)

func TestPasetoMaker_CreateToken(t *testing.T) {

	maker, err := NewPasetoMaker(utils.RandomString(32))
	if err != nil {
		t.Error(err)
	}
	userName := utils.RandomUserName()
	duration := time.Minute

	issueAt := time.Now()
	expiredAt := issueAt.Add(duration)

	token, err := maker.CreateToken(userName, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	verifyToken, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, verifyToken)

	require.Equal(t, userName, verifyToken.Username)
	require.WithinDuration(t, issueAt, verifyToken.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, verifyToken.ExpiredAt, time.Second)
}

func TestPasetoMaker_VerifyToken(t *testing.T) {
	maker, err := NewPasetoMaker(utils.RandomString(32))
	require.NoError(t, err)

	token, err := maker.CreateToken(utils.RandomUserName(), time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
}
