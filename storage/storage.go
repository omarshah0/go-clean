package storage

import (
	"github.com/omarshah0/go-clean-architecture/types"
)

type Storage interface {
	GetAll() ([]*types.User, *types.StorageErrorResponse)
	GetById(id int) (*types.User, *types.StorageErrorResponse)
	Create(user *types.User) (*types.User, *types.StorageErrorResponse)
}
