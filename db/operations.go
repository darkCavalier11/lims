package db

import (
	"database/sql"
	"fmt"
	"github.com/darkCavalier11/lims/models"
)

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

func (lib *library) DeleteBook(bookId string) (id *string, e error) {
	sqlStatement := `DELETE FROM book where book_id = $1 returning book_id`
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
	err = row.Scan(&bookId)
	if err != nil {
		return nil, err
	}
	return &bookId, e
}

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
