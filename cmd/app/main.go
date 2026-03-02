package main

import (
	"Into/internal/config"
	"Into/internal/loggers"
)

func main() {
	logger := loggers.InitLogger()
	db := config.SetUpDatabaseConnection(logger)
	defer func() {
		if err := db.Close(); err != nil {
			logger.Error("Failed to close database", "error", err)
		}
	}()
}
