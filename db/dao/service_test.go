package db

import (
	"context"
	"database/sql"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/orted-org/vyoza/util"
	"github.com/stretchr/testify/require"
)

func createRandomService(t *testing.T) Service {
	id := uuid.NewString()
	id = strings.ReplaceAll(id, "-", "")
	hash, _ := util.HashSecret(util.RandomString(64))
	arg := Service{
		ID:          id,
		Name:        util.RandomString(20),
		Description: util.RandomString(100),
		Secret:      hash,
	}

	service, err := tq.AddService(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, service)
	require.NotEmpty(t, service.CreatedAt)

	require.Equal(t, arg.ID, service.ID)
	require.Equal(t, arg.Name, service.Name)
	require.Equal(t, arg.Description, service.Description)
	require.Equal(t, arg.Secret, service.Secret)
	return service
}

func TestAddService(t *testing.T) {
	createRandomService(t)
}

func TestDeleteService(t *testing.T) {
	i := createRandomService(t)
	delErr := tq.DeleteServiceByID(context.Background(), i.ID)
	require.NoError(t, delErr)

	v, err := tq.GetServiceByID(context.Background(), i.ID)
	require.Error(t, err)
	require.Empty(t, v)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}

func TestGetServiceByID(t *testing.T) {
	i := createRandomService(t)
	v, err := tq.GetServiceByID(context.Background(), i.ID)
	require.NoError(t, err)
	require.NotEmpty(t, v)
	require.Equal(t, i.ID, v.ID)
	require.Equal(t, i.Name, v.Name)
	require.Equal(t, i.Description, v.Description)
	require.Equal(t, i.Secret, v.Secret)
	require.Equal(t, i.CreatedAt, v.CreatedAt)
}
