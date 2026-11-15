package seedhandlers

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Loop3EnemyAbilities(qtx *database.Queries, ctx context.Context) error {
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
		params.DataHash[i] = sh.GenerateDataHash(a)
		params.AbilityID[i] = a.Ability.ID
		params.Effect[i] = h.GetNullString(a.Effect)
	}

	dbRows, err := qtx.CreateEnemyAbilityBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create enemy abilities: %v", err)
	}

	for i, row := range dbRows {
		abilities[i].ID = row.ID
		sh.RawJSON.EnemyAbilities[i].ID = row.ID
		lookup.L.EnemyAbilities[seedtypes.Key(abilities[i])] = abilities[i]
		lookup.L.EnemyAbilitiesID[row.ID] = abilities[i]
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractEnemyAbilities() ([]seedtypes.EnemyAbility, error) {
	abilities := []seedtypes.EnemyAbility{}
	var err error

	for i := range sh.RawJSON.EnemyAbilities {
		ability := &sh.RawJSON.EnemyAbilities[i]

		ability.Ability.ID, err = sh.GetHashID(ability.Ability)
		if err != nil {
			return nil, err
		}

		abilities = append(abilities, *ability)
	}

	return sh.DedupeRows(abilities, sh.Hashes), nil
}

func CompleteEnemyAbilities() error {
	for i := range sh.RawJSON.EnemyAbilities {
		ability := &sh.RawJSON.EnemyAbilities[i]

		err := completeBattleInteractions(ability.BattleInteractions)
		if err != nil {
			return err
		}

		lookup.L.EnemyAbilities[seedtypes.Key(ability)] = *ability
		lookup.L.EnemyAbilitiesID[ability.ID] = *ability
	}

	return nil
}
