package test

import (
	"context"
	"testing"

	"github.com/lmnzx/asyncapi/store"
	"github.com/stretchr/testify/require"
)

func TestUser(t *testing.T) {
	userQueries := store.New(DBPool)
	ctx = context.Background()

	userEmail := "test@test.com"
	userPassword := "testingpassword"

	user, err := userQueries.CreateUser(ctx, store.CreateUserParams{
		Email:    userEmail,
		Password: userPassword,
	})
	require.NoError(t, err)
	require.Equal(t, user.Email, userEmail)
	require.NoError(t, user.ComparePassword(userPassword))
}
