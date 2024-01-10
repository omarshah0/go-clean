package handlers

import (
	"fmt"

	"github.com/omarshah0/go-clean-architecture/middleware"
	"github.com/omarshah0/go-clean-architecture/storage"
	"github.com/omarshah0/go-clean-architecture/types"
)

func HandleLoginUser(s storage.Storage) (string, *types.HandlerErrorResponse) {
	user := types.User{Name: "Omar Farooq Shah", Email: "oemyoem55@gmail.com", Type: types.Admin}

	token, err := middleware.GenerateToken(&user)

	if err != nil {
		errorResponse := &types.HandlerErrorResponse{
			Type:       "InternalError",
			Message:    "Internal Server Error",
			Error:      err.Error(),
			StatusCode: 500,
		}
		return "", errorResponse
	}

	fmt.Println("Token ", token)

	return token, nil
}

func HandleGetAllUsers(s storage.Storage) (string, *types.HandlerErrorResponse) {
	users, err := s.GetAllUsers()

	if err != nil {
		errorResponse := &types.HandlerErrorResponse{
			Type:       "Validation",
			Message:    "Validation failed. Please resubmit your request.",
			Error:      "This is Error From User Handler",
			StatusCode: 400,
		}
		return "", errorResponse
	}

	return users, nil
}
