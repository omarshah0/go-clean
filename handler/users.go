package handlers

import (
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

	return token, nil
}

func HandleGetAllUsers(s storage.Storage) ([]*types.User, *types.HandlerErrorResponse) {
	users, err := s.GetAll()

	if err != nil {
		errorResponse := &types.HandlerErrorResponse{
			Type:       err.Type,
			Message:    err.Message,
			StatusCode: err.StatusCode,
			Error:      err,
		}
		return nil, errorResponse
	}

	return users, nil
}

func HandleGetUserById(s storage.Storage) (*types.User, *types.HandlerErrorResponse) {
	user, err := s.GetById(1)

	if err != nil {
		errorResponse := &types.HandlerErrorResponse{
			Type:       err.Type,
			Message:    err.Message,
			StatusCode: err.StatusCode,
			Error:      err.Error,
		}
		return nil, errorResponse
	}

	return user, nil

}
