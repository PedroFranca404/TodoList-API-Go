package config

import (
	"os"

	"github.com/PedroFranca404/TodoList-API-Go/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitliazieSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	// Check if DB exist
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")
		// Create DB
		err := os.Mkdir("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	// Create DB and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate schema
	err = db.AutoMigrate(&schemas.Todo{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
