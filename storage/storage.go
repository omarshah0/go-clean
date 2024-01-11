package storage

import (
	"github.com/omarshah0/go-clean-architecture/types"
)

type Storage interface {
	GetAll() ([]*types.User, *types.HandlerErrorResponse)
	GetById(id int) (*types.User, *types.HandlerErrorResponse)
	Create(user *types.User) (*types.User, *types.HandlerErrorResponse)
}
