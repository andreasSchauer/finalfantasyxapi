package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type PlayerAbility struct {
	ID int32
	Ability
	Description         *string             `json:"description"`
	Effect              string              `json:"effect"`
	Topmenu             *string             `json:"topmenu"`
	CanUseOutsideBattle bool                `json:"can_use_outside_battle"`
	MPCost              *int32              `json:"mp_cost"`
	Cursor              *string             `json:"cursor"`
	BattleInteractions  []BattleInteraction `json:"battle_interactions"`
}

func (p PlayerAbility) ToHashFields() []any {
	return []any{
		p.Ability.ID,
		derefOrNil(p.Description),
		p.Effect,
		derefOrNil(p.Topmenu),
		p.CanUseOutsideBattle,
		derefOrNil(p.MPCost),
		derefOrNil(p.Cursor),
	}
}

func (p PlayerAbility) GetID() int32 {
	return p.ID
}

func (p PlayerAbility) GetAbilityRef() AbilityReference {
	return AbilityReference{
		Name:        p.Name,
		Version:     p.Version,
		AbilityType: string(database.AbilityTypePlayerAbility),
	}
}

func (l *lookup) seedPlayerAbilities(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/player_abilities.json"

	var playerAbilities []PlayerAbility

	err := loadJSONFile(string(srcPath), &playerAbilities)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, playerAbility := range playerAbilities {
			var err error
			playerAbility.Type = database.AbilityTypePlayerAbility

			playerAbility.Ability, err = seedObjAssignID(qtx, playerAbility.Ability, l.seedAbility)
			if err != nil {
				return err
			}

			dbPlayerAbility, err := qtx.CreatePlayerAbility(context.Background(), database.CreatePlayerAbilityParams{
				DataHash:            generateDataHash(playerAbility),
				AbilityID:           playerAbility.Ability.ID,
				Description:         getNullString(playerAbility.Description),
				Effect:              playerAbility.Effect,
				Topmenu:             nullTopmenuType(playerAbility.Topmenu),
				CanUseOutsideBattle: playerAbility.CanUseOutsideBattle,
				MpCost:              getNullInt32(playerAbility.MPCost),
				Cursor:              nullTargetType(playerAbility.Cursor),
			})
			if err != nil {
				return fmt.Errorf("couldn't create Player Ability: %s: %v", playerAbility.Name, err)
			}

			playerAbility.ID = dbPlayerAbility.ID
			key := createLookupKey(playerAbility.Ability)
			l.playerAbilities[key] = playerAbility
		}
		return nil
	})
}

func (l *lookup) createPlayerAbilitiesRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/player_abilities.json"

	var playerAbilities []PlayerAbility

	err := loadJSONFile(string(srcPath), &playerAbilities)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonAbility := range playerAbilities {
			abilityRef := jsonAbility.GetAbilityRef()

			ability, err := l.getPlayerAbility(abilityRef)
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
