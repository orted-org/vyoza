package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomAlphaNumeric(t *testing.T) {
	n := 16
	str := RandomAlphaNumeric(n)
	fmt.Println(str)
	require.NotEmpty(t, str)
	require.Len(t, str, n)
}
