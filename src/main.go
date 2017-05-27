package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type key int

const (
	confKey key = iota
)

type config struct {
	address string
}

func logger(next http.Handler, conf config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Middleware ran")
		newContext := context.WithValue(r.Context(), confKey, conf)
		r = r.WithContext(newContext)
		next.ServeHTTP(w, r)
	})
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", hello)
	conf := config{address: ":8080"}
	logger := handlers.LoggingHandler(os.Stdout, logger(router, conf))
	server := http.Server{
		Addr:           ":8080",
		Handler:        logger,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	panic(server.ListenAndServe())
}
