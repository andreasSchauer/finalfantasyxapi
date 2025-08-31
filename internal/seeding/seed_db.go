package seeding

import(
	"database/sql"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)



func Seed_database(db *database.Queries, dbConn *sql.DB) error {
	seedFunctions := []func(*database.Queries, *sql.DB) error{
		seedStats,
		seedElements,
		seedAffinities,
		seedAgilityTiers,
		seedOverdriveModes,
		seedStatusConditions,
		seedProperties,
	}

	for _, fn := range seedFunctions {
		err := fn(db, dbConn)
		if err != nil {
			return err
		}
	}

	return nil
}