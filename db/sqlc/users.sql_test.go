package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"sample_banking/db/utils"
	"testing"
)

func CreateTestUser(t *testing.T) User {
	arg := CreateUserParams{
		Username: utils.RandomUserName(),
		Email:    utils.RandomEmail(),
		Password: utils.RandomString(10),
		Name:     utils.RandomString(22),
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	if err != nil {
		t.Errorf("Error creating user: %v", err)
	}
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Password, user.Password)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	CreateTestUser(t)
}

func TestQueries_GetUserByUsername(t *testing.T) {
	user1 := CreateTestUser(t)
	user2, err := testQueries.GetUserByUsername(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Password, user2.Password)
}

func TestQueries_UpdateUser(t *testing.T) {
	createdUser := CreateTestUser(t)

	arg := UpdateUserParams{
		Username: createdUser.Username,
		Email:    utils.RandomEmail(),
		Password: utils.RandomString(3),
		Name:     utils.RandomString(11),
	}

	updatedUser, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedUser)

	require.Equal(t, arg.Username, updatedUser.Username)
	require.Equal(t, arg.Email, updatedUser.Email)
	require.Equal(t, arg.Password, updatedUser.Password)
	require.Equal(t, arg.Name, updatedUser.Name)
	require.Equal(t, createdUser.ID, updatedUser.ID)
}
