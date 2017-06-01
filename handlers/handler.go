package handlers

import "database/sql"

type Handler struct {
	DB *sql.DB
}

func New(db *sql.DB) *Handler {
	return &Handler{db}
}
