package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	APIKey   string         `mapstructure:"api_key"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

type DatabaseConfig struct {
	Path string `mapstructure:"path"`
}

func Load() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// Set default values
	viper.SetDefault("server.port", "3000")
	viper.SetDefault("database.path", "carsearch.db")
	viper.SetDefault("api_key", "default-api-key")

	// Allow environment variables without prefix
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// Read config file if available
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("No config file found, using defaults and environment variables: %v", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Unable to decode config: %v", err)
	}

	return &cfg
}
