package db

import (
	"context"
	"database/sql"
	"fmt"
	// "log"
	"sort"
	"testing"
	"time"

	"github.com/orted-org/vyoza/util"
	"github.com/stretchr/testify/require"
)

func createRandomUptimeWatchRequest(t *testing.T) UptimeWatchRequest {
	arg := AddUptimeWatchRequestParams{
		Title: util.RandomString(5),
		Description: util.RandomString(10),
		Location: util.RandomString(10),
		Enabled: util.RandomBool(),
		Interval:util.RandomInt(20, 60),
		ExpectedStatus: util.RandomInt(100, 600),
		MaxResponseTime: util.RandomInt(10, 20),
		RetainDuration: util.RandomInt(1000,2000),
		HookLevel: util.RandomInt(1, 3),
		HookAddress: util.RandomString(10),
		HookSecret: string(util.NewSHA256([]byte(util.RandomString(20)))),
	}
	

	uwr, err := tq.AddUptimeWatchRequest(context.Background(),arg)

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
	require.WithinDuration(t, time.Now(), uwr.EnableUpdatedAt, time.Second)
	require.Equal(t, arg.Interval, uwr.Interval)
	require.Equal(t, arg.ExpectedStatus, uwr.ExpectedStatus)
	require.Equal(t, arg.MaxResponseTime, uwr.MaxResponseTime)
	require.Equal(t, arg.RetainDuration, uwr.RetainDuration)
	require.Equal(t, arg.HookLevel, uwr.HookLevel)
	require.Equal(t, arg.HookAddress, uwr.HookAddress)
	require.Equal(t, arg.HookSecret, uwr.HookSecret)

	return uwr;
}

func deletingTheTestingData(t *testing.T ,ID int){
	err := tq.DeleteUptimeWatchRequestById(context.Background(), ID)
	require.NoError(t, err)
}

func TestAddUptimeWatchRequest(t *testing.T){
	uwr := createRandomUptimeWatchRequest(t)
	deletingTheTestingData(t, uwr.ID);
}


func TestGetUptimeWatchRequestByID(t *testing.T){
	uwr := createRandomUptimeWatchRequest(t);

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
	require.WithinDuration(t, time.Now(), incomingUWR.EnableUpdatedAt, time.Second)
	require.Equal(t, uwr.Interval, incomingUWR.Interval)
	require.Equal(t, uwr.ExpectedStatus, incomingUWR.ExpectedStatus)
	require.Equal(t, uwr.MaxResponseTime, incomingUWR.MaxResponseTime)
	require.Equal(t, uwr.RetainDuration, incomingUWR.RetainDuration)
	require.Equal(t, uwr.HookLevel, incomingUWR.HookLevel)
	require.Equal(t, uwr.HookAddress, incomingUWR.HookAddress)
	require.Equal(t, uwr.HookSecret, incomingUWR.HookSecret)

	deletingTheTestingData(t, incomingUWR.ID);
}

func TestDeleteUptimeWatchRequestById(t *testing.T){
	uwr := createRandomUptimeWatchRequest(t);

	err := tq.DeleteUptimeWatchRequestById(context.Background(),uwr.ID)
	require.NoError(t, err)

	//checking incomingUWR for being empty of uwr.ID

	incomingUWR, err := tq.GetUptimeWatchRequestByID(context.Background(),uwr.ID)
	require.Error(t, err);
	require.Empty(t, incomingUWR)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}

/*
	1. create 10 requests, Sort them according to there id
	2. Get all the rows from db, sort them also according to there id (there 
	should be no rows, expected created above i table)
	3. start comparing
*/
func TestGetAllUptimeWatchRequest(t *testing.T){
	UWRNumber := 10
	var allUWR []UptimeWatchRequest;
	for i := 0; i <UWRNumber; i++ {
		createdUWR := createRandomUptimeWatchRequest(t);
		allUWR = append(allUWR, createdUWR)
	}
	sort.Slice(allUWR, func(i, j int) bool {
		return allUWR[i].ID < allUWR[j].ID
	})

	incomingAllUWR,err := tq.GetAllUptimeWatchRequest(context.Background()) 

	require.NoError(t, err)
	require.NotEmpty(t, incomingAllUWR)
	sort.Slice(incomingAllUWR, func(i, j int) bool {
		return incomingAllUWR[i].ID < incomingAllUWR[j].ID
	})

	for i:= 0 ; i <UWRNumber ; i++ {
		t.Run(fmt.Sprintf("Subtest Number: %v", i+1), func(t *testing.T) {
			//must be a row in incomingAllUWR, corresponding to a row in allUWR
			require.NotEmpty(t, incomingAllUWR[i]);

			require.NotZero(t, incomingAllUWR[i].ID)
			require.NotZero(t, incomingAllUWR[i].EnableUpdatedAt)

			//comparing with created data with incoming data
			require.Equal(t, allUWR[i].Title, incomingAllUWR[i].Title)
			require.Equal(t, allUWR[i].Description, incomingAllUWR[i].Description)
			require.Equal(t, allUWR[i].Location, incomingAllUWR[i].Location)
			require.Equal(t, allUWR[i].Enabled, incomingAllUWR[i].Enabled)
			require.WithinDuration(t, time.Now(), incomingAllUWR[i].EnableUpdatedAt, time.Second)
			require.Equal(t, allUWR[i].Interval, incomingAllUWR[i].Interval)
			require.Equal(t, allUWR[i].ExpectedStatus, incomingAllUWR[i].ExpectedStatus)
			require.Equal(t, allUWR[i].MaxResponseTime, incomingAllUWR[i].MaxResponseTime)
			require.Equal(t, allUWR[i].RetainDuration, incomingAllUWR[i].RetainDuration)
			require.Equal(t, allUWR[i].HookLevel, incomingAllUWR[i].HookLevel)
			require.Equal(t, allUWR[i].HookAddress, incomingAllUWR[i].HookAddress)
			require.Equal(t, allUWR[i].HookSecret, incomingAllUWR[i].HookSecret)
		})
	}

	//deleting all the created data used in testing
	for _, v := range allUWR {
		deletingTheTestingData(t, v.ID);
	}

}