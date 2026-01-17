package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type OverdriveAbility struct {
	ID int32
	Ability
	RelatedStats       []string            `json:"related_stats"`
	BattleInteractions []BattleInteraction `json:"battle_interactions"`
}

func (o OverdriveAbility) ToHashFields() []any {
	return []any{
		o.Ability.ID,
	}
}

func (o OverdriveAbility) ToKeyFields() []any {
	return []any{
		o.Ability.Name,
		h.DerefOrNil(o.Ability.Version),
	}
}

func (o OverdriveAbility) GetID() int32 {
	return o.ID
}

func (o OverdriveAbility) GetAbilityRef() AbilityReference {
	return AbilityReference{
		Name:        o.Name,
		Version:     o.Version,
		AbilityType: string(database.AbilityTypeOverdriveAbility),
	}
}

func (o OverdriveAbility) Error() string {
	return fmt.Sprintf("overdrive ability %s, version %v", o.Name, h.DerefOrNil(o.Version))
}

func (o OverdriveAbility) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID: o.ID,
		Name: o.Name,
		Version: o.Version,
		Specification: o.Specification,
	}
}

func (l *Lookup) seedOverdriveAbilities(db *database.Queries, dbConn *sql.DB) error {
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
				return h.NewErr(overdriveAbility.Error(), err)
			}

			dbOverdriveAbility, err := qtx.CreateOverdriveAbility(context.Background(), database.CreateOverdriveAbilityParams{
				DataHash:  generateDataHash(overdriveAbility),
				AbilityID: overdriveAbility.Ability.ID,
			})
			if err != nil {
				return h.NewErr(overdriveAbility.Error(), err, "couldn't create overdrive ability")
			}

			overdriveAbility.ID = dbOverdriveAbility.ID
			key := CreateLookupKey(overdriveAbility)
			l.OverdriveAbilities[key] = overdriveAbility
			l.OverdriveAbilitiesID[overdriveAbility.ID] = overdriveAbility
		}
		return nil
	})
}

func (l *Lookup) seedOverdriveAbilitiesRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/overdrive_abilities.json"

	var overdriveAbilities []OverdriveAbility

	err := loadJSONFile(string(srcPath), &overdriveAbilities)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonAbility := range overdriveAbilities {
			abilityRef := jsonAbility.GetAbilityRef()

			overdriveAbility, err := GetResource(abilityRef.Untyped(), l.OverdriveAbilities)
			if err != nil {
				return h.NewErr(abilityRef.Error(), err)
			}

			err = l.seedOverdriveAbilityRelatedStats(qtx, overdriveAbility)
			if err != nil {
				return h.NewErr(overdriveAbility.Error(), err)
			}

			l.currentAbility = overdriveAbility.Ability

			err = l.seedBattleInteractions(qtx, l.currentAbility, overdriveAbility.BattleInteractions)
			if err != nil {
				return h.NewErr(overdriveAbility.Error(), err)
			}
		}

		return nil
	})
}

func (l *Lookup) seedOverdriveAbilityRelatedStats(qtx *database.Queries, ability OverdriveAbility) error {
	for _, jsonStat := range ability.RelatedStats {
		junction, err := createJunction(ability, jsonStat, l.Stats)
		if err != nil {
			return err
		}

		err = qtx.CreateOverdriveAbilitiesRelatedStatsJunction(context.Background(), database.CreateOverdriveAbilitiesRelatedStatsJunctionParams{
			DataHash:           generateDataHash(junction),
			OverdriveAbilityID: junction.ParentID,
			StatID:             junction.ChildID,
		})
		if err != nil {
			return h.NewErr(jsonStat, err, "couldn't junction related stat")
		}
	}

	return nil
}
