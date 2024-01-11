package config

import (
	"log"
	"os"
)

type Config struct {
	Port        string
	DatabaseUrl string
}

func NewConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	return &Config{
		Port:        ":" + port,
		DatabaseUrl: "postgres://postgres:Unclesnoopdog@69@localhost:5432/go_clean",
	}
}
