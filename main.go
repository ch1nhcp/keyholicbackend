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
		handlers.AllowCredentials(),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"Content-Type", "jwt", "Set-Cookie"}),
	)
	r.Use(cors)

	routes.Setup(r)
	// log.Fatal(srv.ListenAndServe())
	log.Fatal(http.ListenAndServe(":8000", r))

}
