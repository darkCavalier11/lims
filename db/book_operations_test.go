package db

import (
	"fmt"
	"github.com/darkCavalier11/lims/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
	"time"
)

func TestAddBook(t *testing.T) {
	err := Connect(host, port, user, password, dbname)

	defer Lib.db.Close()
	require.Equal(t, err, nil)
	bookId := uuid.New().String()
	testBook.BookId = bookId
	id, err := Lib.AddBook(&testBook)
	require.Equal(t, err, nil, "Unable to insert new book")
	require.Equal(t, bookId, *id, "Invalid book id")

	// Fail adding duplicate book
	id, err = Lib.AddBook(&testBook)
	require.NotEqual(t, err, nil, "Adding already book")
	require.Nil(t, id)
}

func TestDeleteBook(t *testing.T) {
	err := Connect(host, port, user, password, dbname)

	defer Lib.db.Close()
	require.Equal(t, err, nil)
	bookId := uuid.New().String()
	testBook.BookId = bookId
	id, err := Lib.AddBook(&testBook)
	require.Equal(t, err, nil, "Unable to insert new book")
	require.Equal(t, bookId, *id, "Invalid book id")
	deleteId, err := Lib.DeleteBook(*id)
	require.Equal(t, err, nil, "Unable to delete the book")
	require.Equal(t, deleteId, id, "Invalid book id")
}

func TestSearchBook(t *testing.T) {
	err := Connect(host, port, user, password, dbname)
	defer Lib.db.Close()
	require.Equal(t, err, nil)
	bookQuery := []string{"Harry Potter", "Wizard ", "World", "Magic", "time", "Asweqzxxxvvffrtder"}
	for _, q := range bookQuery {
		res, err := Lib.SearchBook(q)
		require.Equal(t, err, nil, "Unexpected error")
		require.GreaterOrEqual(t, len(res), 0)
	}
}

func TestGetReviewsOfBook(t *testing.T) {
	err := Connect(host, port, user, password, dbname)
	defer Lib.db.Close()
	require.Equal(t, err, nil)
	bookId := uuid.New().String()
	userId := uuid.New().String()
	testBook.BookId = bookId
	testUser.UserId = userId
	Lib.AddBook(&testBook)
	Lib.AddUser(&testUser)
	defer Lib.DeleteBook(bookId)
	defer Lib.DeleteUser(userId)
	reviews := []*models.Review{}
	for i := 0; i < 10; i++ {
		var bookTestReview models.Review
		bookTestReview.ReviewId = uuid.New().String()
		bookTestReview.UserId = userId
		bookTestReview.BookId = bookId
		bookTestReview.Comment = fmt.Sprintf("comment no %v", i)
		bookTestReview.Rating = rand.Intn(6)
		bookTestReview.Edited = false
		bookTestReview.Date = time.Now().Format(time.RFC3339)
		reviews = append(reviews, &bookTestReview)
	}
	for _, r := range reviews {
		Lib.AddReview(r)
	}
	retReviews, err := Lib.GetReviewsOfBook(bookId)
	require.Equal(t, err, nil, err)
	require.EqualValues(t, retReviews, reviews, "invalid reviews")
}
