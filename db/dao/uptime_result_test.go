package db

import (
	"context"
	"log"
	"testing"

	"github.com/orted-org/vyoza/util"
	"github.com/stretchr/testify/require"
)

func createRamdomUptimeResult(t *testing.T) UptimeResult {
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
	return i
}
func TestAddUptimeResult(t *testing.T) {
	createRamdomUptimeResult(t)
}

func TestGetUptimeResultCount(t *testing.T) {
	i := createRamdomUptimeResult(t)
	c, err := tq.GetUptimeResultCount(context.Background(), i.ID)
	log.Println(c)
	require.NoError(t, err)
	require.NotEmpty(t, c)
	require.NotZero(t, c)
}
