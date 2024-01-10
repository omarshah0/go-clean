package handlers

import (
	"github.com/omarshah0/go-clean-architecture/storage"
	"github.com/omarshah0/go-clean-architecture/types"
)

func HandleGetAllUsers(s storage.Storage) (string, *types.HandlerErrorResponse) {
	users, err := s.GetAllUsers()

	if err != nil {
		errorResponse := &types.HandlerErrorResponse{
			Type:    "Validation",
			Message: "Validation failed. Please resubmit your request.",
			Error:   "This is Error From User Handler",
		}
		return "", errorResponse
	}

	return users, nil
}
