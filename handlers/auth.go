package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/vikrambombhi/DogDate/models"
)

type Claims struct {
	Email              string `json:"email"`
	UserID             int    `json:"userID"`
	dogID              int    `json:"dogID"`
	jwt.StandardClaims `json:"StandardClaims"`
}

var secret = []byte(os.Getenv("SECRET"))

func generateToken(user models.User, dog models.Dog) *jwt.Token {
	claims := Claims{
		user.Email,
		user.ID,
		dog.ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "DogDate",
		},
	}
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

	// Find users dog (to make things easyier later)
	dog := models.GetDogByOwner(handler.DB, user.ID)

	// Generate token
	token := generateToken(user, dog)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		log.Fatal(err)
	}

	// Send JWT
	w.Header().Set("WWW-Authenticatr", `Basic realm="Restricted"`)
	json.NewEncoder(w).Encode(tokenString)
}

func validateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("Token not valid")
	}
	return claims, nil
}

func (handler *Handler) GetUser(r *http.Request) (models.User, error) {
	auth := r.Header.Get("Authorization")
	var user models.User
	if auth == "" {
		return user, fmt.Errorf("Authorization header is empty")
	}
	token := strings.TrimPrefix(auth, "Bearer ")
	claims, err := validateToken(token)
	if err != nil {
		log.Print(err)
		return user, err
	}
	user = models.GetUserByEmail(handler.DB, claims.Email)
	return user, nil
}
