package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vikrambombhi/DogDate/models"
)

func (handler *Handler) GetAccountInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	claims := ctx.Value("user")
	urlParams := mux.Vars(r)
	userID, err := strconv.Atoi(urlParams["userID"])
	if err != nil {
		http.Error(w, "id not valid", http.StatusNotAcceptable)
	}

	user := models.GetUserByID(handler.DB, userID)
	dogs := models.GetDogsByUserID(handler.DB, userID)
	if userID == claims.(models.User).ID {
		userInfo := map[string]interface{}{"user": user, "dog": dogs}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(userInfo)
		return
	} else {
		userInfo := map[string]interface{}{"userName": user.Name, "dog": dogs}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(userInfo)
		return
	}
}
