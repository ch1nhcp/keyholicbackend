package main

import (
<<<<<<< HEAD
	"context"
=======
	"finalbackend/database"
	"finalbackend/routes"
>>>>>>> 28e18b7fe29baedf032232cbcd9cc754ca9c0425
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/olivere/elastic"
)

const (
	index    = "products"
	mappings = `
	{
		"settings":{
			"number_of_shards":2,
			"number_of_replicas":1
		},
		"mappings":{
			"properties":{
				"Name":{
					"type":"text"
				}

			}
		}
	}
	`
)

type Tweet struct {
	User    string `json:"user"`
	Message string `json:"message"`
}

<<<<<<< HEAD
func main() {
	client, err := elastic.NewClient(
		elastic.SetSniff(true),
		elastic.SetURL("http://localhost:9200"),
		elastic.SetHealthcheckInterval(5*time.Second), // quit trying after 5 seconds
=======
	database.Connect()
	r := mux.NewRouter()

	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"http://localhost:8080"}),
		handlers.AllowCredentials(),
>>>>>>> 28e18b7fe29baedf032232cbcd9cc754ca9c0425
	)
	// Check and see if olivere's NewClient() method returned an error
	if err != nil {
		// (Bad Request): Failed to parse content to map if mapping bad
		fmt.Println("elastic.NewClient() ERROR: %v", err)
		log.Fatalf("quiting connection..")
	} else {
		// Print client information
		fmt.Println("client:", client)
		fmt.Println("client TYPE:", reflect.TypeOf(client), "\n")
	}

	ctx := context.Background()

	exist, err := client.IndexExists(index).Do(ctx)
	if err != nil {
		log.Fatalf("IndexExists() ERROR: %v", err)

		// If the index exists..
	} else if exist {
		fmt.Println("The index " + index + " already exists.")
		_, err = client.DeleteIndex(index).Do(ctx)

<<<<<<< HEAD
		// Check for DeleteIndex() errors
		if err != nil {
			log.Fatalf("client.DeleteIndex() ERROR: %v", err)
		}
	}
=======
	routes.Setup(r)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
	log.Fatal(http.ListenAndServe(":8000", cors(r)))
	fmt.Println("back_end:http://localhost:8000")
>>>>>>> 28e18b7fe29baedf032232cbcd9cc754ca9c0425

	// Create a new index and pass the mappings string to the body
	create, err := client.CreateIndex(index).Body(mappings).Do(ctx)
	if err != nil {
		log.Fatalf("CreateIndex() ERROR: %v", err)
	} else {
		fmt.Println("CreateIndex():", create)
	}
	//add to twitch
	tweet := Tweet{User: "olivere", Message: "Take Five"}
	_, err = client.Index().
		Index("tweets").
		Type("doc").
		Id("1").
		BodyJson(tweet).
		Refresh("wait_for").
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	//search
	termQuery := elastic.NewTermQuery("user", "olivere")
	searchResult, err := client.Search().
		Index("tweets").            // search in index "tweets"
		Query(termQuery).           // specify the query
		Sort("user.keyword", true). // sort by "user" field, ascending
		From(0).Size(10).           // take documents 0-9
		Pretty(true).               // pretty print request and response JSON
		Do(context.Background())    // execute
	if err != nil {
		// Handle error
		fmt.Println("sai")
		panic(err)
	}
	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)
}
