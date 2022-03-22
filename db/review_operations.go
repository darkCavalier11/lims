package db

import (
	"database/sql"
	"fmt"
	"github.com/darkCavalier11/lims/models"
)

func (lib *library) AddReview(review *models.Review) (*string, error) {
	var reviewId string
	sqlStatement := `INSERT INTO review (review_id, user_id, book_id, comment, rating, date, edited) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING review_id`
	err := lib.db.QueryRow(sqlStatement, review.ReviewId, review.UserId, review.BookId, review.Comment, review.Rating, review.Date, review.Edited).Scan(&reviewId)
	if err != nil {
		return nil, err
	}
	return &reviewId, nil
}

func (lib *library) DeleteReview(reviewId string) (*string, error) {
	var deletedReviewId string
	sqlStatement := `DELETE FROM review WHERE review_id = $1 RETURNING review_id`
	err := lib.db.QueryRow(sqlStatement, reviewId).Scan(&deletedReviewId)
	if err != nil {
		return nil, err
	}
	return &deletedReviewId, nil
}

func (lib *library) EditReview(review *models.Review) (*string, error) {
	var reviewId string
	sqlStatement := `UPDATE review SET comment = $1, rating = $2, date = $3, edited =$4 WHERE review_id = $5 returning review_id`
	err := lib.db.QueryRow(sqlStatement, review.Comment, review.Rating, review.Date, review.Edited, review.ReviewId).Scan(&reviewId)
	if err != nil {
		return nil, err
	}
	return &reviewId, nil
}

func (lib *library) GetReviewByReviewId(reviewId string) (*models.Review, error) {
	var review models.Review
	sqlStatement := `SELECT * FROM review WHERE review_id = $1`
	err := lib.db.QueryRow(sqlStatement, reviewId).Scan(&review.ReviewId, &review.UserId, &review.BookId, &review.Comment, &review.Rating, &review.Date, &review.Edited)
	if err != nil {
		return nil, err
	}
	return &review, nil
}

func (lib *library) GetReviewsOfBook(bookId string) ([]*models.Review, error) {
	var reviews []*models.Review
	sqlStatement := `SELECT * FROM review WHERE book_id = $1`
	rows, err := lib.db.Query(sqlStatement, bookId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var review models.Review
		err := rows.Scan(&review.ReviewId, &review.UserId, &review.BookId, &review.Comment, &review.Rating, &review.Date, &review.Edited)
		if err != nil {
			if err == sql.ErrNoRows {
				return reviews, nil
			}
			if err != nil {
				return nil, fmt.Errorf(" -> Unable to query %w", err)
			}
		}
		reviews = append(reviews, &review)
	}
	return reviews, nil
}
