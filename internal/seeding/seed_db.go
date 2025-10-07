package seeding

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func Seed_database(db *database.Queries, dbConn *sql.DB) error {
	lookup := lookupInit()
	seedFunctions := []func(*database.Queries, *sql.DB) error {
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

	start := time.Now()

	for _, seedFunc := range seedFunctions {
        if err := seedFunc(db, dbConn); err != nil {
            return err
        }
    }

	duration := time.Since(start)

	fmt.Printf("seeding took %.2f seconds\n", duration.Seconds())

	return nil
}