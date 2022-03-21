package models

type Review struct {
	ReviewId string
	UserId   string
	BookId   string
	Comment  string
	Rating   int
	Date     string
	Edited   bool
}
