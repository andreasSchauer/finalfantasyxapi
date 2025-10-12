package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type PlayerAbility struct {
	Ability
	AbilityAttributes
	AbilityID           int32
	Description         *string `json:"description"`
	Effect              string  `json:"effect"`
	Topmenu             *string `json:"topmenu"`
	CanUseOutsideBattle bool    `json:"can_use_outside_battle"`
	MPCost              *int32  `json:"mp_cost"`
	Cursor              *string `json:"cursor"`
}

func (a PlayerAbility) ToHashFields() []any {
	return []any{
		a.AbilityID,
		derefOrNil(a.Description),
		a.Effect,
		derefOrNil(a.Topmenu),
		a.CanUseOutsideBattle,
		derefOrNil(a.MPCost),
		derefOrNil(a.Cursor),
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
			ability := playerAbility.Ability
			attributes := playerAbility.AbilityAttributes
			ability.Type = database.AbilityTypePlayerAbility

			dbAbility, err := l.seedAbility(qtx, attributes, ability)
			if err != nil {
				return err
			}

			playerAbility.AbilityID = dbAbility.ID

			err = qtx.CreatePlayerAbility(context.Background(), database.CreatePlayerAbilityParams{
				DataHash:            generateDataHash(playerAbility),
				AbilityID:           playerAbility.AbilityID,
				Description:         getNullString(playerAbility.Description),
				Effect:              playerAbility.Effect,
				Topmenu: 			 nullTopmenuType(playerAbility.Topmenu),
				CanUseOutsideBattle: playerAbility.CanUseOutsideBattle,
				MpCost:              getNullInt32(playerAbility.MPCost),
				Cursor:              nullTargetType(playerAbility.Cursor),
			})
			if err != nil {
				return fmt.Errorf("couldn't create Player Ability: %s: %v", ability.Name, err)
			}
		}
		return nil
	})
}
