package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestBcrypt(t *testing.T) {
	secret := RandomString(8)
	hs, err := HashSecret(secret)
	require.NoError(t, err)
	require.NotEmpty(t, hs)

	err = VerifyHashSecret(secret, hs)
	require.NoError(t, err)

	err = VerifyHashSecret(RandomString(8), hs)
	require.Error(t, err)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
