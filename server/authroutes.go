package server

import (
	"github.com/gin-gonic/gin"

	handlers "github.com/omarshah0/go-clean-architecture/handler"
	"github.com/omarshah0/go-clean-architecture/storage"
)

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
