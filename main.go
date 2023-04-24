package main

import (
	"log"

	"github.com/remy-z/gobank/api"
	"github.com/remy-z/gobank/storage"
)

func main() {
	store, err := storage.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}
	server := api.NewAPIServer(":3000", store)
	server.Run()
}
