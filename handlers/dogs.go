package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/vikrambombhi/DogDate/models"
)

func (handler *Handler) GetAllDogs(w http.ResponseWriter, r *http.Request) {
	dogs := models.GetAllDogs(handler.DB)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dogs)
}
