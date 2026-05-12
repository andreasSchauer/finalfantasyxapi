package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/pressly/goose/v3"
)

func SeedDatabase(db *database.Queries, dbConn *sql.DB) (*Lookup, error) {
	defer fmt.Println()
	defer h.MeasureTime("database seeding")()
	ctx := context.Background()

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
		defer h.MeasureTime("initial database seeding")()
		fmt.Println("initial database seeding...")

		return l.seedLoop(qtx, ctx, []seedFunc{
			l.seedLoop1,
			l.seedLoop2,
			l.seedLoop3,
			l.seedLoop4,
			l.seedLoop5,
			l.seedLoop6,
			l.seedLoop7,
		})
	})
	if err != nil {
		return nil, err
	}

	err = l.completeLookups()
	if err != nil {
		return nil, err
	}

	err = queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		defer h.MeasureTime("seeding junctions")()
		return l.seedJunctions(qtx, ctx)
	})
	if err != nil {
		return nil, err
	}

	err = refreshViews(db, ctx)
	if err != nil {
		return nil, err
	}

	fmt.Println()

	return l, nil
}


func setupDB(dbConn *sql.DB, migrationsDir string) error {
	defer h.MeasureTime("\ndatabase setup")()

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

	return nil
}

func (l *Lookup) completeLookups() error {
	defer h.MeasureTime("lookup completion")()

	fns := []func() error{
		l.completeEnemyAbilities,
		l.completeItems,
		l.completeOverdriveAbilities,
		l.completePlayerAbilities,
		l.completeTriggerCommands,
		l.completeMiscAbilities,
		l.completeAeons,
		l.completeAeonStats,
		l.completeAgilityTiers,
		l.completeAutoAbilities,
		l.completeBlitzballPositions,
		l.completeCharacters,
		l.completeEquipment,
		l.completeLocations,
		l.completeMixes,
		l.completeMonsterFormations,
		l.completeMonsters,
		l.completeOverdriveModes,
		l.completeShops,
		l.completeSidequests,
		l.completeSongs,
		l.completeStatusConditions,
		l.completeTreasureLists,
	}

	for _, fn := range fns {
		err := fn()
		if err != nil {
			return err
		}
	}

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


func refreshViews(db *database.Queries, ctx context.Context) error {
	defer h.MeasureTime("refreshing views")()
	
	fns := []func(context.Context) error{
		db.RefreshMonsterItemDropsView,
		db.RefreshMonsterEquipmentDropsView,
		db.RefreshMonsterEncountersView,
		db.RefreshGeographyView,
		db.RefreshGeographyGraphView,
		db.RefreshItemSourcesView,
		db.RefreshEquipmentSourcesView,
		db.RefreshAbilitiesView,
	}

	for _, fn := range fns {
		err := fn(ctx)
		if err != nil {
			return nil
		}
	}

	return nil
}