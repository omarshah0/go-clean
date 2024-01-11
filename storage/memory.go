package storage

import (
	"github.com/omarshah0/go-clean-architecture/config"
	"github.com/omarshah0/go-clean-architecture/types"
)

type MemoryStore struct{}

func NewMemoryStorage(config *config.Config) (*MemoryStore, error) {
	return &MemoryStore{}, nil
}

func (s *MemoryStore) GetAllUsers() ([]*types.User, error) {
	return []*types.User{{Name: "Omar Farooq Shah", Email: "omar@gmail.com"}}, nil
}
