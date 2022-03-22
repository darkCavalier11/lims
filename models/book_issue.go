package models

type BookIssue struct {
	IssueId    string
	UserId     string
	BookId     string
	IssuerId   string
	IssueDate  string
	ReturnDate string
	Returned   bool
}
