package config

import (
	"os"
)

type Config struct {
	DBType string
	DSN    string
}

// LoadConfig loads from env vars (or defaults)
func LoadConfig() *Config {
	dbType := os.Getenv("DB_TYPE")
	if dbType == "" {
		dbType = "memory" // default
	}

	dsn := os.Getenv("DSN")
	if dsn == "" {
		// default for MySQL
		dsn = "root:password@tcp(127.0.0.1:3306)/library"
	}

	return &Config{
		DBType: dbType,
		DSN:    dsn,
	}
}
