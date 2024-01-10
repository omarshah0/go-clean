package storage

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/omarshah0/go-clean-architecture/config"
)

type PostgresStore struct {
	db *gorm.DB
}

func NewPostgresStorage(config *config.Config) (*PostgresStore, error) {
	db, err := gorm.Open(postgres.Open(config.DatabaseUrl), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) GetAllUsers() (string, error) {
	return "Omar Farooq Shah", nil
}
