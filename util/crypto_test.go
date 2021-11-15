package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestSHA256(t *testing.T) {
	text := []byte(RandomString(10))
	sha256 := NewSHA256(text)
	err := CompareSHA256(text, sha256)
	require.NotEmpty(t, text)
	require.NoError(t, err)
}

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

func TestTextEncryption(t *testing.T) {
	text := []byte(RandomString(10))
	e := EncryptText(text)
	require.NotEmpty(t, e)

	d := DecryptText(e)
	require.NotEmpty(t, d)
	require.Equal(t, text, d)
}
