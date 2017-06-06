package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/vikrambombhi/DogDate/models"
)

func (handler *Handler) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	claims := ctx.Value("user")
	user := models.GetUserByID(handler.DB, claims.(models.User).ID)
	dogs := models.GetDogsByUserID(handler.DB, claims.(models.User).ID)
	userInfo := map[string]interface{}{"user": user, "dog": dogs}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userInfo)
}

func (handler *Handler) UpdateDog(w http.ResponseWriter, r *http.Request) {
}
