package junctions

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func getAutoAbilityAddedStatusResists(a seedtypes.AutoAbility) ([]seedtypes.StatusResist, error) {
	return a.AddedStatusResists, nil
}

func AutoAbilitiesAddedStatusResists(qtx *database.Queries, ctx context.Context) error {
	const desc string = "auto-abilities + added status resists"
	jParams, err := processJunctions(desc, sh.RawJSON.AutoAbilities, getAutoAbilityAddedStatusResists)
	if err != nil {
		return err
	}

	return qtx.CreateAutoAbilitiesAddedStatusResistsJunctionBulk(ctx, database.CreateAutoAbilitiesAddedStatusResistsJunctionBulkParams{
		DataHash:       jParams.DataHashes,
		AutoAbilityID:  jParams.ParentIDs,
		StatusResistID: jParams.ChildIDs,
	})
}

func getAutoAbilityAddedStatusses(a seedtypes.AutoAbility) ([]seedtypes.StatusCondition, error) {
	return lookup.GetResources(a.AddedStatusses, lookup.L.StatusConditions)
}

func AutoAbilitiesAddedStatusses(qtx *database.Queries, ctx context.Context) error {
	const desc string = "auto-abilities + added statusses"
	jParams, err := processJunctions(desc, sh.RawJSON.AutoAbilities, getAutoAbilityAddedStatusses)
	if err != nil {
		return err
	}

	return qtx.CreateAutoAbilitiesAddedStatussesJunctionBulk(ctx, database.CreateAutoAbilitiesAddedStatussesJunctionBulkParams{
		DataHash:          jParams.DataHashes,
		AutoAbilityID:     jParams.ParentIDs,
		StatusConditionID: jParams.ChildIDs,
	})
}

func getAutoAbilityAutoItems(a seedtypes.AutoAbility) ([]seedtypes.Item, error) {
	return lookup.GetResources(a.AutoItemUse, lookup.L.Items)
}

func AutoAbilitiesAutoItems(qtx *database.Queries, ctx context.Context) error {
	const desc string = "auto-abilities + auto items"
	jParams, err := processJunctions(desc, sh.RawJSON.AutoAbilities, getAutoAbilityAutoItems)
	if err != nil {
		return err
	}

	return qtx.CreateAutoAbilitiesAutoItemJunctionBulk(ctx, database.CreateAutoAbilitiesAutoItemJunctionBulkParams{
		DataHash:      jParams.DataHashes,
		AutoAbilityID: jParams.ParentIDs,
		ItemID:        jParams.ChildIDs,
	})
}

func getAutoAbilityLockedOutAbilities(a seedtypes.AutoAbility) ([]seedtypes.AutoAbility, error) {
	return lookup.GetResources(a.LockedOutAbilities, lookup.L.AutoAbilities)
}

func AutoAbilitiesLockedOutAbilities(qtx *database.Queries, ctx context.Context) error {
	const desc string = "auto-abilities + locked out"
	jParams, err := processJunctions(desc, sh.RawJSON.AutoAbilities, getAutoAbilityLockedOutAbilities)
	if err != nil {
		return err
	}

	return qtx.CreateAutoAbilitiesLockedOutJunctionBulk(ctx, database.CreateAutoAbilitiesLockedOutJunctionBulkParams{
		DataHash:        jParams.DataHashes,
		ParentAbilityID: jParams.ParentIDs,
		ChildAbilityID:  jParams.ChildIDs,
	})
}

func getAutoAbilityModifierChanges(a seedtypes.AutoAbility) ([]seedtypes.ModifierChange, error) {
	return a.ModifierChanges, nil
}

func AutoAbilitiesModifierChanges(qtx *database.Queries, ctx context.Context) error {
	const desc string = "auto-abilities + modifier changes"
	jParams, err := processJunctions(desc, sh.RawJSON.AutoAbilities, getAutoAbilityModifierChanges)
	if err != nil {
		return err
	}

	return qtx.CreateAutoAbilitiesModifierChangesJunctionBulk(ctx, database.CreateAutoAbilitiesModifierChangesJunctionBulkParams{
		DataHash:         jParams.DataHashes,
		AutoAbilityID:    jParams.ParentIDs,
		ModifierChangeID: jParams.ChildIDs,
	})
}

func getAutoAbilityRelatedStats(a seedtypes.AutoAbility) ([]seedtypes.Stat, error) {
	return lookup.GetResources(a.RelatedStats, lookup.L.Stats)
}

func AutoAbilitiesRelatedStats(qtx *database.Queries, ctx context.Context) error {
	const desc string = "auto-abilities + related stats"
	jParams, err := processJunctions(desc, sh.RawJSON.AutoAbilities, getAutoAbilityRelatedStats)
	if err != nil {
		return err
	}

	return qtx.CreateAutoAbilitiesRelatedStatsJunctionBulk(ctx, database.CreateAutoAbilitiesRelatedStatsJunctionBulkParams{
		DataHash:      jParams.DataHashes,
		AutoAbilityID: jParams.ParentIDs,
		StatID:        jParams.ChildIDs,
	})
}

func getAutoAbilityStatChanges(a seedtypes.AutoAbility) ([]seedtypes.StatChange, error) {
	return a.StatChanges, nil
}

func AutoAbilitiesStatChanges(qtx *database.Queries, ctx context.Context) error {
	const desc string = "auto-abilities + stat changes"
	jParams, err := processJunctions(desc, sh.RawJSON.AutoAbilities, getAutoAbilityStatChanges)
	if err != nil {
		return err
	}

	return qtx.CreateAutoAbilitiesStatChangesJunctionBulk(ctx, database.CreateAutoAbilitiesStatChangesJunctionBulkParams{
		DataHash:      jParams.DataHashes,
		AutoAbilityID: jParams.ParentIDs,
		StatChangeID:  jParams.ChildIDs,
	})
}
