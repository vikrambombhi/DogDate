package models

import (
	"database/sql"
	"log"
)

type Dog struct {
	ID    int    `json:"id"`
	Owner int    `json:"owner"`
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Size  string `json:"size"`
}

func GetDogByOwner(db *sql.DB, userID int) Dog {
	dog := Dog{}
	err := db.QueryRow("select id, owner, name, breed, size from Dogs where owner=?", userID).Scan(&dog.Owner, &dog.ID, &dog.Name, &dog.Breed, &dog.Size)
	if err != nil {
		log.Fatal(err)
	}
	return dog
}
