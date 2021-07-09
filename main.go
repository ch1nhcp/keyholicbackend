package main

import (
	"finalbackend/database"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	database.Connect()
	router := mux.NewRouter()
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
	// log.Fatal(http.ListenAndServe(":8080", route))

}
