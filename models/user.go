package models

import (
	"database/sql"
	"log"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Name     string `json:"name"`
}

func GetUser(db *sql.DB, email string, password string) User {
	var user User
	err := db.QueryRow("select id, email, password, name from Users where email=? and password=?", email, password).Scan(&user.ID, &user.Email, &user.Password, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return user
		} else {
			log.Fatal(err)
		}
	}
	return user
}

func GetUserByID(db *sql.DB, userID int) User {
	var user User
	err := db.QueryRow("select id, email, password, name from Users where id=?", userID).Scan(&user.ID, &user.Email, &user.Password, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return user
		} else {
			log.Fatal(err)
		}
	}
	return user
}

func GetUserByEmail(db *sql.DB, email string) User {
	var user User
	err := db.QueryRow("select id, email, password, name from Users where email=?", email).Scan(&user.ID, &user.Email, &user.Password, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return user
		} else {
			log.Fatal(err)
		}
	}
	return user
}
