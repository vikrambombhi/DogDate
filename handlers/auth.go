package handlers

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/vikrambombhi/DogDate/models"
)

var secret = []byte(os.Getenv("SECRET"))

func generateToken(claims jwt.MapClaims) *jwt.Token {
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
}

func (handler *Handler) Login(w http.ResponseWriter, r *http.Request) {
	// Get email and password send by baiuc auth
	s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(s) != 2 {
		http.Error(w, "Not authorized", 401)
		return
	}
	b, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 {
		http.Error(w, "username or password missing", 401)
		return
	}

	// Check if email and passowrd is valid
	// pair[0] is email
	// pair[1] is password
	user := models.GetUser(handler.DB, pair[0], pair[1])
	if user.Email == "" {
		http.Error(w, "Invaild username or password", 401)
		return
	}

	// Create JWT token to send back
	claims := jwt.MapClaims{
		"email": user.Email,
		"name":  user.Name,
	}
	token := generateToken(claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		log.Fatal(err)
	}

	// Send JWT
	w.Header().Set("WWW-Authenticatr", `Basic realm="Restricted"`)
	json.NewEncoder(w).Encode(tokenString)
}
