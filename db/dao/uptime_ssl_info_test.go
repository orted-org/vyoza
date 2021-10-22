package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/orted-org/vyoza/util"
	"github.com/stretchr/testify/require"
)

func createRandomUptimeSSLInfo(t *testing.T) UptimeSSLInfo {
	uwr := createRandomUptimeWatchRequest(t)
	arg := UptimeSSLInfo{
		UWRID:      uwr.ID,
		IsValid:    util.RandomBool(),
		ExpiryDate: time.Now().AddDate(0, 0, 5).UTC(),
		Remark:     util.RandomString(10),
		UpdatedAt:  time.Now().UTC(),
	}

	createdUptimeSSLInfo, err := tq.AddUptimeSSLInfo(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, createdUptimeSSLInfo)

	require.Equal(t, arg.UWRID, createdUptimeSSLInfo.UWRID)
	require.Equal(t, arg.IsValid, createdUptimeSSLInfo.IsValid)
	require.Equal(t, arg.ExpiryDate, createdUptimeSSLInfo.ExpiryDate)
	require.Equal(t, arg.Remark, createdUptimeSSLInfo.Remark)
	require.Equal(t, arg.UpdatedAt, createdUptimeSSLInfo.UpdatedAt)
	return createdUptimeSSLInfo
}

func TestAddUptimeSSLInfo(t *testing.T) {
	createRandomUptimeSSLInfo(t)
}

func TestDeleteUptimeSSLInfoByUWRID(t *testing.T) {
	usi := createRandomUptimeSSLInfo(t)
	delErr := tq.DeleteUptimeSSLInfoByUWRID(context.Background(), usi.UWRID)

	require.NoError(t, delErr)

	incomingUsi, err := tq.GetUptimeSSLInfoByUWRID(context.Background(), usi.UWRID)
	require.Error(t, err)
	require.Empty(t, incomingUsi)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}

func TestUpdateUptimeSSLInfoByUWRID(t *testing.T) {
	usi := createRandomUptimeSSLInfo(t)

	updatedUsiArg := UptimeSSLInfo{
		UWRID:      usi.UWRID,
		IsValid:    util.RandomBool(),
		ExpiryDate: time.Now().AddDate(0, 0, 10).UTC(),
		Remark:     util.RandomString(10),
		UpdatedAt:  time.Now().UTC(),
	}
	incomingUsi, err := tq.UpdateUptimeSSLInfoByUWRID(context.Background(), updatedUsiArg)
	require.NoError(t, err)
	require.NotEmpty(t, incomingUsi)

	require.Equal(t, updatedUsiArg.UWRID, incomingUsi.UWRID)
	require.Equal(t, updatedUsiArg.IsValid, incomingUsi.IsValid)
	require.Equal(t, updatedUsiArg.ExpiryDate, incomingUsi.ExpiryDate)
	require.Equal(t, updatedUsiArg.Remark, incomingUsi.Remark)
	require.Equal(t, updatedUsiArg.UpdatedAt, incomingUsi.UpdatedAt)
}

func TestGetUptimeSSLInfoByUWRID(t *testing.T) {
	usi := createRandomUptimeSSLInfo(t)
	incomingUsi, err := tq.GetUptimeSSLInfoByUWRID(context.Background(), usi.UWRID)
	require.NoError(t, err)
	require.NotEmpty(t, incomingUsi)
	require.Equal(t, usi.UWRID, incomingUsi.UWRID)
	require.Equal(t, usi.IsValid, incomingUsi.IsValid)
	require.Equal(t, usi.ExpiryDate, incomingUsi.ExpiryDate)
	require.Equal(t, usi.Remark, incomingUsi.Remark)
	require.Equal(t, usi.UpdatedAt, incomingUsi.UpdatedAt)
}

func TestGetAllUptimeSSLInfo(t *testing.T) {
	n := 10
	var insertedUSI []UptimeSSLInfo
	for i := 0; i < n; i++ {
		insertedUSI = append(insertedUSI, createRandomUptimeSSLInfo(t))
	}

	limit := 3
	offset := 0
	matchedCount := 0

	for {
		incomingUSIGroup, err := tq.GetAllUptimeSSLInfo(context.Background(), getAllUptimeSSLInfoParams{
			Limit:  limit,
			Offset: offset,
		})
		require.NoError(t, err)
		if len(incomingUSIGroup) == 0 {
			break
		}
		for _, insertedUSI := range insertedUSI {
			for _, incmgUSI := range incomingUSIGroup {
				if insertedUSI.UWRID == incmgUSI.UWRID {
					//matched
					matchedCount++
					//Checking the equality of fields
					require.Equal(t, insertedUSI.UWRID, incmgUSI.UWRID)
					require.Equal(t, insertedUSI.IsValid, incmgUSI.IsValid)
					require.Equal(t, insertedUSI.ExpiryDate, incmgUSI.ExpiryDate)
					require.Equal(t, insertedUSI.Remark, incmgUSI.Remark)
					require.Equal(t, insertedUSI.UpdatedAt, incmgUSI.UpdatedAt)
				}
			}
		}
		offset = offset + limit
	}

	require.Equal(t, n, matchedCount)
}
