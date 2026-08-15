package db

import (
	"github.com/ducminhgd/intelligent-inventory/internal/adapter/postgresql"
	"github.com/ducminhgd/intelligent-inventory/internal/infrastructure/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg config.DatabaseConfig) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// AutoMigrate automatically migrates the database schema for the given models.
// This function should be called during application startup to ensure that the database schema is up-to-date.
//
//	IMPORTANT: Use this for development and testing purposes only. In production, consider using a proper migration tool to manage schema changes.
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&postgresql.ManufacturerModel{},
	)
}

func CloseDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
