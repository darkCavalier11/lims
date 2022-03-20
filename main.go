package main

import (
	"fmt"
	"github.com/darkCavalier11/lims/models"
	"log"

	"github.com/darkCavalier11/lims/db"
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
	//bookQuery := flag.String("title", "", "Title of the book to query")
	//flag.Parse()
	//book, err := db.Lib.SearchBook(*bookQuery)
	//if err != nil {
	//	log.Fatalf("Unable to query book %v", err)
	//}
	//fmt.Println(book[0].BookId)
	user, err := db.Lib.AddUser(&testUser)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Println(user)
}
