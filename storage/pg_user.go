package storage

import (
	"errors"

	"github.com/omarshah0/go-clean-architecture/types"
	"gorm.io/gorm"
)

func (s *PostgresStore) GetAll() ([]*types.User, *types.HandlerErrorResponse) {
	var users []*types.User

	if err := s.db.Find(&users).Error; err != nil {
		return nil, &types.HandlerErrorResponse{
			Type:       "InternalError",
			Message:    "Internal Server Error",
			Error:      err.Error(),
			StatusCode: 500,
		}
	}

	return users, nil
}

func (s *PostgresStore) GetById(id int) (*types.User, *types.HandlerErrorResponse) {
	user := new(types.User)
	err := new(types.HandlerErrorResponse)

	result := s.db.First(&user, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		err.Type = "NotFound"
		err.Message = "Resource not found"
		err.StatusCode = 404
		err.Error = result.Error.Error()

		return nil, err
	}

	return user, nil
}

func (s *PostgresStore) Create(user *types.User) (*types.User, *types.HandlerErrorResponse) {
	err := new(types.HandlerErrorResponse)

	result := s.db.Create(&user)

	if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
		err.Type = "DuplicateError"
		err.Message = "Resource already exists"
		err.StatusCode = 409
		err.Error = result.Error.Error()

		return nil, err
	}

	return user, nil
}
