package server

import (
	"github.com/gin-gonic/gin"

	handlers "github.com/omarshah0/go-clean-architecture/handler"
	"github.com/omarshah0/go-clean-architecture/middleware"
	"github.com/omarshah0/go-clean-architecture/storage"
	"github.com/omarshah0/go-clean-architecture/types"
)

func getTraceID(c *gin.Context) string {
	traceID, exists := c.Get("X-Trace-ID")
	if !exists {
		return "unknown"
	}

	traceIDStr, ok := traceID.(string)
	if !ok {
		return "unknown"
	}

	return traceIDStr
}

func sendErrorResponse(c *gin.Context, err *types.HandlerErrorResponse, statusCode int) {
	traceID := getTraceID(c)
	c.JSON(statusCode, &types.ErrorResponse{
		Type:    err.Type,
		Trace:   traceID,
		Message: err.Message,
		Error:   err.Error,
	})
}

func sendSuccessResponse(c *gin.Context, data interface{}, statusCode int) {
	c.JSON(statusCode, &types.SuccessResponse{
		Data: data,
	})
}

func setupRoutes(router *gin.Engine, storage storage.Storage) error {
	// Customer Routes
	customerRoutes := router.Group("/customer")
	customerRoutes.Use(middleware.AuthMiddleware("customer"))
	customerRoutes.Use(middleware.Logging())
	customerRoutes.GET("/ping", handlePingRoute(storage))

	// Driver Routes
	driverRoutes := router.Group("/driver")
	driverRoutes.Use(middleware.AuthMiddleware("driver"))
	driverRoutes.Use(middleware.Logging())
	driverRoutes.GET("/ping", handlePingRoute(storage))

	// Not Found
	router.Use(middleware.Logging())
	router.NoRoute(handleNotFoundRoute)
	return nil
}

func handlePingRoute(storage storage.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		response, err := handlers.HandleGetAllUsers(storage)

		if err != nil {
			sendErrorResponse(c, err, 400)
			return
		}

		sendSuccessResponse(c, response, 200)

	}
}

func handleNotFoundRoute(c *gin.Context) {
	sendErrorResponse(c, &types.HandlerErrorResponse{
		Type:    "NotFound",
		Message: "Not Found",
	}, 404)
}
