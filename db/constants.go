package db

import (
	"time"

	"github.com/darkCavalier11/lims/models"
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
