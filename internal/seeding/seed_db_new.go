package seeding

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/pressly/goose/v3"
)

func Seed(db *database.Queries, dbConn *sql.DB) (*Lookup, error) {
	const migrationsDir = "sql/schema/"
	fullPath, err := h.GetAbsoluteFilepath(migrationsDir)
	if err != nil {
		return nil, err
	}

	err = setupDB(dbConn, fullPath)
	if err != nil {
		return nil, fmt.Errorf("couldn't setup database: %v", err)
	}

	l, err := lookupInit()
	if err != nil {
		return nil, err
	}	

	err = queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		start := time.Now()

		err := l.seedLoop(qtx, context.Background(), []seedFunc{
			l.seedLoop1,
			l.seedLoop2,
			l.seedLoop3,
			l.seedLoop4,
			l.seedLoop5,
			l.seedLoop6,
			l.seedLoop7,
		})
		if err != nil {
			return err
		}

		duration := time.Since(start)
		fmt.Printf("database seeding took %.3f seconds\n\n", duration.Seconds())

		return nil
	})
	if err != nil {
		return nil, err
	}

	return l, nil
}



func setupDB(dbConn *sql.DB, migrationsDir string) error {
	start := time.Now()

	err := goose.SetDialect("postgres")
	if err != nil {
		return err
	}

	err = goose.DownTo(dbConn, migrationsDir, 0)
	if err != nil {
		return err
	}

	err = goose.Up(dbConn, migrationsDir)
	if err != nil {
		return err
	}

	duration := time.Since(start)
	fmt.Printf("\ndatabase setup took %.3f seconds\n", duration.Seconds())

	return nil
}


func dedupeRows[T Hashable](rows []T, hashes map[string]int32) []T {
    seen := make(map[string]bool)
    new := []T{}

    for _, row := range rows {
        hash := generateDataHash(row)
		
		_, ok := hashes[hash]
		if ok || seen[hash] {
			continue
		}

		seen[hash] = true
		new = append(new, row)
    }
    return new
}