package seeding

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/pressly/goose/v3"
)


func SeedDatabase(db *database.Queries, dbConn *sql.DB) error {
	const migrationsDir = "./sql/schema/"

	err := setupDB(dbConn, migrationsDir)
	if err != nil {
		return fmt.Errorf("couldn't setup database: %v", err)
	}

	l := lookupInit()

	seedFunctions := []func(*database.Queries, *sql.DB) error{
		l.seedStats,
		l.seedElements,
		l.seedAffinities,
		l.seedAgilityTiers,
		l.seedOverdriveModes,
		l.seedStatusConditions,
		l.seedProperties,
		l.seedModifiers,
		l.seedCharacters,
		l.seedAeons,
		l.seedBlitzballItems,
		l.seedSidequests,
		l.seedMonsterArenaCreations,
		l.seedMonsters,
		l.seedAeonCommands,
		l.seedSubmenus,
		l.seedPlayerAbilities,
		l.seedEnemyAbilities,
		l.seedOverdriveAbilities,
		l.seedTriggerCommands,
		l.seedOverdriveCommands,
		l.seedItems,
		l.seedKeyItems,
		l.seedPrimers,
		l.seedMixes,
		l.seedCelestialWeapons,
		l.seedAutoAbilities,
		l.seedEquipment,
		l.seedLocations,
		l.seedTreasures,
		l.seedShops,
		l.seedMonsterFormations,
		l.seedSongs,
		l.seedFMVs,
	}

	relationshipFunctions := []func(*database.Queries, *sql.DB) error{
		l.createStatsRelationships,
		l.createElementsRelationships,
		l.createOverdriveModesRelationships,
		l.createStatusConditionsRelationships,
		l.createPropertiesRelationships,
		l.createCharactersRelationships,
		l.createSubmenusRelationships,
		l.createMixesRelationships,
		l.createCelestialWeaponsRelationships,
		l.createAutoAbilitiesRelationships,
		l.createEquipmentRelationships,
		l.createDefaultAbilitiesRelationships,
		l.createAeonsRelationships,
		l.createAreasRelationships,
	}

	err = handleDBFunctions(db, dbConn, seedFunctions, "initial seeding")
	if err != nil {
		return err
	}

	err = handleDBFunctions(db, dbConn, relationshipFunctions, "establishing relationships")
	if err != nil {
		return err
	}

	return nil
}


func handleDBFunctions(db *database.Queries, dbConn *sql.DB, functions []func(*database.Queries, *sql.DB) error, infotext string) error {
	start := time.Now()

	for _, function := range functions {
		err := function(db, dbConn)
		if err != nil {
			return err
		}
	}

	duration := time.Since(start)
	fmt.Printf("%s took %.2f seconds\n", infotext, duration.Seconds())

	return nil
}

func setupDB(dbConn *sql.DB, migrationsDir string) error {
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
