package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	db             *database.Queries
}

func main() {
	//const filepathRoot = "."
	//const port = "8080"
	const testmail = "idontknow@gmail.com"

	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL must be set")
	}

	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database: %s", err)
	}
	dbQueries := database.New(dbConn)

	apiCfg := apiConfig{
		db:             dbQueries,
	}

	for i := range 10 {
		email := testmail + strconv.Itoa(i+1)
		apiCfg.handlerUsersCreate(email)
	}
}



func (cfg *apiConfig) handlerUsersCreate(email string) {
	_, err := cfg.db.CreateUser(context.Background(), email)
	if err != nil {
		fmt.Println(err)
	}
}