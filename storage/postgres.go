package storage

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/omarshah0/go-clean-architecture/config"
	"github.com/omarshah0/go-clean-architecture/types"
)

type PostgresStore struct {
	db *gorm.DB
}

func NewPostgresStorage(config *config.Config) (*PostgresStore, error) {
	db, err := gorm.Open(postgres.Open(config.DatabaseUrl), &gorm.Config{TranslateError: true})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Perform the migrations
	err = db.AutoMigrate(&types.User{})

	if err != nil {
		return nil, fmt.Errorf("failed to perform migrations: %w", err)
	}

	return &PostgresStore{
		db: db,
	}, nil
}
