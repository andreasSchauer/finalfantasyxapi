package seedhandlers

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Loop3SeedOverdriveAbilities(qtx *database.Queries, ctx context.Context) error {
	abilities, err := extractOverdriveAbilities()
	if err != nil {
		return err
	}

	params := database.CreateOverdriveAbilityBulkParams{
		DataHash:  make([]string, len(abilities)),
		AbilityID: make([]int32, len(abilities)),
	}

	for i, a := range abilities {
		params.DataHash[i] = sh.GenerateDataHash(a)
		params.AbilityID[i] = a.Ability.ID
	}

	dbRows, err := qtx.CreateOverdriveAbilityBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create overdrive abilities: %v", err)
	}

	for i, row := range dbRows {
		abilities[i].ID = row.ID
		sh.RawJSON.OverdriveAbilities[i].ID = row.ID
		lookup.L.OverdriveAbilities[seedtypes.Key(abilities[i])] = abilities[i]
		lookup.L.OverdriveAbilitiesID[row.ID] = abilities[i]
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractOverdriveAbilities() ([]seedtypes.OverdriveAbility, error) {
	abilities := []seedtypes.OverdriveAbility{}
	var err error

	for i := range sh.RawJSON.OverdriveAbilities {
		ability := &sh.RawJSON.OverdriveAbilities[i]

		ability.Ability.ID, err = sh.GetHashID(ability.Ability)
		if err != nil {
			return nil, err
		}

		abilities = append(abilities, *ability)
	}

	return sh.DedupeRows(abilities, sh.Hashes), nil
}

func CompleteOverdriveAbilities() error {
	for i := range sh.RawJSON.OverdriveAbilities {
		ability := &sh.RawJSON.OverdriveAbilities[i]

		err := completeBattleInteractions(ability.BattleInteractions)
		if err != nil {
			return err
		}

		lookup.L.OverdriveAbilities[seedtypes.Key(ability)] = *ability
		lookup.L.OverdriveAbilitiesID[ability.ID] = *ability
	}

	return nil
}
