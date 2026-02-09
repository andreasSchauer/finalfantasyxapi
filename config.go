package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
	_ "github.com/lib/pq"
)

type Config struct {
	db          	*database.Queries
	dbConn      	*sql.DB
	l           	*seeding.Lookup
	t           	*TypeLookup
	q				*QueryLookup
	e				*endpoints
	platform   		string
	adminApiKey 	string
	host        	string
}

func ConfigInit() (*Config, error) {
	const domain = "localhost:8080"

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
	dbQueries := database.New(dbConn)

	apiCfg := Config{
		db:          dbQueries,
		dbConn:      dbConn,
		platform:    platform,
		adminApiKey: adminApiKey,
		host:        domain,
	}

	apiCfg.l, err = seeding.SeedDatabase(apiCfg.db, apiCfg.dbConn)
	if err != nil {
		return nil, err
	}

	apiCfg.TypeLookupInit()
	apiCfg.QueryLookupInit()
	apiCfg.EndpointsInit()

	return &apiCfg, nil
}
