package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(cfg Config) {
	var err error
	
	if cfg.DBDriver == "postgres" {
		DB, err = gorm.Open(postgres.Open(cfg.DBDSN), &gorm.Config{})
	} else {
		// Default to SQLite
		DB, err = gorm.Open(sqlite.Open(cfg.DBDSN), &gorm.Config{})
	}

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connection established")
}
