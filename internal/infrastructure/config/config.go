package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

const (
	ConfigFile = "config.yaml"
)

type DatabaseConfig struct {
	DSN string `mapstructure:"dsn"`
}

// Config is the top-level application configuration.
type Config struct {
	Database DatabaseConfig `mapstructure:"database"`
}

func Load(configFile string) (*Config, error) {
	v := viper.NewWithOptions(viper.KeyDelimiter("."))
	v.SetEnvPrefix("IID")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if configFile != "" {
		v.SetConfigFile(configFile)
		if err := v.ReadInConfig(); err != nil {
			return nil, fmt.Errorf("read config file %q: %w", configFile, err)
		}
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}

	return &cfg, nil
}

func LoadAndValidate(configFile string) (*Config, error) {
	cfg, err := Load(configFile)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
