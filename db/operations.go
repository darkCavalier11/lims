package db

import (
	"database/sql"
	"fmt"
	"github.com/darkCavalier11/lims/models"
)

// SearchBook search a book based on title matching keyword
func (lib *library) SearchBook(keyword string) (bookList []*models.Book, err error) {
	book := &models.Book{}
	sqlStatement := `SELECT * FROM book WHERE title LIKE $1`
	rows, err := lib.db.Query(sqlStatement, "%"+keyword+"%")
	defer func(rows *sql.Rows) {
		if rows == nil {
			return
		}
		e := rows.Close()
		if e != nil {
			err = e
		}
	}(rows)
	for rows.Next() {
		err := rows.Scan(&book.BookId, &book.Isbn, &book.Title, &book.Subtitle, &book.Author, &book.Published, &book.Publisher, &book.Pages, &book.Description, &book.ImageUrl)
		if err != nil {
			return nil, fmt.Errorf(" -> Unable to query %w", err)
		}
		bookList = append(bookList, book)
	}
	if err != nil {
		if err == sql.ErrNoRows {
			return bookList, nil
		} else {
			return nil, fmt.Errorf(" -> Unable to query %w", err)
		}
	}
	return bookList, err
}

// AddBook adds a new book to the db. Should be only invoked by admin users.
func (lib *library) AddBook(book *models.Book) (id *string, e error) {
	sqlStatement := `INSERT INTO book(book_id, isbn, title, subtitle, author, published, publisher, pages, description, image_url)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING book_id`
	row, err := lib.db.Query(sqlStatement, book.BookId, book.Isbn, book.Title, book.Subtitle, book.Author, book.Published, book.Publisher,
		book.Pages, book.Description, book.ImageUrl)
	defer func(row *sql.Rows) {
		if row == nil {
			return
		}
		err := row.Close()
		if err != nil {
			e = err
		}
	}(row)
	if err != nil {
		return nil, fmt.Errorf("-> unable to add book %w", err)
	}
	var bookId string
	row.Next()
	err = row.Scan(&bookId)
	if err != nil {
		return nil, err
	}
	return &bookId, e
}

// DeleteBook deletes book from db based on the bookId
func (lib *library) DeleteBook(bookId string) (id *string, e error) {
	sqlStatement := `DELETE FROM book where book_id = $1 returning book_id`
	var deletedBookId string
	row, err := lib.db.Query(sqlStatement, bookId)
	defer func(row *sql.Rows) {
		if row == nil {
			return
		}
		err := row.Close()
		if err != nil {
			e = err
		}
	}(row)
	if err != nil {
		return nil, fmt.Errorf("-> unable to add book %w", err)
	}
	row.Next()
	err = row.Scan(&deletedBookId)
	if err != nil {
		return nil, err
	}
	return &deletedBookId, e
}

// AddUser adds an unique new user to the db. Fails if user try to
// use an already used email address.
func (lib *library) AddUser(user *models.User) (id *string, e error) {
	sqlStatement := `INSERT INTO reguser(user_id, first_name, last_name, gender, email, password, is_admin) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING user_id`
	row, err := lib.db.Query(sqlStatement, user.UserId, user.FirstName, user.LastName, user.Gender, user.Email, user.Password, user.IsAdmin)
	defer func(row *sql.Rows) {
		if row == nil {
			return
		}
		err := row.Close()
		if err != nil {
			e = err
		}
	}(row)
	if err != nil {
		return nil, fmt.Errorf("unable to add user %w", err)
	}
	var userId string
	row.Next()
	err = row.Scan(&userId)
	if err != nil {
		return nil, fmt.Errorf("unable to add user %w", err)
	}
	return &userId, e
}

// DeleteUser deletes an user from db. Only be invoked by the
// admin users.
func (lib *library) DeleteUser(userId string) (id *string, e error) {
	var deletedUserId string
	sqlStatement := `DELETE FROM reguser where user_id = $1 returning user_id`
	row, err := lib.db.Query(sqlStatement, userId)
	defer func(row *sql.Rows) {
		if row == nil {
			return
		}
		err := row.Close()
		if err != nil {
			deletedUserId = ""
			e = err
		}
	}(row)
	if err != nil {
		return nil, fmt.Errorf("-> unable delete user %w", err)
	}
	row.Next()
	err = row.Scan(&deletedUserId)
	if err != nil {
		return nil, err
	}
	return &deletedUserId, e
}

// SearchUserByEmail searches an user with a given emailId.
func (lib *library) SearchUserByEmail(email string) (resultUser *models.User, err error) {
	var user models.User
	sqlStatement := `SELECT * FROM reguser WHERE email = $1`
	err = lib.db.QueryRow(sqlStatement, email).Scan(&user.UserId, &user.FirstName, &user.LastName, &user.Gender, &user.Email, &user.Password, &user.IsAdmin)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

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

func (lib *library) AddReview(review *models.Review) (*string, error) {
	var reviewId string
	sqlStatement := `INSERT INTO review (review_id, user_id, book_id, comment, rating, date, edited) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING review_id`
	err := lib.db.QueryRow(sqlStatement, review.ReviewId, review.UserId, review.BookId, review.Comment, review.Rating, review.Date, review.Edited).Scan(&reviewId)
	if err != nil {
		return nil, err
	}
	return &reviewId, nil
}
