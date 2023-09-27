package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/jaynewdee/gogoswerver/postgres"
	"github.com/jaynewdee/gogoswerver/web"
)

const uri string = "postgres://postgres:secret@localhost:5432/postgres?sslmode=disable"

// 1. Establish postgres connection
// 2. Pass db store at server creation
// 3. Listen and serve
func main() {

	port := flag.String("port", ":3001", "server port")
	flag.Parse()

	store, err := postgres.NewStore(uri)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Database connection established.")
	}

	server := &http.Server{
		Addr:    *port,
		Handler: web.NewHandler(store),
	}

	fmt.Printf("Starting server @ %v\n", *port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("encountered error @ server start: %w", err)
	}

}
