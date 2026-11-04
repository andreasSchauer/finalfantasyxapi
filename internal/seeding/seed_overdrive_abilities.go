package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func loop3SeedOverdriveAbilities(qtx *database.Queries, ctx context.Context) error {
	abilities, err := extractOverdriveAbilities()
	if err != nil {
		return err
	}

	params := database.CreateOverdriveAbilityBulkParams{
		DataHash:  make([]string, len(abilities)),
		AbilityID: make([]int32, len(abilities)),
	}

	for i, a := range abilities {
		params.DataHash[i] = generateDataHash(a)
		params.AbilityID[i] = a.Ability.ID
	}

	dbRows, err := qtx.CreateOverdriveAbilityBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create overdrive abilities: %v", err)
	}

	for i, row := range dbRows {
		abilities[i].ID = row.ID
		rawJSON.overdriveAbilities[i].ID = row.ID
		lookup.L.OverdriveAbilities[seedtypes.Key(abilities[i])] = abilities[i]
		lookup.L.OverdriveAbilitiesID[row.ID] = abilities[i]
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractOverdriveAbilities() ([]seedtypes.OverdriveAbility, error) {
	abilities := []seedtypes.OverdriveAbility{}
	var err error

	for i := range rawJSON.overdriveAbilities {
		ability := &rawJSON.overdriveAbilities[i]

		ability.Ability.ID, err = GetHashID(ability.Ability)
		if err != nil {
			return nil, err
		}

		abilities = append(abilities, *ability)
	}

	return dedupeRows(abilities, hashes), nil
}

func completeOverdriveAbilities() error {
	for i := range rawJSON.overdriveAbilities {
		ability := &rawJSON.overdriveAbilities[i]

		err := completeBattleInteractions(ability.BattleInteractions)
		if err != nil {
			return err
		}

		lookup.L.OverdriveAbilities[seedtypes.Key(ability)] = *ability
		lookup.L.OverdriveAbilitiesID[ability.ID] = *ability
	}

	return nil
}

func getOverdriveAbilityRelatedStats(oa seedtypes.OverdriveAbility) ([]seedtypes.Stat, error) {
	return lookup.GetResources(oa.RelatedStats, lookup.L.Stats)
}

func seedJuncOverdriveAbilitiesRelatedStats(qtx *database.Queries, ctx context.Context) error {
	const desc string = "overdrive abilities + related stats"
	jParams, err := processJunctions(desc, rawJSON.overdriveAbilities, getOverdriveAbilityRelatedStats)
	if err != nil {
		return err
	}

	return qtx.CreateOverdriveAbilitiesRelatedStatsJunctionBulk(ctx, database.CreateOverdriveAbilitiesRelatedStatsJunctionBulkParams{
		DataHash:           jParams.DataHashes,
		OverdriveAbilityID: jParams.ParentIDs,
		StatID:             jParams.ChildIDs,
	})
}
