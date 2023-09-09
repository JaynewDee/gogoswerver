package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gogoswerver/postgres"
	"github.com/gogoswerver/web"
)

const uri string = "postgres://postgres:secret@localhost:5432/postgres?sslmode=disable"

func main() {

	// establish db connection and init store
	store, err := postgres.NewStore(uri)

	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Addr:    ":3001",
		Handler: web.NewHandler(store),
	}

	fmt.Println("Server listening @ 3001 ... ")
	server.ListenAndServe()
}
