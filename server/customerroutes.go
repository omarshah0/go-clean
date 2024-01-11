package server

import (
	"strconv"

	"github.com/gin-gonic/gin"

	handlers "github.com/omarshah0/go-clean-architecture/handler"
	"github.com/omarshah0/go-clean-architecture/storage"
	"github.com/omarshah0/go-clean-architecture/types"
)

func handleGetAllUsers(storage storage.Storage) gin.HandlerFunc {
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
		idStr := c.Param("id")
		id, convErr := strconv.Atoi(idStr)

		if convErr != nil {
			errMessage := &types.HandlerErrorResponse{
				Type:       "BadRequest",
				Message:    "Invalid ID",
				StatusCode: 400,
				Error:      convErr.Error(),
			}
			sendErrorResponse(c, errMessage)
			return
		}

		response, err := handlers.HandleGetUserById(id, storage)

		if err != nil {
			sendErrorResponse(c, err)
			return
		}
		sendSuccessResponse(c, response, 200)
	}
}
