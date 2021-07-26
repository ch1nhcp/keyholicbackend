package main

import (
	"finalbackend/database"
	"finalbackend/rabitmq"
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
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"http://localhost:8080"}),
		handlers.AllowCredentials(),
	)
	routes.Setup(r)
	//kết nối rabitMQ
	rmq := rabitmq.RabbitMQ{
		ConnectionString: "amqp://tfs:tfs-ocg@174.138.40.239:5672/",
	}
	rmq.CreateConnection()
	defer rmq.Close()
	fmt.Println("Successfuly Connected To our RMQ Instance")
	ch := rmq.GetChannel()
	defer ch.Close()
	msgs := rabitmq.Consume(ch, "")
	fmt.Println(msgs)
	// go func() {
	// 	for data := range msgs {
	// 		controller.SendEmailBySendGrid(data.Body)
	// 	}
	// }()
	fmt.Println("http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", cors(r)))

}
