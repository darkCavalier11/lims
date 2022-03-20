package db

import (
	"testing"

	"github.com/darkCavalier11/lims/models"
	"github.com/google/uuid"

	"github.com/stretchr/testify/require"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "limsdb"
	password = "password123"
	dbname   = "lims"
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

var testUser = models.User{
	UserId:    "",
	FirstName: "Test",
	LastName:  "User",
	Gender:    "Male",
	Email:     "test@user.com",
	Password:  "123456",
	IsAdmin:   false,
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

func TestAddBook(t *testing.T) {
	err := Connect(host, port, user, password, dbname)

	defer Lib.db.Close()
	require.Equal(t, err, nil)
	bookId := uuid.New().String()
	book.BookId = bookId
	id, err := Lib.AddBook(&book)
	require.Equal(t, err, nil, "Unable to insert new book")
	require.Equal(t, bookId, *id, "Invalid book id")

	// Fail adding duplicate book
	id, err = Lib.AddBook(&book)
	require.NotEqual(t, err, nil, "Adding already book")
	require.Nil(t, id)
}

func TestDeleteBook(t *testing.T) {
	err := Connect(host, port, user, password, dbname)

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
