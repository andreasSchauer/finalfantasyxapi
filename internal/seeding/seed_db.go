package seeding

import (
	"database/sql"

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
		seedSidequests,
		seedMonsterArenaCreations,
		seedTreasures,
		seedLocations,
		seedCelestialWeapons,
		seedAutoAbilities,
		seedMonsters,
	}

	for _, seedFunc := range seedFunctions {
        if err := seedFunc(db, dbConn); err != nil {
            return err
        }
    }

	return nil
}