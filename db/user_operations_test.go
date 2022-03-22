package db

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAddAndDeleteUser(t *testing.T) {
	err := Connect(host, port, user, password, dbname)
	defer Lib.db.Close()
	require.Equal(t, err, nil)
	userId := uuid.New().String()
	testUser.UserId = userId
	id, err := Lib.AddUser(&testUser)
	require.Equal(t, err, nil, "Unable to add user")
	require.Equal(t, *id, userId, "Invalid user id")

	// Inserting again the user with same email fails.
	userId = uuid.New().String()
	testUser.UserId = userId
	duplicateUserid, err := Lib.AddUser(&testUser)
	require.NotNilf(t, err, "Added duplicate user")
	require.Nil(t, duplicateUserid, "Invalid user id")
	deleteUserId, err := Lib.DeleteUser(*id)
	require.Nil(t, err, "unable to delete user", err)
	require.NotNil(t, deleteUserId, "Invalid id")
}

func TestSearchUser(t *testing.T) {
	err := Connect(host, port, user, password, dbname)
	defer Lib.db.Close()
	require.Nil(t, err, "unable to connect to db")
	testUser.UserId = uuid.New().String()
	Lib.AddUser(&testUser)
	resultUser, err := Lib.SearchUserByEmail(testUser.Email)
	require.Equal(t, err, nil, "Error while searching", err)
	require.Equal(t, *resultUser, testUser)
	Lib.DeleteUser(testUser.UserId)
}
