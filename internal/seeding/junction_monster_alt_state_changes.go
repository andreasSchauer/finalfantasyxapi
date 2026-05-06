package seeding

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func (l *Lookup) getAltStateChanges() []Alt {
	changes := []Alt{}

	for _, mon := range l.json.monsters {
		for _, state := range mon.AlteredStates {
			changes = append(changes, state.Changes...)
		}
	}

	return changes
}

func (l *Lookup) getAltStateChangeAutoAbilities(c Alt) ([]AutoAbility, error) {
	return getResources(c.AutoAbilities, l.AutoAbilities)
}

func (l *Lookup) seedJuncAltStateChangesAutoAbilities(qtx *database.Queries, ctx context.Context) error {
	const desc string = "alt state changes + auto-abilities"
	jParams, err := processJunctions(l, desc, l.getAltStateChanges(), l.getAltStateChangeAutoAbilities)
	if err != nil {
		return err
	}

	return qtx.CreateAltStateChangesAutoAbilitiesJunctionBulk(ctx, database.CreateAltStateChangesAutoAbilitiesJunctionBulkParams{
		DataHash:         jParams.DataHashes,
		AltStateChangeID: jParams.ParentIDs,
		AutoAbilityID:    jParams.ChildIDs,
	})
}

func (l *Lookup) getAltStateChangeBaseStats(c Alt) ([]BaseStat, error) {
	return c.BaseStats, nil
}

func (l *Lookup) seedJuncAltStateChangesBaseStats(qtx *database.Queries, ctx context.Context) error {
	const desc string = "alt state changes + base stats"
	jParams, err := processJunctions(l, desc, l.getAltStateChanges(), l.getAltStateChangeBaseStats)
	if err != nil {
		return err
	}

	return qtx.CreateAltStateChangesBaseStatsJunctionBulk(ctx, database.CreateAltStateChangesBaseStatsJunctionBulkParams{
		DataHash:         jParams.DataHashes,
		AltStateChangeID: jParams.ParentIDs,
		BaseStatID:       jParams.ChildIDs,
	})
}

func (l *Lookup) getAltStateChangeElementalResists(c Alt) ([]ElementalResist, error) {
	return c.ElemResists, nil
}

func (l *Lookup) seedJuncAltStateChangesElementalResists(qtx *database.Queries, ctx context.Context) error {
	const desc string = "alt state changes + elemental resists"
	jParams, err := processJunctions(l, desc, l.getAltStateChanges(), l.getAltStateChangeElementalResists)
	if err != nil {
		return err
	}

	return qtx.CreateAltStateChangesElemResistsJunctionBulk(ctx, database.CreateAltStateChangesElemResistsJunctionBulkParams{
		DataHash:         jParams.DataHashes,
		AltStateChangeID: jParams.ParentIDs,
		ElemResistID:     jParams.ChildIDs,
	})
}

func (l *Lookup) getAltStateChangeProperties(c Alt) ([]Property, error) {
	return getResources(c.Properties, l.Properties)
}

func (l *Lookup) seedJuncAltStateChangesProperties(qtx *database.Queries, ctx context.Context) error {
	const desc string = "alt state changes + properties"
	jParams, err := processJunctions(l, desc, l.getAltStateChanges(), l.getAltStateChangeProperties)
	if err != nil {
		return err
	}

	return qtx.CreateAltStateChangesPropertiesJunctionBulk(ctx, database.CreateAltStateChangesPropertiesJunctionBulkParams{
		DataHash:         jParams.DataHashes,
		AltStateChangeID: jParams.ParentIDs,
		PropertyID:       jParams.ChildIDs,
	})
}

func (l *Lookup) getAltStateChangeStatusImmunities(c Alt) ([]StatusCondition, error) {
	return getResources(c.StatusImmunities, l.StatusConditions)
}

func (l *Lookup) seedJuncAltStateChangesStatusImmunities(qtx *database.Queries, ctx context.Context) error {
	const desc string = "alt state changes + status immunities"
	jParams, err := processJunctions(l, desc, l.getAltStateChanges(), l.getAltStateChangeStatusImmunities)
	if err != nil {
		return err
	}

	return qtx.CreateAltStateChangesStatusImmunitiesJunctionBulk(ctx, database.CreateAltStateChangesStatusImmunitiesJunctionBulkParams{
		DataHash:          jParams.DataHashes,
		AltStateChangeID:  jParams.ParentIDs,
		StatusConditionID: jParams.ChildIDs,
	})
}
