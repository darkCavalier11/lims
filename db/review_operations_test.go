package db

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestAddEditDeleteReview(t *testing.T) {
	err := Connect(host, port, user, password, dbname)
	defer Lib.db.Close()
	require.Nil(t, err, "unable to connect to db")
	userId := uuid.New().String()
	bookId := uuid.New().String()
	reviewId := uuid.New().String()
	testUser.UserId = userId
	testBook.BookId = bookId
	testReview.ReviewId = reviewId
	testReview.UserId = userId
	testReview.BookId = bookId
	Lib.AddBook(&testBook)
	Lib.AddUser(&testUser)
	defer Lib.DeleteBook(bookId)
	defer Lib.DeleteUser(userId)

	retReviewId, err := Lib.AddReview(&testReview)
	require.Equal(t, err, nil, "error adding review to the book", err)
	require.Equal(t, reviewId, *retReviewId, "invalid issue id")

	testReview.Date = time.Now().Format(time.RFC3339)
	testReview.Edited = true
	testReview.Rating = 4
	testReview.Comment = "edited comment"

	retReviewId, err = Lib.EditReview(&testReview)
	require.Equal(t, err, nil, "error editing review from the book", err)
	require.Equal(t, reviewId, *retReviewId, "invalid issue id")

	editedReview, err := Lib.GetReviewByReviewId(reviewId)
	require.Equal(t, err, nil, "error getting review from the book", err)
	require.Equal(t, testReview, *editedReview, "invalid review")

	retReviewId, err = Lib.DeleteReview(reviewId)
	require.Equal(t, err, nil, "error deleting review from the book", err)
	require.Equal(t, reviewId, *retReviewId, "invalid issue id")

}
