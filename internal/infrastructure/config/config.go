package config

import (
	"context"
	"fmt"

	"github.com/sethvargo/go-envconfig"
)

type DatabaseConfig struct {
	DSN         string `mapstructure:"dsn" env:"DSN"`
	AutoMigrate bool   `mapstructure:"auto_migrate" env:"AUTO_MIGRATE, default=false"`
}

type RESTConfig struct {
	Host string `mapstructure:"host" env:"HOST, default=0.0.0.0"`
	Port int    `mapstructure:"port" env:"PORT, default=8080"`
}

// Config is the top-level application configuration.
type Config struct {
	REST RESTConfig `mapstructure:"rest" env:",prefix=REST_"`

	Database DatabaseConfig `mapstructure:"database" env:",prefix=DATABASE_"`

	LogLevel string `mapstructure:"log_level" env:"LOG_LEVEL"`
}

func Load() (*Config, error) {
	var cfg Config
	ctx := context.Background()
	if err := envconfig.Process(ctx, &cfg); err != nil {
		return nil, fmt.Errorf("cannot load config: %w", err)
	}

	return &cfg, nil
}
