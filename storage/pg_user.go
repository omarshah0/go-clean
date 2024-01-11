package storage

import (
	"errors"

	"github.com/omarshah0/go-clean-architecture/types"
	"gorm.io/gorm"
)

func (s *PostgresStore) GetAll() ([]*types.User, *types.StorageErrorResponse) {
	var users []*types.User

	if err := s.db.Find(&users).Error; err != nil {
		return nil, &types.StorageErrorResponse{
			Type:       "InternalError",
			Message:    "Internal Server Error",
			Error:      err.Error(),
			StatusCode: 500,
		}
	}

	return users, nil
}

func (s *PostgresStore) GetById(id int) (*types.User, *types.StorageErrorResponse) {
	user := new(types.User)

	if err := s.db.First(&user, id).Error; err != nil {
		return nil, &types.StorageErrorResponse{
			Type:       "NotFound",
			Message:    err.Error(),
			StatusCode: 404,
			Error:      err.Error(),
		}
	}

	return user, nil
}

func (s *PostgresStore) Create(user *types.User) (*types.User, *types.StorageErrorResponse) {
	result := s.db.Create(&user)

	err := new(types.StorageErrorResponse)

	if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
		err.Type = "DuplicateError"
		err.Message = "Resource already exists"
		err.StatusCode = 409
		err.Error = result.Error.Error()

		return nil, err
	}

	return user, nil
}
