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

func GetDogsByUserID(db *sql.DB, userID int) []Dog {
	dogs := []Dog{}
	rows, err := db.Query("select id, owner, name, breed, size from Dogs where owner=?", userID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var dog Dog
		err := rows.Scan(&dog.Owner, &dog.ID, &dog.Name, &dog.Breed, &dog.Size)
		if err != nil {
			log.Fatal(err)
		}
		dogs = append(dogs, dog)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return dogs
}
