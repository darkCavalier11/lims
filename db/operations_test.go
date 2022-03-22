package db

import (
	"testing"
	"time"

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

var testBook = models.Book{
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

var testIssueBook = models.BookIssue{
	IssueId:    "",
	UserId:     "",
	BookId:     "",
	IssueDate:  time.Now().Format(time.RFC3339),
	ReturnDate: time.Date(2050, 11, 12, 12, 25, 45, 1, time.UTC).Format(time.RFC3339),
	Returned:   false,
}

var testReview = models.Review{
	ReviewId: "",
	UserId:   "",
	BookId:   "",
	Comment:  "a test comment",
	Rating:   5,
	Date:     time.Now().Format(time.RFC3339),
	Edited:   false,
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
	Lib.DeleteBook(bookId)
	Lib.DeleteUser(userId)
}

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

	retIssueId, err := Lib.AddReview(&testReview)
	require.Equal(t, err, nil, "error adding review to the book", err)
	require.Equal(t, reviewId, *retIssueId, "invalid issue id")

	testReview.Date = time.Now().Format(time.RFC3339)
	testReview.Edited = true
	testReview.Rating = 4
	testReview.Comment = "edited comment"

	retIssueId, err = Lib.EditReview(&testReview)
	require.Equal(t, err, nil, "error deleting review from the book", err)
	require.Equal(t, reviewId, *retIssueId, "invalid issue id")

	retIssueId, err = Lib.DeleteReview(reviewId)
	require.Equal(t, err, nil, "error deleting review from the book", err)
	require.Equal(t, reviewId, *retIssueId, "invalid issue id")
	Lib.DeleteBook(bookId)
	Lib.DeleteUser(userId)
}
