package db

import (
	"fmt"
	"math/rand"
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
	Title:       fmt.Sprintf("Testbook %v", rand.Int63()),
	Subtitle:    "None",
	Author:      "None",
	Published:   "1990",
	Publisher:   "test publisher",
	Pages:       44,
	Description: "description",
	ImageUrl:    "url",
	Reviews:     1,
	Rating:      rand.Float64() * 5,
}

var testUser = models.User{
	UserId:    "",
	FirstName: "Test",
	LastName:  "User",
	Gender:    "Male",
	Email:     fmt.Sprintf("test%v@user.com", rand.Int63()),
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
	IssuerId:   "issuer-id-7578",
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
