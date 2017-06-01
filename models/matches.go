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

func findMatches(db *sql.DB, userid int) []Match {
	matches := []Match{}
	rows, err := db.Query("select id, userID, otherPartyID, liked from Matches where otherPartyID where otherPartyID = ?", userid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var match Match
		err := rows.Scan(&match.ID, &match.UserID, &match.OtherPartyID, &match.Liked)
		if err != nil {
			log.Fatal(err)
		}
		matches = append(matches, match)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return matches
}
