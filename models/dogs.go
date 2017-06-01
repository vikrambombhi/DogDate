package models

import (
	"database/sql"
	"log"
)

type Dog struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Breed string `json:"breed`
}

func GetAllDogs(db *sql.DB) []Dog {
	dogs := []Dog{}
	rows, err := db.Query("select id, name, email, breed from Dogs")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var dog Dog
		err := rows.Scan(&dog.ID, &dog.Name, &dog.Email, &dog.Breed)
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
