package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)


type apiConfig struct {
	db          *database.Queries
	dbConn      *sql.DB
	l           *seeding.Lookup
	t           *TypeLookup
	platform    string
	adminApiKey string
	host        string
}


func apiConfigInit() (apiConfig, error) {
	const domain = "localhost:8080"

	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		return apiConfig{}, errors.New("DB_URL must be set")
	}

	platform := os.Getenv("PLATFORM")
	if platform == "" {
		return apiConfig{}, errors.New("PLATFORM must be set")
	}

	adminApiKey := os.Getenv("ADMIN_API_KEY")
	if adminApiKey == "" {
		return apiConfig{}, errors.New("ADMIN_API_KEY must be set")
	}

	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		return apiConfig{}, fmt.Errorf("Error opening database: %v", err)
	}
	dbQueries := database.New(dbConn)

	apiCfg := apiConfig{
		db:          dbQueries,
		dbConn:      dbConn,
		platform:    platform,
		adminApiKey: adminApiKey,
		host:        domain,
	}

	apiCfg.l, err = seeding.SeedDatabase(apiCfg.db, apiCfg.dbConn)
	if err != nil {
		return apiConfig{}, err
	}

	typeLookup := TypeLookupInit()
	apiCfg.t = &typeLookup

	return apiCfg, nil
}