package db

import (
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

func CloseDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
