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

func sendErrorResponse(c *gin.Context, err *types.HandlerErrorResponse) {
	traceID := getTraceID(c)
	c.JSON(err.StatusCode, &types.ErrorResponse{
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

func setupRoutes(rawRouter *gin.Engine, storage storage.Storage) error {
	router := rawRouter.Group("/api/v1")

	// Auth Routes
	router.GET("/auth", handleLoginRoute(storage))

	// Admin Routes
	driverRoutes := router.Group("/admin")
	driverRoutes.Use(middleware.AuthMiddleware("admin"))
	driverRoutes.Use(middleware.Logging())
	driverRoutes.GET("/", handleUserRoutes(storage))

	// Customer Routes
	customerRoutes := router.Group("/customer")
	customerRoutes.Use(middleware.AuthMiddleware("customer"))
	customerRoutes.Use(middleware.Logging())
	customerRoutes.GET("/", handleUserRoutes(storage))
	customerRoutes.POST("/", handleCreateUser(storage))
	customerRoutes.GET("/:id", handleGetUserById(storage))

	// Not Found
	rawRouter.Use(middleware.Logging())
	rawRouter.NoRoute(handleNotFoundRoute)
	return nil
}

// Auth Routes

func handleLoginRoute(storage storage.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		response, err := handlers.HandleLoginUser(storage)
		if err != nil {
			sendErrorResponse(c, err)
			return
		}
		sendSuccessResponse(c, response, 200)
	}
}

// Customer Routes

func handleUserRoutes(storage storage.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		response, err := handlers.HandleGetAllUsers(storage)
		if err != nil {
			sendErrorResponse(c, err)
			return
		}
		sendSuccessResponse(c, response, 200)
	}
}

func handleCreateUser(storage storage.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := new(types.User)

		if err := c.ShouldBindJSON(&user); err != nil {
			errMessage := &types.HandlerErrorResponse{
				Type:       "BadRequest",
				Message:    "Invalid request body",
				StatusCode: 400,
				Error:      err.Error(),
			}
			sendErrorResponse(c, errMessage)
			return
		}

		response, err := handlers.HandleCreateUser(user, storage)
		if err != nil {
			sendErrorResponse(c, err)
			return
		}
		sendSuccessResponse(c, response, 200)
	}
}

func handleGetUserById(storage storage.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		response, err := handlers.HandleGetUserById(storage)
		if err != nil {
			sendErrorResponse(c, err)
			return
		}
		sendSuccessResponse(c, response, 200)
	}
}

// Not Found Routes

func handleNotFoundRoute(c *gin.Context) {
	sendErrorResponse(c, &types.HandlerErrorResponse{
		Type:       "NotFound",
		Message:    "Not Found",
		Error:      "Endpoint not defined",
		StatusCode: 404,
	})
}
