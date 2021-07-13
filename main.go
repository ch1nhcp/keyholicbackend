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
		handlers.AllowedOrigins([]string{"http://localhost:8080"}),
		handlers.AllowCredentials(),
	)

	routes.Setup(r)
	log.Fatal(http.ListenAndServe(":8000", cors(r)))

}
