package main

import (
	"finalbackend/database"
	"finalbackend/routes"
	"fmt"
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
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
	log.Fatal(http.ListenAndServe(":8000", cors(r)))
<<<<<<< HEAD
	fmt.Println("back_end:http://localhost:8000")
=======
>>>>>>> 44bcd0f744366aff30f1af43193fd62cce234479

}
