package seedhandlers

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Loop3SeedMonsterAbilities(qtx *database.Queries, ctx context.Context) error {
	abilities, err := extractMonsterAbilities()
	if err != nil {
		return err
	}

	params := database.CreateMonsterAbilityBulkParams{
		DataHash:  make([]string, len(abilities)),
		AbilityID: make([]int32, len(abilities)),
		IsForced:  make([]bool, len(abilities)),
		IsUnused:  make([]bool, len(abilities)),
	}

	for i, a := range abilities {
		params.DataHash[i] = sh.GenerateDataHash(a)
		params.AbilityID[i] = a.AbilityID
		params.IsForced[i] = a.IsForced
		params.IsUnused[i] = a.IsUnused
	}

	dbRows, err := qtx.CreateMonsterAbilityBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create monster abilities: %v", err)
	}

	for _, row := range dbRows {
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractMonsterAbilities() ([]seedtypes.MonsterAbility, error) {
	abilities := []seedtypes.MonsterAbility{}
	var err error

	for i := range sh.RawJSON.Monsters {
		mon := &sh.RawJSON.Monsters[i]

		for j := range mon.Abilities {
			ability := &mon.Abilities[j]

			ability.AbilityID, err = sh.AssignFK(ability.AbilityReference, lookup.L.Abilities)
			if err != nil {
				return nil, err
			}

			abilities = append(abilities, *ability)
		}
	}

	return sh.DedupeRows(abilities, sh.Hashes), nil
}
