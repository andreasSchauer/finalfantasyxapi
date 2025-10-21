package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type OverdriveAbility struct {
	Ability
}

func (a OverdriveAbility) ToHashFields() []any {
	return []any{
		a.Ability.ID,
	}
}

func (l *lookup) seedOverdriveAbilities(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/overdrive_abilities.json"

	var overdriveAbilities []OverdriveAbility

	err := loadJSONFile(string(srcPath), &overdriveAbilities)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, overdriveAbility := range overdriveAbilities {
			var err error
			overdriveAbility.Type = database.AbilityTypeOverdriveAbility

			overdriveAbility.Ability, err = seedObjAssignID(qtx, overdriveAbility.Ability, l.seedAbility)
			if err != nil {
				return err
			}

			dbOverdriveAbility, err := qtx.CreateOverdriveAbility(context.Background(), database.CreateOverdriveAbilityParams{
				DataHash:  generateDataHash(overdriveAbility),
				AbilityID: overdriveAbility.Ability.ID,
			})
			if err != nil {
				return fmt.Errorf("couldn't create Overdrive Ability: %s: %v", overdriveAbility.Name, err)
			}

			overdriveAbility.ID = dbOverdriveAbility.ID
			key := createLookupKey(overdriveAbility.Ability)
			l.overdriveAbilities[key] = overdriveAbility
		}
		return nil
	})
}
