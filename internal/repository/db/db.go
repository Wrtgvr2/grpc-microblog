package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MustOpenDb(dsn string) *gorm.DB {
	db, err := OpenDb(dsn)
	if err != nil {
		panic(err)
	}
	return db
}

func OpenDb(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}
