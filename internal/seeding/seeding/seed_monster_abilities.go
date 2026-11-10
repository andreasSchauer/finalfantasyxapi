package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func loop3SeedMonsterAbilities(qtx *database.Queries, ctx context.Context) error {
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
		params.DataHash[i] = generateDataHash(a)
		params.AbilityID[i] = a.AbilityID
		params.IsForced[i] = a.IsForced
		params.IsUnused[i] = a.IsUnused
	}

	dbRows, err := qtx.CreateMonsterAbilityBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create monster abilities: %v", err)
	}

	for _, row := range dbRows {
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractMonsterAbilities() ([]seedtypes.MonsterAbility, error) {
	abilities := []seedtypes.MonsterAbility{}
	var err error

	for i := range rawJSON.monsters {
		mon := &rawJSON.monsters[i]

		for j := range mon.Abilities {
			ability := &mon.Abilities[j]

			ability.AbilityID, err = assignFK(ability.AbilityReference, lookup.L.Abilities)
			if err != nil {
				return nil, err
			}

			abilities = append(abilities, *ability)
		}
	}

	return dedupeRows(abilities, hashes), nil
}
