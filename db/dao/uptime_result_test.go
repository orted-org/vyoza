package db

import (
	"context"
	"testing"

	"github.com/orted-org/vyoza/util"
	"github.com/stretchr/testify/require"
)

func TestAddUptimeResult(t *testing.T) {

	uwr := createRandomUptimeWatchRequest(t)

	arg := AddUptimeResultParams{
		ID:           uwr.ID,
		ResponseTime: util.RandomInt(1, 3),
	}
	i, err := tq.AddUptimeResult(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, i)

	require.NotZero(t, i.CreatedAt)
	require.Equal(t, uwr.ID, i.ID)
	require.Equal(t, i.ResponseTime, arg.ResponseTime)
}
