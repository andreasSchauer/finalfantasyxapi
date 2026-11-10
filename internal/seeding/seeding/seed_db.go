// Package seeding includes the whole process from mapping the raw JSON files to go structs and populating the database, as well as the lookup that is used for the api
package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
	"github.com/pressly/goose/v3"
)

func LoadDatabase(db *database.Queries, dbConn *sql.DB) error {
	defer fmt.Println()
	defer h.MeasureTime("database seeding")()
	ctx := context.Background()
	lookup.L.Init()

	err := rawJSON.loadFiles()
	if err != nil {
		return err
	}

	// if schema needs to be converted to something else, do it here (use h.ConvertJsonData)

	cacheHash, err := getCacheHash()
	if err != nil {
		return err
	}

	dbHash, _ := db.GetDbState(ctx)

	if cacheHash == dbHash {
		err = skipSeeding(db, ctx)
		if err != nil {
			return err
		}

		return nil
	}

	err = seedDatabase(db, dbConn, ctx, cacheHash)
	if err != nil {
		return err
	}

	return nil
}

func skipSeeding(db *database.Queries, ctx context.Context) error {
	fmt.Println("database is up to date")
	fmt.Println()

	err := lookup.L.InitCached()
	if err != nil {
		return fmt.Errorf("couldn't assign lookups: %v", err)
	}

	err = refreshViews(db, ctx)
	if err != nil {
		return err
	}

	fmt.Println()

	return nil
}

func seedDatabase(db *database.Queries, dbConn *sql.DB, ctx context.Context, dataHash string) error {
	fmt.Println("updating database...")
	fmt.Println()

	const migrationsDir = "sql/schema/"
	fullPath, err := h.GetAbsoluteFilepath(migrationsDir)
	if err != nil {
		return err
	}

	err = setupDB(dbConn, fullPath)
	if err != nil {
		return fmt.Errorf("couldn't setup database: %v", err)
	}

	err = queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		defer h.MeasureTime("initial database seeding")()
		fmt.Println("initial database seeding...")

		return seedLoop(qtx, ctx, []seedFunc{
			seedLoop1,
			seedLoop2,
			seedLoop3,
			seedLoop4,
			seedLoop5,
			seedLoop6,
			seedLoop7,
		})
	})
	if err != nil {
		return err
	}

	err = completeLookups()
	if err != nil {
		return err
	}

	err = lookup.L.Cache()
	if err != nil {
		return err
	}

	err = queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		defer h.MeasureTime("seeding junctions")()
		return seedJunctions(qtx, ctx)
	})
	if err != nil {
		return err
	}

	err = queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		return db.UpdateDbState(ctx, dataHash)
	})
	if err != nil {
		return fmt.Errorf("couldn't update dbState hash: %v", err)
	}

	err = refreshViews(db, ctx)
	if err != nil {
		return err
	}

	fmt.Println()

	return nil
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

func completeLookups() error {
	defer h.MeasureTime("lookup completion")()

	fns := []func() error{
		completeEnemyAbilities,
		completeItems,
		completeOverdriveAbilities,
		completePlayerAbilities,
		completeTriggerCommands,
		completeMiscAbilities,
		completeAeons,
		completeAeonStats,
		completeAgilityTiers,
		completeAutoAbilities,
		completeBlitzballPositions,
		completeCharacters,
		completeEquipment,
		completeLocations,
		completeMixes,
		completeMonsterFormations,
		completeMonsters,
		completeOverdriveModes,
		completeShops,
		completeSidequests,
		completeSongs,
		completeStatusConditions,
		completeTreasureLists,
	}

	for _, fn := range fns {
		err := fn()
		if err != nil {
			return err
		}
	}

	return nil
}

func dedupeRows[T seedtypes.Hashable](rows []T, hashes map[string]int32) []T {
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
		db.RefreshGeographyView,
		db.RefreshGeographyGraphView,
		db.RefreshMonsterEncountersView,
		db.RefreshMonsterItemDropsView,
		db.RefreshMonsterEquipmentDropsView,
		db.RefreshItemSourcesView,
		db.RefreshEquipmentSourcesView,
		db.RefreshShopEquipAvlsView,
		db.RefreshAutoAbilitySourcesView,
		db.RefreshAvailabilityView,
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
