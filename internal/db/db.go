package db

import (
	"log"

	"goUniAdmin/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB holds the GORM database connection
type DB struct {
	*gorm.DB
}

// NewDB initializes a new GORM database connection
func NewDB(cfg *config.Config) (*DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.DATABASE_URL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Test the connection
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	log.Println("Connected to PostgreSQL database with URL: ", cfg.DATABASE_URL)
	return &DB{db}, nil
}
