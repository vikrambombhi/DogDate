package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Setup(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	return db, err
}
