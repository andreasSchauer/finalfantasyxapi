package seeding

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/pressly/goose/v3"
)

func SeedDatabase(db *database.Queries, dbConn *sql.DB, migrationsDir string) error {
	lookup := lookupInit()

	seedFunctions := []func(*database.Queries, *sql.DB) error{
		lookup.seedStats,
		lookup.seedElements,
		lookup.seedAffinities,
		lookup.seedAgilityTiers,
		lookup.seedOverdriveModes,
		lookup.seedStatusConditions,
		lookup.seedProperties,
		lookup.seedCharacters,
		lookup.seedAeons,
		lookup.seedDefaultAbilitiesEntries,
		lookup.seedBlitzballItems,
		lookup.seedMonsterArenaCreations,
		lookup.seedSidequests,
		lookup.seedMonsters,
		lookup.seedAeonCommands,
		lookup.seedMenuCommands,
		lookup.seedPlayerAbilities,
		lookup.seedEnemyAbilities,
		lookup.seedOverdriveAbilities,
		lookup.seedTriggerCommands,
		lookup.seedOverdriveCommands,
		lookup.seedItems,
		lookup.seedKeyItems,
		lookup.seedPrimers,
		lookup.seedMixes,
		lookup.seedCelestialWeapons,
		lookup.seedAutoAbilities,
		lookup.seedEquipment,
		lookup.seedLocations,
		lookup.seedTreasures,
		lookup.seedShops,
		lookup.seedMonsterFormations,
		lookup.seedSongs,
		lookup.seedFMVs,
	}

	err := databaseSetup(dbConn, migrationsDir)
	if err != nil {
		return fmt.Errorf("couldn't setup database: %v", err)
	}

	seedStart := time.Now()

	for _, seedFunc := range seedFunctions {
		if err := seedFunc(db, dbConn); err != nil {
			return err
		}
	}

	seedDuration := time.Since(seedStart)
	fmt.Printf("initial seeding took %.2f seconds\n", seedDuration.Seconds())

	// will do relationships here

	relationshipFunctions := []func(*database.Queries, *sql.DB) error{
		lookup.createStatsRelationships,
	}

	relStart := time.Now()

	for _, relFunc := range relationshipFunctions {
		if err := relFunc(db, dbConn); err != nil {
			return err
		}
	}

	relDuration := time.Since(relStart)
	fmt.Printf("establishing relationships took %.2f seconds\n", relDuration.Seconds())

	return nil
}


func databaseSetup(dbConn *sql.DB, migrationsDir string) error {
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