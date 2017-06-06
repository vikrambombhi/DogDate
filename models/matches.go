package models

import (
	"database/sql"
	"log"
)

type Match struct {
	ID           int    `json:"id"`
	UserID       string `json:"userID"`
	OtherPartyID string `json:"otherPartyID"`
	Liked        bool   `json:"liked"`
}

func GetMatched(db *sql.DB, dogID int) []Dog {
	matches := []Dog{}
	rows, err := db.Query("select id, owner, name, breed, size from Dogs where id in (select Matches.otherDogID from Matches inner join Matches as AlsoLikesMe on Matches.dogID = AlsoLikesMe.otherDogID and Matches.otherDogID = AlsoLikesMe.dogID where Matches.liked=true and AlsoLikesMe.liked=true and Matches.dogID=?)", dogID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var dog Dog
		err := rows.Scan(&dog.ID, &dog.Owner, &dog.Name, &dog.Breed, &dog.Size)
		if err != nil {
			log.Fatal(err)
		}
		matches = append(matches, dog)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return matches
}

// Currently returns all dogs not including the users
func GetPotentialMatches(db *sql.DB, dogID int) []Dog {
	dogs := []Dog{}
	rows, err := db.Query("select id, owner, name, breed, size from Dogs where owner <> ?", dogID)
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

func GetLikedBy(db *sql.DB, dogID int) []Dog {
	dogs := []Dog{}
	rows, err := db.Query("select id, owner, name, breed, size from Dogs where id in (select dogID from Matches where otherDogID=?)", dogID)
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
