package db

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConnect(t *testing.T) {
	err := Connect()
	require.Equal(t, err, nil, "Couldn't connect to the DB instance")
}
