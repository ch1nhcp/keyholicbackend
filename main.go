package main

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/olivere/elastic"
)




	database.Connect()
	r := mux.NewRouter()

	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"http://localhost:8080"}),
		handlers.AllowCredentials(),
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

	routes.Setup(r)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
	log.Fatal(http.ListenAndServe(":8000", cors(r)))
	fmt.Println("back_end:http://localhost:8000")

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
