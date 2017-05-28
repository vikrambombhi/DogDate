package models

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Setup(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	return db, err
}

func GetUsers(db *sql.DB) {
	var (
		id    int
		email string
	)
	rows, err := db.Query("select id, email from Users where id = ?", 2)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &email)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, email)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
