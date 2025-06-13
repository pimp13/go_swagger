package config

import "time"

type DatabaseConfig struct {
	Host            string        `mapstructure:"DB_HOST"`
	Port            string        `mapstructure:"DB_PORT"`
	User            string        `mapstructure:"DB_USER"`
	Password        string        `mapstructure:"DB_PASSWORD"`
	Name            string        `mapstructure:"DB_NAME"`
	MaxIdleConns    int           `mapstructure:"DB_MAX_IDLE_CONNS"`
	MaxOpenConns    int           `mapstructure:"DB_MAX_OPEN_CONNS"`
	ConnMaxLifetime time.Duration `mapstructure:"DB_CONN_MAX_LIFETIME"`
	SSLMode         string        `mapstructure:"DB_SSL_MODE"`
	TimeZone        string        `mapstructure:"DB_TIMEZONE"`
}

var DBConfig = newDatabaseConfig()

func newDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:            getEnv("DB_HOST", "localhost"),
		Port:            getEnv("DB_PORT", "3306"),
		User:            getEnv("DB_USER", "root"),
		Password:        getEnv("DB_PASSWORD", ""),
		Name:            getEnv("DB_NAME", "myapp"),
		MaxIdleConns:    getEnvAsInt("DB_MAX_IDLE_CONNS", 10),
		MaxOpenConns:    getEnvAsInt("DB_MAX_OPEN_CONNS", 100),
		ConnMaxLifetime: getEnvAsDuration("DB_CONN_MAX_LIFETIME", time.Hour),
		SSLMode:         getEnv("DB_SSL_MODE", "disable"),
		TimeZone:        getEnv("DB_TIMEZONE", "UTC"),
	}
}
