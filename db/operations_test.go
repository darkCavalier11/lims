package db

import (
	"github.com/darkCavalier11/lims/models"
	"github.com/google/uuid"
	"testing"

	"github.com/stretchr/testify/require"
)

var bookId = uuid.New().String()

var book = models.Book{
	bookId,
	"123456789",
	"Testbook",
	"None",
	"None",
	"1990",
	"test publisher",
	44,
	"description",
	"url",
	4.3,
}

func TestSearchBook(t *testing.T) {
	err := Connect()
	defer Lib.db.Close()
	require.Equal(t, err, nil)
	bookQuery := []string{"Harry Potter", "Wizard ", "World", "Magic", "time"}
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
	id, err := Lib.AddBook(&book)
	require.Equal(t, err, nil, "Unable to insert new book")
	require.Equal(t, bookId, *id, "Invalid book id")
}

func TestDeleteBook(t *testing.T) {

}
