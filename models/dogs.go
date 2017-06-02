package models

import (
	"database/sql"
	"log"
)

type Dog struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Owner int    `json:"owner"`
}

func GetAllDogs(db *sql.DB) []Dog {
	dogs := []Dog{}
	rows, err := db.Query("select id, name, breed, owner from Dogs")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var dog Dog
		err := rows.Scan(&dog.ID, &dog.Name, &dog.Breed, &dog.Owner)
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
