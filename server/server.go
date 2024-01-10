package server

import (
	"errors"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/omarshah0/go-clean-architecture/config"
	"github.com/omarshah0/go-clean-architecture/middleware"
	"github.com/omarshah0/go-clean-architecture/storage"
)

type GinServer struct {
	listenAddr string
	storage    storage.Storage
}

func NewGinServer(config *config.Config, storage storage.Storage) (*GinServer, error) {
	if config.Port == "" {
		return nil, errors.New("empty port configuration")
	}

	return &GinServer{
		listenAddr: config.Port,
		storage:    storage,
	}, nil
}

func (s *GinServer) Start() error {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"OPTIONS", "GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"*"}
	config.AllowCredentials = true

	router.Use(middleware.Tracing())

	if err := setupRoutes(router, s.storage); err != nil {
		return err
	}

	router.Run(s.listenAddr)
	return nil
}
