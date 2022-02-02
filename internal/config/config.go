package config

import (
	"os"
)

type Config struct {
	Environment string
	LogLevel    string
	Context     struct {
		Timeout string
	}
	Server struct {
		Protocol string
		Host     string
		Port     string
	}
	DB struct {
		Host     string
		Port     string
		Name     string
		User     string
		Password string
		Sslmode  string
	}
}

func New() *Config {
	var config Config

	config.Environment = getEnv("ENVIRONMENT", "develop")
	config.LogLevel = getEnv("LOG_LEVEL", "debug")
	config.Context.Timeout = getEnv("CONTEXT_TIMEOUT", "30s")

	// initialization server
	config.Server.Protocol = getEnv("SERVER_PROTOCOL", "http")
	config.Server.Host = getEnv("SERVER_HOST", "localhost")
	config.Server.Port = getEnv("SERVER_PORT", ":9000")

	// initialization db
	config.DB.Host = getEnv("DATABASE_HOST", "localhost")
	config.DB.Port = getEnv("DATABASE_PORT", "5432")
	config.DB.Name = getEnv("DATABASE_NAME", "postgres")
	config.DB.User = getEnv("DATABASE_USER", "postgres")
	config.DB.Password = getEnv("DATABASE_PASSWORD", "")
	config.DB.Sslmode = getEnv("DATABASE_SSLMODE", "false")

	return &config
}

func getEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultValue
}
