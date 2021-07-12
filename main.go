package main

import (
	"finalbackend/database"
	"finalbackend/routes"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	database.Connect()
	r := mux.NewRouter()
	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}),
	)

	routes.Setup(r)
	log.Fatal(http.ListenAndServe(":8001", cors(r)))

}
