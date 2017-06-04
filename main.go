package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/vikrambombhi/DogDate/handlers"
)

var handler *handlers.Handler

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.String() == "/login" {
			next.ServeHTTP(w, r)
			return
		}

		user, err := handler.GetUser(r)
		if err != nil {
			log.Print("error occoured")
			http.Error(w, "Token not valid", http.StatusForbidden)
			return
		}
		log.Print(user)

		next.ServeHTTP(w, r)
	})
}

func main() {
	dsn := "root:Vi20bo17@tcp(localhost:3306)/DogDate"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	handler = handlers.New(db)

	router := mux.NewRouter()
	router.HandleFunc("/", handler.GetAllDogs)
	router.HandleFunc("/login", handler.Login)

	server := http.Server{
		Addr:           ":8080",
		Handler:        middleware(router),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	panic(server.ListenAndServe())
}
