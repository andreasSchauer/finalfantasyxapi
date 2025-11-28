package seeding

import (
	"context"
	"database/sql"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type DefaultAbilitiesEntry struct {
	Name             string             `json:"name"`
	DefaultAbilities []AbilityReference `json:"default_abilities"`
}

func (l *Lookup) seedDefaultAbilitiesRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/default_abilities.json"

	var entries []DefaultAbilitiesEntry
	err := loadJSONFile(string(srcPath), &entries)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, entry := range entries {
			class, err := getResource(entry.Name, l.charClasses)
			if err != nil {
				return err
			}

			err = l.seedCharClassDefaultAbilities(qtx, class, entry)
			if err != nil {
				return h.GetErr(class.Error(), err)
			}
		}
		return nil
	})
}

func (l *Lookup) seedCharClassDefaultAbilities(qtx *database.Queries, class CharacterClass, entry DefaultAbilitiesEntry) error {
	for _, abilityRef := range entry.DefaultAbilities {
		junction, err := createJunction(class, abilityRef, l.playerAbilities)
		if err != nil {
			return err
		}

		err = qtx.CreateDefaultAbility(context.Background(), database.CreateDefaultAbilityParams{
			DataHash:  generateDataHash(junction),
			ClassID:   junction.ParentID,
			AbilityID: junction.ChildID,
		})
		if err != nil {
			return h.GetErr(abilityRef.Error(), err, "couldn't junction default ability")
		}
	}

	return nil
}
