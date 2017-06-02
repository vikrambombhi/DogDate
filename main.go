package main

import (
	"database/sql"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/vikrambombhi/DogDate/handlers"
)

func main() {
	dsn := "root:Vi20bo17@tcp(localhost:3306)/DogDate"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	handler := handlers.New(db)

	router := mux.NewRouter()
	router.HandleFunc("/", handler.GetAllDogs)
	router.HandleFunc("/login", handler.Login)

	server := http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	panic(server.ListenAndServe())
}
