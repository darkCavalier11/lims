package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/darkCavalier11/lims/db"
)

func main() {
	err := db.Connect()
	if err != nil {
		log.Fatalf("Unable to connect to db %v", err)
	}
	bookQuery := flag.String("title", "", "Title of the book to query")
	flag.Parse()
	book, err := db.Lib.SearchBook(*bookQuery)
	if err != nil {
		log.Fatalf("Unable to query book %v", err)
	}
	fmt.Println(book[0].BookId)
}
