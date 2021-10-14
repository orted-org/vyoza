package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/orted-org/vyoza/util"
	"github.com/stretchr/testify/require"
)

func createRandomKeyValue(t *testing.T) KeyValue {
	arg := KeyValue{Key: util.RandomString(5), Value: util.RandomString(5)}
	kv, err := tq.AddKeyValue(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, kv)
	require.NotZero(t, kv.Key)
	require.NotZero(t, kv.Value)
	require.NotZero(t, kv.UpdateAt)
	require.Equal(t, arg.Key, kv.Key)
	require.Equal(t, arg.Value, kv.Value)
	require.WithinDuration(t, time.Now(), kv.UpdateAt, time.Second)
	return kv
}

func TestSetKeyValue(t *testing.T) {
	createRandomKeyValue(t)
}
func TestUpdateKeyValue(t *testing.T) {
	kv := createRandomKeyValue(t)
	arg := KeyValue{
		Key:   kv.Key,
		Value: util.RandomString(5),
	}
	updatedKv, err := tq.UpdateKeyValue(context.Background(), arg)

	// checking for incoming kv for intactness
	require.NoError(t, err)
	require.NotEmpty(t, updatedKv)
	require.NotZero(t, updatedKv.Key)
	require.NotZero(t, updatedKv.Value)
	require.NotZero(t, updatedKv.UpdateAt)

	// comparing arg with updated kv
	require.Equal(t, arg.Key, updatedKv.Key)
	require.Equal(t, arg.Value, updatedKv.Value)
	require.WithinDuration(t, time.Now(), updatedKv.UpdateAt, time.Second)
}

func TestGetKeyValue(t *testing.T) {
	kv := createRandomKeyValue(t)
	incomingKv, err := tq.GetKeyValue(context.Background(), kv.Key)

	// checking the incoming kv for intactness
	require.NoError(t, err)
	require.NotEmpty(t, incomingKv)
	require.NotZero(t, incomingKv.Key)
	require.NotZero(t, incomingKv.Value)
	require.NotZero(t, incomingKv.UpdateAt)

	// comparing with the created kv
	require.Equal(t, kv.Key, incomingKv.Key)
	require.Equal(t, kv.Value, incomingKv.Value)
}

func TestDeleteKeyValue(t *testing.T) {
	kv := createRandomKeyValue(t)
	err := tq.DeleteKeyValue(context.Background(), kv.Key)
	require.NoError(t, err)

	// checking the incoming kv for being empty
	incomingKv, err := tq.GetKeyValue(context.Background(), kv.Key)
	require.Error(t, err)
	require.Empty(t, incomingKv)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}
