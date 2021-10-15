package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSHA256(t *testing.T) {
	text := []byte(RandomString(10))
	sha256 := NewSHA256(text)
	err := CompareSHA256(text, sha256)
	require.NotEmpty(t, text)
	require.NoError(t, err)
}
