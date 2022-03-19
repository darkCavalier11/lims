package models

type User struct {
	UserId    string
	FirstName string
	LastName  string
	Gender    string
	Email     string
	Password  string
	IsAdmin   bool
}
