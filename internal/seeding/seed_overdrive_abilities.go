package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type OverdriveAbility struct {
	ID					int32
	Ability
	RelatedStats		[]string			`json:"related_stats"`
	BattleInteractions 	[]BattleInteraction `json:"battle_interactions"`
}

func (a OverdriveAbility) ToHashFields() []any {
	return []any{
		a.Ability.ID,
	}
}

func (a OverdriveAbility) GetID() int32 {
	return a.ID
}


func (a OverdriveAbility) GetAbilityRef() AbilityReference {
	return AbilityReference{
		Name:        a.Name,
		Version:     a.Version,
		AbilityType: string(database.AbilityTypeOverdriveAbility),
	}
}

func (a OverdriveAbility) Error() string {
	return fmt.Sprintf("overdrive ability %s, version %v", a.Name, derefOrNil(a.Version))
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


func (l *lookup) seedOverdriveAbilitiesRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/overdrive_abilities.json"

	var overdriveAbilities []OverdriveAbility

	err := loadJSONFile(string(srcPath), &overdriveAbilities)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonAbility := range overdriveAbilities {
			abilityRef := jsonAbility.GetAbilityRef()

			ability, err := l.getOverdriveAbility(abilityRef)
			if err != nil {
				return err
			}

			err = l.seedOverdriveAbilityRelatedStats(qtx, ability)
			if err != nil {
				return err
			}

			l.currentAbility = ability.Ability

			err = l.seedBattleInteractions(qtx, l.currentAbility, ability.BattleInteractions)
			if err != nil {
				return err
			}
		}

		return nil
	})
}



func (l *lookup) seedOverdriveAbilityRelatedStats(qtx *database.Queries, ability OverdriveAbility) error {
	for _, jsonStat := range ability.RelatedStats {
		junction, err := createJunction(ability, jsonStat, l.getStat)
		if err != nil {
			return err
		}

		err = qtx.CreateOverdriveAbilitiesRelatedStatsJunction(context.Background(), database.CreateOverdriveAbilitiesRelatedStatsJunctionParams{
			DataHash: 			generateDataHash(junction),
			OverdriveAbilityID: junction.ParentID,
			StatID: 			junction.ChildID,
		})
		if err != nil {
			return fmt.Errorf("ability %s: %v", createLookupKey(ability.Ability), err)
		}
	}

	return nil
}