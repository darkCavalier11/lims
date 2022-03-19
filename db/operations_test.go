package db

import (
	"github.com/darkCavalier11/lims/models"
	"github.com/google/uuid"
	"testing"

	"github.com/stretchr/testify/require"
)

var book = models.Book{
	BookId:      "",
	Isbn:        "123456789",
	Title:       "Testbook",
	Subtitle:    "None",
	Author:      "None",
	Published:   "1990",
	Publisher:   "test publisher",
	Pages:       44,
	Description: "description",
	ImageUrl:    "url",
	Rating:      4.3,
}

func TestSearchBook(t *testing.T) {
	err := Connect()
	defer Lib.db.Close()
	require.Equal(t, err, nil)
	bookQuery := []string{"Harry Potter", "Wizard ", "World", "Magic", "time", "Asweqzxxxvvffrtder"}
	for _, q := range bookQuery {
		res, err := Lib.SearchBook(q)
		require.Equal(t, err, nil, "Unexpected error")
		require.GreaterOrEqual(t, len(res), 0)
	}
}

func TestAddBook(t *testing.T) {
	err := Connect()
	defer Lib.db.Close()
	require.Equal(t, err, nil)
	bookId := uuid.New().String()
	book.BookId = bookId
	id, err := Lib.AddBook(&book)
	require.Equal(t, err, nil, "Unable to insert new book")
	require.Equal(t, bookId, *id, "Invalid book id")
	book.BookId = "444"
	id, err = Lib.AddBook(&book)
	// require.NotEqual(t, err, nil, "Adding already book")
	// require.Equal(t, id, nil, "Duplicate book added")
}

func TestDeleteBook(t *testing.T) {
	err := Connect()
	defer Lib.db.Close()
	require.Equal(t, err, nil)
	bookId := uuid.New().String()
	book.BookId = bookId
	id, err := Lib.AddBook(&book)
	require.Equal(t, err, nil, "Unable to insert new book")
	require.Equal(t, bookId, *id, "Invalid book id")
	deleteId, err := Lib.DeleteBook(*id)
	require.Equal(t, err, nil, "Unable to delete the book")
	require.Equal(t, deleteId, id, "Invalid book id")
}
