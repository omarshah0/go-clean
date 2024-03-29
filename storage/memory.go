package storage

import (
	"github.com/omarshah0/go-clean-architecture/config"
	"github.com/omarshah0/go-clean-architecture/types"
)

type MemoryStore struct{}

func NewMemoryStorage(config *config.Config) (*MemoryStore, error) {
	return &MemoryStore{}, nil
}

func (s *MemoryStore) GetAll() ([]*types.User, *types.HandlerErrorResponse) {
	return []*types.User{{Name: "Omar Farooq Shah", Email: "omar@gmail.com"}}, nil
}

func (s *MemoryStore) GetById(id int) (*types.User, *types.HandlerErrorResponse) {
	return nil, &types.HandlerErrorResponse{
		Type:       "NotFound",
		Message:    "Resource not found",
		StatusCode: 404,
		Error:      "Resource not found",
	}
}

func (s *MemoryStore) Create(user *types.User) (*types.User, *types.HandlerErrorResponse) {
	return user, nil
}
