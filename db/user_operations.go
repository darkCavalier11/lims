package db

import (
	"database/sql"
	"fmt"
	"github.com/darkCavalier11/lims/models"
)

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
