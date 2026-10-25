package api

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
	db          *database.Queries
	DbConn      *sql.DB
	L           *seeding.Lookup
	T           *Enums
	Q           *QueryLookup
	E           *endpoints
	platform    string
	adminApiKey string
	host        string
	fetchLimit  int
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

	adminApiKey := os.Getenv("ADMIN_API_KEY")
	if adminApiKey == "" {
		return nil, errors.New("ADMIN_API_KEY must be set")
	}

	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("Error opening database: %v", err)
	}
	dbConn.SetMaxOpenConns(connLimit)
	dbConn.SetMaxIdleConns(connLimit)
	dbConn.SetConnMaxLifetime(time.Hour)
	dbQueries := database.New(dbConn)

	apiCfg := Config{
		db:          dbQueries,
		DbConn:      dbConn,
		platform:    platform,
		adminApiKey: adminApiKey,
		host:        domain,
		fetchLimit:  fetchLimit,
	}

	apiCfg.L, err = seeding.LoadDatabase(apiCfg.db, apiCfg.DbConn)
	if err != nil {
		return nil, err
	}

	apiCfg.EnumLookupInit()
	apiCfg.EndpointsInit()
	apiCfg.QueryLookupInit()

	return &apiCfg, nil
}
