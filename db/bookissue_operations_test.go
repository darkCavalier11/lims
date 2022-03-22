package db

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIssueBook(t *testing.T) {
	err := Connect(host, port, user, password, dbname)
	defer Lib.db.Close()
	require.Nil(t, err, "unable to connect to db")
	userId := uuid.New().String()
	bookId := uuid.New().String()
	issueId := uuid.New().String()
	testUser.UserId = userId
	testBook.BookId = bookId
	testIssueBook.IssueId = issueId
	testIssueBook.UserId = userId
	testIssueBook.BookId = bookId
	Lib.AddBook(&testBook)
	Lib.AddUser(&testUser)
	defer Lib.DeleteBook(bookId)
	defer Lib.DeleteUser(userId)

	// Issue a book
	retIssueId, err := Lib.IssueBook(&testIssueBook)
	require.Equal(t, err, nil, "error issuing book", err)
	require.Equal(t, issueId, *retIssueId, "invalid issue id")

	// Check the availability of book
	isAvailable, resIssueId, err := Lib.CheckBookAvailability(bookId)
	require.Nil(t, err, "error ", err)
	require.False(t, *isAvailable, "book is still available")
	require.Equal(t, *resIssueId, issueId, "invalid book id")

	// Unissue the book
	unIssueId, err := Lib.ReturnBook(bookId)
	require.Nil(t, err, err)
	require.Equal(t, *unIssueId, issueId, "invalid issue id")
	isAvailable, retIssueId, err = Lib.CheckBookAvailability(bookId)
	require.Nil(t, err, "error ", err)
	require.Nil(t, retIssueId, "invalid issue id", err)
	require.True(t, *isAvailable, "book is  unavailable", err)

}
