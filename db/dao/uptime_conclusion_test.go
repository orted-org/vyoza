package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/orted-org/vyoza/util"
	"github.com/stretchr/testify/require"
)

func createRandomUptimeConclusion(t *testing.T) UptimeConclusion {
	n := 10
	uwr := createRandomUptimeWatchRequest(t)
	for i := 0; i < n; i++ {
		randomResponeTime := util.RandomInt(-1, 4000)
		if randomResponeTime > uwr.MaxResponseTime {
			randomResponeTime = -1
		}
		arg := AddUptimeResultParams{
			UWRID:        uwr.ID,
			ResponseTime: randomResponeTime,
			Remark:       util.RandomString(10),
		}
		_, err := tq.AddUptimeResult(context.Background(), arg)
		require.NoError(t, err)
	}
	stateEx1, err := tq.GetUptimeResultStatsForID(context.Background(), uwr.ID)
	require.NoError(t, err)

	var uc UptimeConclusion
	uc.UWRID = stateEx1.UWRID
	uc.SuccessCount = stateEx1.SuccessCount
	uc.WarningCount = stateEx1.WarningCount
	uc.ErrorCount = stateEx1.ErrorCount
	uc.MinResponseTime = stateEx1.MinResponseTime
	uc.MaxResponseTime = stateEx1.MaxResponseTime
	uc.AvgSuccessResponseTime = stateEx1.AvgSuccessResponseTime
	uc.AvgWarningResponseTime = stateEx1.AvgWarningResponseTime
	uc.StartDate = stateEx1.StartDate
	uc.EndDate = stateEx1.EndDate

	inuc, err := tq.AddUptimeConclusion(context.Background(), uc)

	require.NoError(t, err)
	require.NotEmpty(t, inuc)

	require.Equal(t, uc.UWRID, inuc.UWRID)
	require.Equal(t, uc.SuccessCount, inuc.SuccessCount)
	require.Equal(t, uc.WarningCount, inuc.WarningCount)
	require.Equal(t, uc.ErrorCount, inuc.ErrorCount)
	require.Equal(t, uc.MinResponseTime, inuc.MinResponseTime)
	require.Equal(t, uc.MaxResponseTime, inuc.MaxResponseTime)
	require.Equal(t, uc.AvgSuccessResponseTime, inuc.AvgSuccessResponseTime)
	require.Equal(t, uc.AvgWarningResponseTime, inuc.AvgWarningResponseTime)
	require.WithinDuration(t, uc.StartDate, inuc.StartDate, time.Second)
	require.WithinDuration(t, uc.EndDate, inuc.EndDate, time.Second)

	return inuc
}

func TestAddUptimeConclusion(t *testing.T) {
	createRandomUptimeConclusion(t)
}

func TestGetUptimeConclusionByUWRID(t *testing.T) {
	uc := createRandomUptimeConclusion(t)
	inuc, err := tq.GetUptimeConclusionByUWRID(context.Background(), uc.UWRID)

	require.NoError(t, err)
	require.NotEmpty(t, inuc)

	require.Equal(t, uc.UWRID, inuc.UWRID)
	require.Equal(t, uc.SuccessCount, inuc.SuccessCount)
	require.Equal(t, uc.WarningCount, inuc.WarningCount)
	require.Equal(t, uc.ErrorCount, inuc.ErrorCount)
	require.Equal(t, uc.MinResponseTime, inuc.MinResponseTime)
	require.Equal(t, uc.MaxResponseTime, inuc.MaxResponseTime)
	require.Equal(t, uc.AvgSuccessResponseTime, inuc.AvgSuccessResponseTime)
	require.Equal(t, uc.AvgWarningResponseTime, inuc.AvgWarningResponseTime)
	require.WithinDuration(t, uc.StartDate, inuc.StartDate, time.Second)
	require.WithinDuration(t, uc.EndDate, inuc.EndDate, time.Second)
}

func TestDeleteUptimeConclusionByUWRID(t *testing.T) {
	uc := createRandomUptimeConclusion(t)
	err := tq.DeleteUptimeConclusionByUWRID(context.Background(), uc.UWRID)

	require.NoError(t, err)

	inuc, getErr := tq.GetUptimeConclusionByUWRID(context.Background(), uc.UWRID)
	require.Error(t, getErr)
	require.Empty(t, inuc)
	require.EqualError(t, getErr, sql.ErrNoRows.Error())
}

/*
Algorithm
1. insert n new rows there
2. Get all in groups
3. Check all the newly created rows should be in all the groups
if yes then we are sorted
*/

func TestGetAllUptimeConclusion(t *testing.T) {
	n := 10
	var insertedUCs []UptimeConclusion
	for i := 0; i < n; i++ {
		insertedUCs = append(insertedUCs, createRandomUptimeConclusion(t))
	}
	limit := 3
	offset := 0
	matchedCount := 0

	for {
		incomingUCGroup, err := tq.GetAllUptimeConclusion(context.Background(), getAllUptimeConclusionParams{
			Limit:  limit,
			Offset: offset,
		})
		require.NoError(t, err)
		if len(incomingUCGroup) == 0 {
			break
		}
		for _, insrtUC := range insertedUCs {
			for _, incmgUC := range incomingUCGroup {
				if insrtUC.UWRID == incmgUC.UWRID {
					//matched
					matchedCount++
					//Checking the equality of fields
					require.Equal(t, insrtUC.UWRID, incmgUC.UWRID)
					require.Equal(t, insrtUC.SuccessCount, incmgUC.SuccessCount)
					require.Equal(t, insrtUC.WarningCount, incmgUC.WarningCount)
					require.Equal(t, insrtUC.ErrorCount, incmgUC.ErrorCount)
					require.Equal(t, insrtUC.MinResponseTime, incmgUC.MinResponseTime)
					require.Equal(t, insrtUC.MaxResponseTime, incmgUC.MaxResponseTime)
					require.Equal(t, insrtUC.AvgSuccessResponseTime, incmgUC.AvgSuccessResponseTime)
					require.Equal(t, insrtUC.AvgWarningResponseTime, incmgUC.AvgWarningResponseTime)
					require.WithinDuration(t, insrtUC.StartDate, incmgUC.StartDate, time.Second)
					require.WithinDuration(t, insrtUC.EndDate, incmgUC.EndDate, time.Second)
				}
			}
		}
		offset = offset + limit
	}

	require.Equal(t, n, matchedCount)
}
