package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/vikrambombhi/DogDate/models"
)

func (handler *Handler) GetPotentialMatches(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	claims := ctx.Value("user")
	singles := models.GetPotentialMatches(handler.DB, claims.(models.User).ID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(singles)
}

func (handler *Handler) GetMatched(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	claims := ctx.Value("user")
	matches := models.GetMatched(handler.DB, claims.(models.User).ID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(matches)
}

func (handler *Handler) GetLikedBy(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	claims := ctx.Value("user")
	interestedUsers := models.GetLikedBy(handler.DB, claims.(models.User).ID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(interestedUsers)
}
