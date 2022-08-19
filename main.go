package main

import (
	"fmt"
	"log"

	"github.com/darkCavalier11/lims/db"
	"github.com/darkCavalier11/lims/models"
)

var testUser = models.User{
	UserId:    "",
	FirstName: "Test",
	LastName:  "User",
	Gender:    "Male",
	Email:     "test@user.com",
	Password:  "123456",
	IsAdmin:   false,
}

const (
	host     = "localhost"
	port     = 5432
	user     = "limsdb"
	password = "password123"
	dbname   = "lims"
)

func main() {
	err := db.Connect(host, port, user, password, dbname)
	if err != nil {
		log.Fatalf("Unable to connect to db %v", err)
	}
	bookList, _ := db.Lib.SearchBook("c")
	fmt.Println(bookList[0])
}
