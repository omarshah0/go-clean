package storage

import (
	"github.com/omarshah0/go-clean-architecture/types"
)

type Storage interface {
	GetAllUsers() ([]*types.User, error)
}
