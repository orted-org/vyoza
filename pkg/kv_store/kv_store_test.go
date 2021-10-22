package kvstore

import (
	"testing"

	"github.com/orted-org/vyoza/util"
	"github.com/stretchr/testify/require"
)

func TestKVStore(t *testing.T) {
	i := New()
	val, err := i.Get(util.RandomString(10))
	require.Error(t, err)
	require.EqualError(t, err, ErrKeyValueNotExists.Error())
	require.Empty(t, val)

	// set
	argKey := util.RandomString(10)
	argValue := util.RandomString(10)
	err = i.Set(argKey, argValue)
	require.NoError(t, err)

	// get
	val, err = i.Get(argKey)
	require.NoError(t, err)
	require.NotEmpty(t, val)
	require.Equal(t, argValue, val)

	// delete
	err = i.Delete(argKey)
	require.NoError(t, err)
	val, err = i.Get(argKey)
	require.Error(t, err)
	require.EqualError(t, err, ErrKeyValueNotExists.Error())
	require.Empty(t, val)

	// turncate
	err = i.Set(argKey, argValue)
	require.NoError(t, err)

	err = i.Truncate()
	require.NoError(t, err)
	val, err = i.Get(argKey)
	require.Error(t, err)
	require.EqualError(t, err, ErrKeyValueNotExists.Error())
	require.Empty(t, val)
}
