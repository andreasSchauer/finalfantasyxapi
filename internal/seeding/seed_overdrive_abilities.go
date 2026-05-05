package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func (l *Lookup) loop3SeedOverdriveAbilities(qtx *database.Queries, ctx context.Context) error {
	abilities, err := l.extractOverdriveAbilities()
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
		l.json.overdriveAbilities[i].ID = row.ID
		l.OverdriveAbilities[Key(abilities[i])] = abilities[i]
		l.OverdriveAbilitiesID[row.ID] = abilities[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractOverdriveAbilities() ([]OverdriveAbility, error) {
	abilities := []OverdriveAbility{}
	var err error

	for i := range l.json.overdriveAbilities {
		ability := &l.json.overdriveAbilities[i]

		ability.Ability.ID, err = l.getHashID(ability.Ability)
		if err != nil {
			return nil, err
		}

		abilities = append(abilities, *ability)
	}

	return dedupeRows(abilities, l.Hashes), nil
}

func (l *Lookup) completeOverdriveAbilities() error {
	for i := range l.json.overdriveAbilities {
		ability := &l.json.overdriveAbilities[i]

		err := l.completeBattleInteractions(ability.BattleInteractions)
		if err != nil {
			return err
		}

		l.OverdriveAbilities[Key(ability)] = *ability
		l.OverdriveAbilitiesID[ability.ID] = *ability
	}

	return nil
}

func (l *Lookup) getOverdriveAbilityRelatedStats(oa OverdriveAbility) ([]Stat, error) {
	return getResources(oa.RelatedStats, l.Stats)
}

func (l *Lookup) seedJuncOverdriveAbilitiesRelatedStats(qtx *database.Queries, ctx context.Context) error {
	const desc string = "overdrive abilities + related stats"
	jParams, err := processJunctions(l, desc, l.json.overdriveAbilities, l.getOverdriveAbilityRelatedStats)
	if err != nil {
		return err
	}

	return qtx.CreateOverdriveAbilitiesRelatedStatsJunctionBulk(ctx, database.CreateOverdriveAbilitiesRelatedStatsJunctionBulkParams{
		DataHash:           jParams.DataHashes,
		OverdriveAbilityID: jParams.ParentIDs,
		StatID:             jParams.ChildIDs,
	})
}
