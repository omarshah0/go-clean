package main

import (
	"github.com/omarshah0/go-clean-architecture/config"
	"github.com/omarshah0/go-clean-architecture/server"
	"github.com/omarshah0/go-clean-architecture/storage"
)

func main() {
	config := config.NewConfig()

	// Init Storage
	store, err := storage.NewPostgresStorage(config)

	if err != nil {
		panic(err)
	}

	// Init Gin Server
	server, err := server.NewGinServer(config, store)

	if err != nil {
		panic(err)
	}

	server.Start()
}
