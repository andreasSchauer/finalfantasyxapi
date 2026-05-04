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
	const srcPath = "data/default_abilities.json"

	var entries []DefaultAbilitiesEntry
	err := loadJSONFile(string(srcPath), &entries)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, entry := range entries {
			class, err := GetResource(entry.Name, l.CharClasses)
			if err != nil {
				return err
			}

			err = l.seedCharClassDefaultAbilities(qtx, class, entry)
			if err != nil {
				return h.NewErr(class.Error(), err)
			}
		}
		return nil
	})
}

func (l *Lookup) seedCharClassDefaultAbilities(qtx *database.Queries, class CharacterClass, entry DefaultAbilitiesEntry) error {
	for _, abilityRef := range entry.DefaultAbilities {
		junction, err := createJunction(class, abilityRef, l.Abilities)
		if err != nil {
			return err
		}

		err = qtx.CreateDefaultAbilityJunction(context.Background(), database.CreateDefaultAbilityJunctionParams{
			DataHash:  generateDataHash(junction),
			ClassID:   junction.ParentID,
			AbilityID: junction.ChildID,
		})
		if err != nil {
			return h.NewErr(abilityRef.Error(), err, "couldn't junction default ability")
		}
	}

	return nil
}


func (l *Lookup) seedJuncDefaultAbilities(qtx *database.Queries, ctx context.Context) error {
	const desc string = "default abilities"
	params := database.CreateDefaultAbilityJunctionBulkParams{
		DataHash: 	make([]string, 0),
		ClassID: 	make([]int32, 0),
		AbilityID: 	make([]int32, 0),
	}

	for _, entry := range l.json.defaultAbilities {
		class, err := GetResource(entry.Name, l.CharClasses)
		if err != nil {
			return err
		}

		for _, ref := range entry.DefaultAbilities {
			ability, err := GetResource(ref, l.Abilities)
			if err != nil {
				return err
			}

			j := StdJunction{
				ParentID: 	class.ID,
				ChildID: 	ability.ID,
			}
			dataHash := generateJunctionHash(j, desc)

			params.DataHash = append(params.DataHash, dataHash)
			params.ClassID = append(params.ClassID, class.ID)
			params.AbilityID = append(params.AbilityID, ability.ID)
		}
	}

	return qtx.CreateDefaultAbilityJunctionBulk(ctx, params)
}