package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type DefaultAbilitiesEntry struct {
	Name 				string 				`json:"name"`
	DefaultAbilities 	[]AbilityReference 	`json:"default_abilities"`
}



func (l *lookup) seedDefaultAbilitiesRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/default_abilities.json"

	var entries []DefaultAbilitiesEntry
	err := loadJSONFile(string(srcPath), &entries)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, entry := range entries {
			class, err := l.getCharacterClass(entry.Name)
			if err != nil {
				return err
			}

			for _, abilityRef := range entry.DefaultAbilities {
				junction, err := createJunction(class, abilityRef, l.getPlayerAbility)
				if err != nil {
					return fmt.Errorf("couldn't create junction between character class %s and ability %s: %v", class.Name, createLookupKey(abilityRef), err)
				}

				err = qtx.CreateDefaultAbility(context.Background(), database.CreateDefaultAbilityParams{
					DataHash: generateDataHash(junction),
					ClassID: junction.ParentID,
					AbilityID: junction.ChildID,
				})
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
}