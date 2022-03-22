package db

import (
	"database/sql"
	"fmt"
	"github.com/darkCavalier11/lims/models"
)

func (lib *library) GetBookById(bookId string) (bookList *models.Book, err error) {
	var book models.Book
	sqlStatement := `SELECT * FROM book WHERE book_id = $1`
	err = lib.db.QueryRow(sqlStatement, bookId).Scan(&book.BookId, &book.Isbn, &book.Title, &book.Subtitle, &book.Author, &book.Published, &book.Publisher, &book.Pages, &book.Description, &book.ImageUrl, &book.Reviews, &book.Rating)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

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
		err := rows.Scan(&book.BookId, &book.Isbn, &book.Title, &book.Subtitle, &book.Author, &book.Published, &book.Publisher, &book.Pages, &book.Description, &book.ImageUrl, &book.Reviews, &book.Rating)
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
	sqlStatement := `INSERT INTO book(book_id, isbn, title, subtitle, author, published, publisher, pages, description, image_url, reviews, rating)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING book_id`
	row, err := lib.db.Query(sqlStatement, book.BookId, book.Isbn, book.Title, book.Subtitle, book.Author, book.Published, book.Publisher,
		book.Pages, book.Description, book.ImageUrl, book.Reviews, book.Rating)
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

func (lib *library) GetBooksIssuedIdByUser(userId string) ([]*string, error) {
	var bookList []*string
	sqlStatement := `SELECT book_id FROM bookissue WHERE user_id = $1`
	rows, err := lib.db.Query(sqlStatement, userId)
	if err != nil {
		return nil, fmt.Errorf("unable to find books issued by user %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var bookId string
		err := rows.Scan(&bookId)
		if err != nil {
			return nil, fmt.Errorf(" -> Unable to query %w", err)
		}
		bookList = append(bookList, &bookId)
	}
	return bookList, nil
}
