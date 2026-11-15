package junctions

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func getAlts() []seedtypes.Alt {
	changes := []seedtypes.Alt{}

	for _, mon := range sh.RawJSON.Monsters {
		for _, state := range mon.AlteredStates {
			changes = append(changes, state.Changes...)
		}
	}

	return changes
}

func getAltAutoAbilities(c seedtypes.Alt) ([]seedtypes.AutoAbility, error) {
	return lookup.GetResources(c.AutoAbilities, lookup.L.AutoAbilities)
}

func SeedJuncAltsAutoAbilities(qtx *database.Queries, ctx context.Context) error {
	const desc string = "alts + auto-abilities"
	jParams, err := processJunctions(desc, getAlts(), getAltAutoAbilities)
	if err != nil {
		return err
	}

	return qtx.CreateAltsAutoAbilitiesJunctionBulk(ctx, database.CreateAltsAutoAbilitiesJunctionBulkParams{
		DataHash:      jParams.DataHashes,
		AltID:         jParams.ParentIDs,
		AutoAbilityID: jParams.ChildIDs,
	})
}

func getAltBaseStats(c seedtypes.Alt) ([]seedtypes.BaseStat, error) {
	return c.BaseStats, nil
}

func SeedJuncAltsBaseStats(qtx *database.Queries, ctx context.Context) error {
	const desc string = "alts + base stats"
	jParams, err := processJunctions(desc, getAlts(), getAltBaseStats)
	if err != nil {
		return err
	}

	return qtx.CreateAltsBaseStatsJunctionBulk(ctx, database.CreateAltsBaseStatsJunctionBulkParams{
		DataHash:   jParams.DataHashes,
		AltID:      jParams.ParentIDs,
		BaseStatID: jParams.ChildIDs,
	})
}

func getAltElementalResists(c seedtypes.Alt) ([]seedtypes.ElementalResist, error) {
	return c.ElemResists, nil
}

func SeedJuncAltsElementalResists(qtx *database.Queries, ctx context.Context) error {
	const desc string = "alts + elemental resists"
	jParams, err := processJunctions(desc, getAlts(), getAltElementalResists)
	if err != nil {
		return err
	}

	return qtx.CreateAltsElemResistsJunctionBulk(ctx, database.CreateAltsElemResistsJunctionBulkParams{
		DataHash:     jParams.DataHashes,
		AltID:        jParams.ParentIDs,
		ElemResistID: jParams.ChildIDs,
	})
}

func getAltProperties(c seedtypes.Alt) ([]seedtypes.Property, error) {
	return lookup.GetResources(c.Properties, lookup.L.Properties)
}

func SeedJuncAltsProperties(qtx *database.Queries, ctx context.Context) error {
	const desc string = "alts + properties"
	jParams, err := processJunctions(desc, getAlts(), getAltProperties)
	if err != nil {
		return err
	}

	return qtx.CreateAltsPropertiesJunctionBulk(ctx, database.CreateAltsPropertiesJunctionBulkParams{
		DataHash:   jParams.DataHashes,
		AltID:      jParams.ParentIDs,
		PropertyID: jParams.ChildIDs,
	})
}

func getAltStatusImmunities(c seedtypes.Alt) ([]seedtypes.StatusCondition, error) {
	return lookup.GetResources(c.StatusImmunities, lookup.L.StatusConditions)
}

func SeedJuncAltsStatusImmunities(qtx *database.Queries, ctx context.Context) error {
	const desc string = "alts + status immunities"
	jParams, err := processJunctions(desc, getAlts(), getAltStatusImmunities)
	if err != nil {
		return err
	}

	return qtx.CreateAltsStatusImmunitiesJunctionBulk(ctx, database.CreateAltsStatusImmunitiesJunctionBulkParams{
		DataHash:          jParams.DataHashes,
		AltID:             jParams.ParentIDs,
		StatusConditionID: jParams.ChildIDs,
	})
}
