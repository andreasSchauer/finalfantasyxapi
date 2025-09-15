package seeding

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func Seed_database(db *database.Queries, dbConn *sql.DB) error {
	seedFunctions := []func(*database.Queries, *sql.DB) error {
		seedStats,
		seedElements,
		seedAffinities,
		seedAgilityTiers,
		seedOverdriveModes,
		seedStatusConditions,
		seedProperties,
		seedCharacters,
		seedAeons,
		seedDefaultAbilitiesEntries,
		seedBlitzballItems,
		seedMonsterArenaCreations,
		seedSidequests,
		seedMonsters,
		seedAeonCommands,
		seedMenuCommands,
		seedPlayerAbilities,
		seedEnemyAbilities,
		seedOverdriveAbilities,
		seedTriggerCommands,
		seedOverdriveCommands,
		seedItems,
		seedKeyItems,
		seedPrimers,
		seedMixes,
		seedCelestialWeapons,
		seedAutoAbilities,
		seedEquipment,
		seedLocations,
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