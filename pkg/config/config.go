package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	*DatabaseConfig
}

var Envs = NewConfig()

func NewConfig() *Config {
	return &Config{
		DatabaseConfig: DBConfig,
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.Atoi(value)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}

func getEnvAsDuration(key string, fallback time.Duration) time.Duration {
	if value, ok := os.LookupEnv(key); ok {
		d, err := time.ParseDuration(value)
		if err != nil {
			return fallback
		}
		return d
	}
	return fallback
}
