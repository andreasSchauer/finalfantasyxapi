// Package config asdoasd
package config

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
	_ "github.com/lib/pq"
)

type Config struct {
	DB          *database.Queries
	DBConn      *sql.DB
	Platform    string
	AdminAPIKey string
	Host        string
	FetchLimit  int
}

func ConfigInit() (*Config, error) {
	const domain = "localhost:8080"
	const fetchLimit = 1000
	const connLimit = 25

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		return nil, errors.New("DB_URL must be set")
	}

	platform := os.Getenv("PLATFORM")
	if platform == "" {
		return nil, errors.New("PLATFORM must be set")
	}

	adminAPIKey := os.Getenv("ADMIN_API_KEY")
	if adminAPIKey == "" {
		return nil, errors.New("ADMIN_API_KEY must be set")
	}

	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}
	dbConn.SetMaxOpenConns(connLimit)
	dbConn.SetMaxIdleConns(connLimit)
	dbConn.SetConnMaxLifetime(time.Hour)
	dbQueries := database.New(dbConn)

	cfg := Config{
		DB:          dbQueries,
		DBConn:      dbConn,
		Platform:    platform,
		AdminAPIKey: adminAPIKey,
		Host:        domain,
		FetchLimit:  fetchLimit,
	}

	err = seeding.LoadDatabase(cfg.DB, cfg.DBConn)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
