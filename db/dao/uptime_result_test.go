package db

import (
	"context"
	"testing"
	"time"

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

func TestGetUptimeResultStatsForID(t *testing.T) {
	n := 10
	uwr := createRandomUptimeWatchRequest(t)
	mustStats := UptimeResultStats{}
	mustStats.ID = uwr.ID
	successRespSum := 0
	warningRespSum := 0
	successRespCnt := 0
	warningRespCnt := 0
	for i := 0; i < n; i++ {
		arg := AddUptimeResultParams{
			ID:           uwr.ID,
			ResponseTime: util.RandomInt(-1, 4000),
		}
		res, err := tq.AddUptimeResult(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, res)
		if arg.ResponseTime <= uwr.StdResponseTime {
			successRespCnt++
			successRespSum += arg.ResponseTime
			mustStats.SuccessCount++
		} else if arg.ResponseTime > uwr.StdResponseTime && arg.ResponseTime <= uwr.MaxResponseTime {
			warningRespCnt++
			warningRespSum += arg.ResponseTime
			mustStats.WarningCount++
		} else {
			mustStats.ErrorCount++
		}
		if arg.ResponseTime <= mustStats.MinResponseTime {
			mustStats.MinResponseTime = arg.ResponseTime
		}
		if arg.ResponseTime >= mustStats.MaxResponseTime {
			mustStats.MaxResponseTime = arg.ResponseTime
		}
		if arg.ResponseTime == -1 {
			mustStats.ErrorCount++
		}
		if i == 0 {
			mustStats.StartDate = time.Now().UTC()
			mustStats.MinResponseTime = arg.ResponseTime
			mustStats.MaxResponseTime = arg.ResponseTime
		}
		if i == n-1 {
			mustStats.EndDate = time.Now().UTC()
			if successRespCnt > 0 {
				mustStats.AvgSuccessResponseTime = successRespSum / successRespCnt
			}
			if warningRespCnt > 0 {

				mustStats.AvgWarningResponseTime = warningRespSum / warningRespCnt
			}
		}
	}
	inStats, err := tq.GetUptimeResultStatsForID(context.Background(), uwr.ID)
	require.NoError(t, err)
	require.NotEmpty(t, inStats)

	require.Equal(t, mustStats.ID, inStats.ID)

	require.Equal(t, mustStats.SuccessCount, inStats.SuccessCount)
	require.Equal(t, mustStats.WarningCount, inStats.WarningCount)
	require.Equal(t, mustStats.ErrorCount, inStats.ErrorCount)

	require.Equal(t, mustStats.MinResponseTime, inStats.MinResponseTime)
	require.Equal(t, mustStats.MaxResponseTime, inStats.MaxResponseTime)

	require.Equal(t, mustStats.AvgSuccessResponseTime, inStats.AvgSuccessResponseTime)
	require.Equal(t, mustStats.AvgWarningResponseTime, inStats.AvgWarningResponseTime)

	require.WithinDuration(t, mustStats.StartDate, inStats.StartDate, time.Second)
	require.WithinDuration(t, mustStats.EndDate, inStats.EndDate, time.Second)

}
