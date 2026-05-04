package seeding

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func (l *Lookup) getAutoAbilityAddedStatusResists(a AutoAbility) ([]StatusResist, error) {
	return a.AddedStatusResists, nil
}

func (l *Lookup) seedJuncAutoAbilitiesAddedStatusResists(qtx *database.Queries, ctx context.Context) error {
	const desc string = "auto-abilities + added status resists"
	jParams, err := processJunctions(l, desc, l.json.autoAbilities, l.getAutoAbilityAddedStatusResists)
	if err != nil {
		return err
	}

	return qtx.CreateAutoAbilitiesAddedStatusResistsJunctionBulk(ctx, database.CreateAutoAbilitiesAddedStatusResistsJunctionBulkParams{
		DataHash:       jParams.DataHashes,
		AutoAbilityID:  jParams.ParentIDs,
		StatusResistID: jParams.ChildIDs,
	})
}

func (l *Lookup) getAutoAbilityAddedStatusses(a AutoAbility) ([]StatusCondition, error) {
	return getResources(a.AddedStatusses, l.StatusConditions)
}

func (l *Lookup) seedJuncAutoAbilitiesAddedStatusses(qtx *database.Queries, ctx context.Context) error {
	const desc string = "auto-abilities + added statusses"
	jParams, err := processJunctions(l, desc, l.json.autoAbilities, l.getAutoAbilityAddedStatusses)
	if err != nil {
		return err
	}

	return qtx.CreateAutoAbilitiesAddedStatussesJunctionBulk(ctx, database.CreateAutoAbilitiesAddedStatussesJunctionBulkParams{
		DataHash:          jParams.DataHashes,
		AutoAbilityID:     jParams.ParentIDs,
		StatusConditionID: jParams.ChildIDs,
	})
}

func (l *Lookup) getAutoAbilityAutoItems(a AutoAbility) ([]Item, error) {
	return getResources(a.AutoItemUse, l.Items)
}

func (l *Lookup) seedJuncAutoAbilitiesAutoItems(qtx *database.Queries, ctx context.Context) error {
	const desc string = "auto-abilities + auto items"
	jParams, err := processJunctions(l, desc, l.json.autoAbilities, l.getAutoAbilityAutoItems)
	if err != nil {
		return err
	}

	return qtx.CreateAutoAbilitiesAutoItemJunctionBulk(ctx, database.CreateAutoAbilitiesAutoItemJunctionBulkParams{
		DataHash:      jParams.DataHashes,
		AutoAbilityID: jParams.ParentIDs,
		ItemID:        jParams.ChildIDs,
	})
}

func (l *Lookup) getAutoAbilityLockedOutAbilities(a AutoAbility) ([]AutoAbility, error) {
	return getResources(a.LockedOutAbilities, l.AutoAbilities)
}

func (l *Lookup) seedJuncAutoAbilitiesLockedOutAbilities(qtx *database.Queries, ctx context.Context) error {
	const desc string = "auto-abilities + locked out"
	jParams, err := processJunctions(l, desc, l.json.autoAbilities, l.getAutoAbilityLockedOutAbilities)
	if err != nil {
		return err
	}

	return qtx.CreateAutoAbilitiesLockedOutJunctionBulk(ctx, database.CreateAutoAbilitiesLockedOutJunctionBulkParams{
		DataHash:        jParams.DataHashes,
		ParentAbilityID: jParams.ParentIDs,
		ChildAbilityID:  jParams.ChildIDs,
	})
}

func (l *Lookup) getAutoAbilityModifierChanges(a AutoAbility) ([]ModifierChange, error) {
	return a.ModifierChanges, nil
}

func (l *Lookup) seedJuncAutoAbilitiesModifierChanges(qtx *database.Queries, ctx context.Context) error {
	const desc string = "auto-abilities + modifier changes"
	jParams, err := processJunctions(l, desc, l.json.autoAbilities, l.getAutoAbilityModifierChanges)
	if err != nil {
		return err
	}

	return qtx.CreateAutoAbilitiesModifierChangesJunctionBulk(ctx, database.CreateAutoAbilitiesModifierChangesJunctionBulkParams{
		DataHash:         jParams.DataHashes,
		AutoAbilityID:    jParams.ParentIDs,
		ModifierChangeID: jParams.ChildIDs,
	})
}

func (l *Lookup) getAutoAbilityRelatedStats(a AutoAbility) ([]Stat, error) {
	return getResources(a.RelatedStats, l.Stats)
}

func (l *Lookup) seedJuncAutoAbilitiesRelatedStats(qtx *database.Queries, ctx context.Context) error {
	const desc string = "auto-abilities + related stats"
	jParams, err := processJunctions(l, desc, l.json.autoAbilities, l.getAutoAbilityRelatedStats)
	if err != nil {
		return err
	}

	return qtx.CreateAutoAbilitiesRelatedStatsJunctionBulk(ctx, database.CreateAutoAbilitiesRelatedStatsJunctionBulkParams{
		DataHash:      jParams.DataHashes,
		AutoAbilityID: jParams.ParentIDs,
		StatID:        jParams.ChildIDs,
	})
}

func (l *Lookup) getAutoAbilityStatChanges(a AutoAbility) ([]StatChange, error) {
	return a.StatChanges, nil
}

func (l *Lookup) seedJuncAutoAbilitiesStatChanges(qtx *database.Queries, ctx context.Context) error {
	const desc string = "auto-abilities + stat changes"
	jParams, err := processJunctions(l, desc, l.json.autoAbilities, l.getAutoAbilityStatChanges)
	if err != nil {
		return err
	}

	return qtx.CreateAutoAbilitiesStatChangesJunctionBulk(ctx, database.CreateAutoAbilitiesStatChangesJunctionBulkParams{
		DataHash:      jParams.DataHashes,
		AutoAbilityID: jParams.ParentIDs,
		StatChangeID:  jParams.ChildIDs,
	})
}
