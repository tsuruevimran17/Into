package main

import (
	"fmt"
	"log/slog"
	"net/url"
	"os"

	"github.com/joho/godotenv"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	if err := godotenv.Load(); err != nil {
		logger.Error("Error loading .env file", "error", err)
		os.Exit(1)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbSSL := os.Getenv("DB_SSLMODE")
	if dbSSL == "" {
		dbSSL = "disable"
	}

	u := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(dbUser, dbPass),
		Host:   fmt.Sprintf("%s:%s", dbHost, dbPort),
		Path:   "/" + dbName,
	}
	q := u.Query()
	q.Set("sslmode", dbSSL)
	u.RawQuery = q.Encode()
	dbURL := u.String()

	m, err := migrate.New("file://migrations", dbURL)
	if err != nil {
		logger.Error("Failed to initialize migrations", "error", err)
		os.Exit(1)
	}

	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			logger.Info("No migrations to apply")
			return
		}
		logger.Error("Migration failed", "error", err)
		os.Exit(1)
	}

	logger.Info("Migrations applied")
}
