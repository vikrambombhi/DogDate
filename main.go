package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/vikrambombhi/DogDate/models"
)

type Env struct {
	db *sql.DB
}

var env Env

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Middleware ran")
		next.ServeHTTP(w, r)
	})
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
	models.GetUsers(env.db)
}

func main() {
	var err error
	env.db, err = models.Setup("root:Vi20bo17@tcp(localhost:3306)/DogDate")
	if err != nil {
		panic(err)
	}
	defer env.db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/", hello)
	middleware := handlers.LoggingHandler(os.Stdout, middleware(router))

	server := http.Server{
		Addr:           ":8080",
		Handler:        middleware,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	panic(server.ListenAndServe())
}
