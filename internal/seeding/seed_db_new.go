package seeding

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func Seed(db *database.Queries, dbConn *sql.DB) (*Lookup, error) {
	const migrationsDir = "sql/schema/"
	fullPath, err := h.GetAbsoluteFilepath(migrationsDir)
	if err != nil {
		return nil, err
	}

	start := time.Now()

	err = setupDB(dbConn, fullPath)
	if err != nil {
		return nil, fmt.Errorf("couldn't setup database: %v", err)
	}
	
	l := lookupInit()
	err = l.loadJSONFiles()
	if err != nil {
		return nil, err
	}

	err = queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		err = l.seedLoop1(qtx, context.Background())
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	
	totalDuration := time.Since(start)
	fmt.Printf("database seeding took %.3f seconds\n\n", totalDuration.Seconds())

	return l, nil
}


func (l *Lookup) seedLoop(qtx *database.Queries, ctx context.Context, fns []func(*database.Queries, context.Context) error) error {
	for _, fn := range fns {
		err := fn(qtx, ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func dedupeRows[T Hashable](rows []T, hashes map[string]int32) []T {
    seen := make(map[string]bool)
    ordered := []T{}

    for _, row := range rows {
        hash := generateDataHash(row)
		_, ok := hashes[hash]
		if ok || seen[hash] {
			continue
		}

		seen[hash] = true
		ordered = append(ordered, row)
    }
    return ordered
}