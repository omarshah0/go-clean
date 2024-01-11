package server

import (
	"github.com/gin-gonic/gin"

	"github.com/omarshah0/go-clean-architecture/middleware"
	"github.com/omarshah0/go-clean-architecture/storage"
)

func setupRoutes(rawRouter *gin.Engine, storage storage.Storage) error {
	router := rawRouter.Group("/api/v1")

	// Auth Routes
	router.GET("/auth", handleLoginRoute(storage))

	// Admin Routes
	driverRoutes := router.Group("/admin")
	driverRoutes.Use(middleware.AuthMiddleware("admin"))
	driverRoutes.Use(middleware.Logging())

	// Customer Routes
	customerRoutes := router.Group("/customer")
	customerRoutes.Use(middleware.AuthMiddleware("customer"))
	customerRoutes.Use(middleware.Logging())

	customerRoutes.GET("/", handleGetAllUsers(storage))
	customerRoutes.GET("/:id", handleGetUserById(storage))
	customerRoutes.POST("/", handleCreateUser(storage))

	// Not Found
	rawRouter.Use(middleware.Logging())
	rawRouter.NoRoute(handleNotFoundRoute)
	return nil
}
