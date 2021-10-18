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

func TestGetUptimeResults(t *testing.T) {
	n := 10
	var allUR []UptimeResult
	var savedRes []UptimeResult
	var avgCreated = 0
	var avgSaved = 0
	uwr := createRandomUptimeWatchRequest(t)
	for i := 0; i < n; i++ {
		arg := AddUptimeResultParams{
			ID:           uwr.ID,
			ResponseTime: util.RandomInt(1, 10),
		}
		res, err := tq.AddUptimeResult(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, res)
		allUR = append(allUR, res)
		avgCreated = avgCreated + res.ResponseTime
	}

	i := 0
	for {
		res, err := tq.GetUptimeResults(context.Background(), GetUptimeResultsParams{
			ID:     uwr.ID,
			Offset: i * 2,
			Limit:  2,
		})
		require.NoError(t, err)
		if len(res) == 0 {
			break
		}
		savedRes = append(savedRes, res...)
		for _, inv := range res {
			avgSaved = avgSaved + inv.ResponseTime
		}
		i++
	}

	require.Len(t, savedRes, len(allUR))
	require.Equal(t, avgSaved/n, avgCreated/n)

}

func TestDeleteUptimeResults(t *testing.T) {
	n := 10
	var allUR []UptimeResult
	var savedRes []UptimeResult
	var avgCreated = 0
	var avgSaved = 0
	uwr := createRandomUptimeWatchRequest(t)
	for i := 0; i < n; i++ {
		arg := AddUptimeResultParams{
			ID:           uwr.ID,
			ResponseTime: util.RandomInt(1, 10),
		}
		res, err := tq.AddUptimeResult(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, res)
		allUR = append(allUR, res)
		avgCreated = avgCreated + res.ResponseTime
	}

	err := tq.DeleteUptimeResults(context.Background(), uwr.ID)
	require.NoError(t, err)

	i := 0
	for {
		res, err := tq.GetUptimeResults(context.Background(), GetUptimeResultsParams{
			ID:     uwr.ID,
			Offset: i * 2,
			Limit:  2,
		})
		require.NoError(t, err)
		if len(res) == 0 {
			break
		}
		savedRes = append(savedRes, res...)
		for _, inv := range res {
			avgSaved = avgSaved + inv.ResponseTime
		}
		i++
	}

	require.Len(t, allUR, n)
	require.Len(t, savedRes, 0)
	require.Equal(t, avgSaved/n, 0)
}
