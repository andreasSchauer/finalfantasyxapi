package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func loop3SeedEnemyAbilities(qtx *database.Queries, ctx context.Context) error {
	abilities, err := extractEnemyAbilities()
	if err != nil {
		return err
	}

	params := database.CreateEnemyAbilityBulkParams{
		DataHash:  make([]string, len(abilities)),
		AbilityID: make([]int32, len(abilities)),
		Effect:    make([]sql.NullString, len(abilities)),
	}

	for i, a := range abilities {
		params.DataHash[i] = generateDataHash(a)
		params.AbilityID[i] = a.Ability.ID
		params.Effect[i] = h.GetNullString(a.Effect)
	}

	dbRows, err := qtx.CreateEnemyAbilityBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create enemy abilities: %v", err)
	}

	for i, row := range dbRows {
		abilities[i].ID = row.ID
		rawJSON.enemyAbilities[i].ID = row.ID
		L.EnemyAbilities[seedtypes.Key(abilities[i])] = abilities[i]
		L.EnemyAbilitiesID[row.ID] = abilities[i]
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractEnemyAbilities() ([]seedtypes.EnemyAbility, error) {
	abilities := []seedtypes.EnemyAbility{}
	var err error

	for i := range rawJSON.enemyAbilities {
		ability := &rawJSON.enemyAbilities[i]

		ability.Ability.ID, err = GetHashID(ability.Ability)
		if err != nil {
			return nil, err
		}

		abilities = append(abilities, *ability)
	}

	return dedupeRows(abilities, hashes), nil
}

func completeEnemyAbilities() error {
	for i := range rawJSON.enemyAbilities {
		ability := &rawJSON.enemyAbilities[i]

		err := completeBattleInteractions(ability.BattleInteractions)
		if err != nil {
			return err
		}

		L.EnemyAbilities[seedtypes.Key(ability)] = *ability
		L.EnemyAbilitiesID[ability.ID] = *ability
	}

	return nil
}
