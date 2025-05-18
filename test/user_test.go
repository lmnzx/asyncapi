package test

import (
	"context"
	"testing"
	"time"

	"github.com/lmnzx/asyncapi/store"
	"github.com/stretchr/testify/require"
)

func TestUser(t *testing.T) {
	userQueries := store.New(DBPool)
	ctx = context.Background()

	now := time.Now()
	userEmail := "test@test.com"
	userPassword := "testingpassword"

	user, err := userQueries.CreateUser(ctx, store.CreateUserParams{
		Email:    userEmail,
		Password: userPassword,
	})
	require.NoError(t, err)
	require.Equal(t, user.Email, userEmail)
	require.NoError(t, user.ComparePassword(userPassword))
	require.Less(t, now.UnixNano(), user.CreatedAt.UnixNano())

	userFromId, err := userQueries.GetUserById(ctx, user.ID)
	require.NoError(t, err)
	require.Equal(t, userFromId.Email, userEmail)
	require.Equal(t, userFromId.ID, user.ID)
	require.NoError(t, userFromId.ComparePassword(userPassword))
	require.Equal(t, userFromId.CreatedAt.UnixNano(), user.CreatedAt.UnixNano())

	userFromEmail, err := userQueries.GetUserByEmail(ctx, user.Email)
	require.NoError(t, err)
	require.Equal(t, userFromEmail.Email, userEmail)
	require.Equal(t, userFromEmail.ID, user.ID)
	require.NoError(t, userFromEmail.ComparePassword(userPassword))
	require.Equal(t, userFromEmail.CreatedAt.UnixNano(), user.CreatedAt.UnixNano())
}
