package db

import (
	"github.com/ducminhgd/intelligent-inventory/internal/infrastructure/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	cfg, err := config.Load(config.ConfigFile)
	if err != nil {
		return nil, err
	}
	dsn := cfg.Database.DSN
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
