// Package config asdoasd
package config

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seeding"
	_ "github.com/lib/pq"
)

var Cfg = &Config{}

type Config struct {
	DB          *database.Queries
	DBConn      *sql.DB
	Platform    string
	AdminAPIKey string
	Host        string
	FetchLimit  int
}

func (cfg *Config) Init() error {
	const domain = "localhost:8080"
	const fetchLimit = 1000
	const connLimit = 25

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		return errors.New("DB_URL must be set")
	}

	platform := os.Getenv("PLATFORM")
	if platform == "" {
		return errors.New("PLATFORM must be set")
	}

	adminAPIKey := os.Getenv("ADMIN_API_KEY")
	if adminAPIKey == "" {
		return errors.New("ADMIN_API_KEY must be set")
	}

	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}
	dbConn.SetMaxOpenConns(connLimit)
	dbConn.SetMaxIdleConns(connLimit)
	dbConn.SetConnMaxLifetime(time.Hour)
	dbQueries := database.New(dbConn)

	cfg.DB = dbQueries
	cfg.DBConn = dbConn
	cfg.Platform = platform
	cfg.AdminAPIKey = adminAPIKey
	cfg.Host = domain
	cfg.FetchLimit = fetchLimit

	err = seeding.LoadDatabase(cfg.DB, cfg.DBConn)
	if err != nil {
		return err
	}

	return nil
}
