package db

import (
	"database/sql"
	"fmt"
	"github.com/darkCavalier11/lims/models"
)

// IssueBook issues a book to an user
func (lib *library) IssueBook(issue *models.BookIssue) (id *string, e error) {
	var issueId string
	sqlStatement := `INSERT INTO bookissue (issue_id, user_id, book_id, issue_date, return_date, returned) VALUES ($1, $2, $3, $4, $5, $6) RETURNING issue_id`
	err := lib.db.QueryRow(sqlStatement, issue.IssueId, issue.UserId, issue.BookId, issue.IssueDate, issue.ReturnDate, issue.Returned).Scan(&issueId)
	if err != nil {
		return nil, fmt.Errorf("unable to issue book %w", err)
	}
	return &issueId, nil
}

func (lib *library) CheckBookAvailability(bookId string) (*bool, *string, error) {
	var issueId string
	var isAvailable bool = true
	sqlStatement := `SELECT issue_id, returned FROM bookissue where book_id = $1`
	err := lib.db.QueryRow(sqlStatement, bookId).Scan(&issueId, &isAvailable)
	if err == sql.ErrNoRows {
		isAvailable = true
		return &isAvailable, nil, nil
	}
	if err != nil {
		return nil, nil, err
	}
	if !isAvailable {
		return &isAvailable, &issueId, nil
	}
	return &isAvailable, nil, nil
}

func (lib *library) ReturnBook(bookId string) (*string, error) {
	var issueId string
	sqlStatement := `UPDATE bookissue set returned = $1 where book_id = $2 returning issue_id`
	err := lib.db.QueryRow(sqlStatement, true, bookId).Scan(&issueId)
	if err != nil {
		return nil, err
	}
	return &issueId, nil
}
