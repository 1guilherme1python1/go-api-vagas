package config

import (
	"github.com/1guilherme1python1/go-api-vagas/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

func InitializeSqlite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	dbPath := "./db/main.db"

	//Check if the database file exists
	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("database file doesn't exist, creating...")

		//create the database file and directory
		err := os.MkdirAll("./db", os.ModePerm)

		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}

		err = file.Close()
		if err != nil {
			return nil, err
		}
	}

	open, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Error("sqlite open failed", err)
		return nil, err
	}

	err = open.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Error("sqlite migrate failed", err)
	}
	return open, nil
}
