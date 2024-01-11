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
