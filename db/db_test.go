package db

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConnect(t *testing.T) {
	err := Connect(host, port, user, password, dbname)
	require.Equal(t, err, nil, "Couldn't connect to the DB instance")
}
