package config

import (
	"os"
)

type Config struct {
	AppEnv     string
	Port       string
	DBDriver   string
	DBDSN      string
	JWTSecret  string
}

func LoadConfig() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	dbDriver := os.Getenv("DB_DRIVER")
	if dbDriver == "" {
		dbDriver = "sqlite"
	}

	dbDSN := os.Getenv("DB_DSN")
	if dbDSN == "" {
		dbDSN = "labelin.db"
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "super-secret-key-change-in-prod"
	}

	return Config{
		AppEnv:    os.Getenv("APP_ENV"),
		Port:      port,
		DBDriver:  dbDriver,
		DBDSN:     dbDSN,
		JWTSecret: jwtSecret,
	}
}
