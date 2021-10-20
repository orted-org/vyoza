package db

import (
	"context"
	"database/sql"
	"fmt"

	"testing"
	"time"

	"github.com/orted-org/vyoza/util"
	"github.com/stretchr/testify/require"
)

func createRandomUptimeWatchRequest(t *testing.T) UptimeWatchRequest {
	arg := AddUptimeWatchRequestParams{
		Title:           util.RandomString(5),
		Description:     util.RandomString(10),
		Location:        util.RandomString(10),
		Enabled:         util.RandomBool(),
		Interval:        util.RandomInt(20, 60),
		ExpectedStatus:  util.RandomInt(100, 600),
		StdResponseTime: util.RandomInt(500, 1500),
		MaxResponseTime: util.RandomInt(1501, 3000),
		RetainDuration:  util.RandomInt(1000, 2000),
		HookLevel:       util.RandomInt(1, 3),
		HookAddress:     util.RandomString(10),
		HookSecret:      string(util.NewSHA256([]byte(util.RandomString(20)))),
	}

	uwr, err := tq.AddUptimeWatchRequest(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, uwr)

	//checking notZero/notNil only for fields set automatically
	require.NotZero(t, uwr.ID)
	require.NotZero(t, uwr.EnableUpdatedAt)

	//comparing with created data
	require.Equal(t, arg.Title, uwr.Title)
	require.Equal(t, arg.Description, uwr.Description)
	require.Equal(t, arg.Location, uwr.Location)
	require.Equal(t, arg.Enabled, uwr.Enabled)
	require.WithinDuration(t, time.Now().UTC(), uwr.EnableUpdatedAt, time.Second)
	require.Equal(t, arg.Interval, uwr.Interval)
	require.Equal(t, arg.ExpectedStatus, uwr.ExpectedStatus)
	require.Equal(t, arg.MaxResponseTime, uwr.MaxResponseTime)
	require.Equal(t, arg.RetainDuration, uwr.RetainDuration)
	require.Equal(t, arg.HookLevel, uwr.HookLevel)
	require.Equal(t, arg.HookAddress, uwr.HookAddress)
	require.Equal(t, arg.HookSecret, uwr.HookSecret)

	return uwr
}

func deletingTheTestingData(t *testing.T, ID int) {
	err := tq.DeleteUptimeWatchRequestById(context.Background(), ID)
	require.NoError(t, err)
}

func TestAddUptimeWatchRequest(t *testing.T) {
	uwr := createRandomUptimeWatchRequest(t)
	deletingTheTestingData(t, uwr.ID)
}

func TestGetUptimeWatchRequestByID(t *testing.T) {
	uwr := createRandomUptimeWatchRequest(t)

	incomingUWR, err := tq.GetUptimeWatchRequestByID(context.Background(), uwr.ID)

	require.NoError(t, err)
	require.NotEmpty(t, incomingUWR)

	//checking notZero/notNil only for fields set automatically
	require.NotZero(t, incomingUWR.ID)
	require.NotZero(t, incomingUWR.EnableUpdatedAt)

	//comparing with created data
	require.Equal(t, uwr.Title, incomingUWR.Title)
	require.Equal(t, uwr.Description, incomingUWR.Description)
	require.Equal(t, uwr.Location, incomingUWR.Location)
	require.Equal(t, uwr.Enabled, incomingUWR.Enabled)
	require.WithinDuration(t, time.Now().UTC(), incomingUWR.EnableUpdatedAt, time.Second)
	require.Equal(t, uwr.Interval, incomingUWR.Interval)
	require.Equal(t, uwr.ExpectedStatus, incomingUWR.ExpectedStatus)
	require.Equal(t, uwr.MaxResponseTime, incomingUWR.MaxResponseTime)
	require.Equal(t, uwr.RetainDuration, incomingUWR.RetainDuration)
	require.Equal(t, uwr.HookLevel, incomingUWR.HookLevel)
	require.Equal(t, uwr.HookAddress, incomingUWR.HookAddress)
	require.Equal(t, uwr.HookSecret, incomingUWR.HookSecret)

	deletingTheTestingData(t, incomingUWR.ID)
}

func TestDeleteUptimeWatchRequestById(t *testing.T) {
	uwr := createRandomUptimeWatchRequest(t)

	err := tq.DeleteUptimeWatchRequestById(context.Background(), uwr.ID)
	require.NoError(t, err)

	//checking incomingUWR for being empty of uwr.ID

	incomingUWR, err := tq.GetUptimeWatchRequestByID(context.Background(), uwr.ID)
	require.Error(t, err)
	require.Empty(t, incomingUWR)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}

/*
	1. create 10 requests, Sort them according to there id
	2. Get all the rows from db, sort them also according to there id (there
	should be no rows, expected created above i table)
	3. start comparing
*/
func TestGetAllUptimeWatchRequest(t *testing.T) {
	UWRNumber := 10
	var allUWR []UptimeWatchRequest
	for i := 0; i < UWRNumber; i++ {
		createdUWR := createRandomUptimeWatchRequest(t)
		allUWR = append(allUWR, createdUWR)
	}

	incomingAllUWR, err := tq.GetAllUptimeWatchRequest(context.Background())

	require.NoError(t, err)
	require.NotEmpty(t, incomingAllUWR)

	for k, oneFromCreated := range allUWR {
		t.Run(fmt.Sprintf("Subtest Number: %v", k+1), func(t *testing.T) {
			//must be a row in incomingAllUWR, corresponding to a row in allUWR
			var i UptimeWatchRequest
			for _, oneFromAll := range incomingAllUWR {
				if oneFromCreated.ID == oneFromAll.ID {
					i = oneFromAll
				}
			}
			require.NotEmpty(t, i)
			require.NotZero(t, i.ID)
			require.NotZero(t, i.EnableUpdatedAt)

			//comparing with created data with incoming data
			require.Equal(t, oneFromCreated.Title, i.Title)
			require.Equal(t, oneFromCreated.Description, i.Description)
			require.Equal(t, oneFromCreated.Location, i.Location)
			require.Equal(t, oneFromCreated.Enabled, i.Enabled)
			require.WithinDuration(t, time.Now().UTC(), i.EnableUpdatedAt, time.Second)
			require.Equal(t, oneFromCreated.Interval, i.Interval)
			require.Equal(t, oneFromCreated.ExpectedStatus, i.ExpectedStatus)
			require.Equal(t, oneFromCreated.MaxResponseTime, i.MaxResponseTime)
			require.Equal(t, oneFromCreated.RetainDuration, i.RetainDuration)
			require.Equal(t, oneFromCreated.HookLevel, i.HookLevel)
			require.Equal(t, oneFromCreated.HookAddress, i.HookAddress)
			require.Equal(t, oneFromCreated.HookSecret, i.HookSecret)
		})
	}

	//deleting all the created data used in testing
	for _, v := range allUWR {
		deletingTheTestingData(t, v.ID)
	}

}

func TestUpdateUptimeWatchRequestById(t *testing.T) {
	i := createRandomUptimeWatchRequest(t)
	updates := make(map[string]interface{})
	arg := UptimeWatchRequest{
		Title:       util.RandomString(10),
		Description: util.RandomString(40),
		HookLevel:   util.RandomInt(1, 3),
	}
	updates["title"] = arg.Title
	updates["description"] = arg.Description
	updates["hook_level"] = arg.HookLevel

	updated, err := tq.UpdateUptimeWatchRequestById(context.Background(), updates, i.ID)
	require.NoError(t, err)
	require.NotEmpty(t, updated)

	require.Equal(t, updated.Title, arg.Title)
	require.Equal(t, updated.Description, arg.Description)
	require.Equal(t, updated.HookLevel, arg.HookLevel)
	deletingTheTestingData(t, i.ID)

}
