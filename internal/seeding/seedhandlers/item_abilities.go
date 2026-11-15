package seedhandlers

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Loop3ItemAbilities(qtx *database.Queries, ctx context.Context) error {
	abilities, err := extractItemAbilities()
	if err != nil {
		return err
	}

	params := database.CreateItemAbilityBulkParams{
		DataHash:  make([]string, len(abilities)),
		ItemID:    make([]int32, len(abilities)),
		AbilityID: make([]int32, len(abilities)),
		Cursor:    make([]database.TargetType, len(abilities)),
	}

	for i, a := range abilities {
		params.DataHash[i] = sh.GenerateDataHash(a)
		params.AbilityID[i] = a.Ability.ID
		params.ItemID[i] = a.ItemID
		params.Cursor[i] = database.TargetType(a.Cursor)
	}

	dbRows, err := qtx.CreateItemAbilityBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create item abilities: %v", err)
	}

	for i, row := range dbRows {
		abilities[i].ID = row.ID
		lookup.L.ItemAbilities[seedtypes.Key(abilities[i])] = abilities[i]
		lookup.L.ItemAbilitiesID[row.ID] = abilities[i]
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractItemAbilities() ([]seedtypes.ItemAbility, error) {
	abilities := []seedtypes.ItemAbility{}
	var err error

	for i := range sh.RawJSON.Items {
		item := &sh.RawJSON.Items[i]

		if len(item.BattleInteractions) == 0 {
			continue
		}

		ability := &item.ItemAbility
		ability.Name = item.Name
		ability.Type = database.AbilityTypeItemAbility
		ability.ItemID = item.ID

		ability.Ability.ID, err = sh.GetHashID(ability.Ability)
		if err != nil {
			return nil, err
		}

		abilities = append(abilities, *ability)
	}

	return sh.DedupeRows(abilities, sh.Hashes), nil
}
