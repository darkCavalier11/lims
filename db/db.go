package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "limsdb"
	password = "password123"
	dbname   = "lims"
)

type library struct {
	db *sql.DB
}

var Lib *library

// Connect ensures the connection to right database.
// It creates a db struct Lib internally that can be accessed
// by other module that can use listed methods only.
// Must be invoked in main to be used in other modules.
func Connect() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return fmt.Errorf(" -> Unable to open db %w", err)
	}
	db.SetConnMaxIdleTime(time.Duration(time.Second * 5))
	Lib = &library{
		db,
	}
	err = Lib.db.Ping()
	if err != nil {
		return fmt.Errorf(" -> Unable to ping db %w", err)
	}
	return nil
}
