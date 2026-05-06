package seeding

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func (l *Lookup) getAlts() []Alt {
	changes := []Alt{}

	for _, mon := range l.json.monsters {
		for _, state := range mon.AlteredStates {
			changes = append(changes, state.Changes...)
		}
	}

	return changes
}

func (l *Lookup) getAltAutoAbilities(c Alt) ([]AutoAbility, error) {
	return getResources(c.AutoAbilities, l.AutoAbilities)
}

func (l *Lookup) seedJuncAltsAutoAbilities(qtx *database.Queries, ctx context.Context) error {
	const desc string = "alts + auto-abilities"
	jParams, err := processJunctions(l, desc, l.getAlts(), l.getAltAutoAbilities)
	if err != nil {
		return err
	}

	return qtx.CreateAltsAutoAbilitiesJunctionBulk(ctx, database.CreateAltsAutoAbilitiesJunctionBulkParams{
		DataHash:         jParams.DataHashes,
		AltID: jParams.ParentIDs,
		AutoAbilityID:    jParams.ChildIDs,
	})
}

func (l *Lookup) getAltBaseStats(c Alt) ([]BaseStat, error) {
	return c.BaseStats, nil
}

func (l *Lookup) seedJuncAltsBaseStats(qtx *database.Queries, ctx context.Context) error {
	const desc string = "alts + base stats"
	jParams, err := processJunctions(l, desc, l.getAlts(), l.getAltBaseStats)
	if err != nil {
		return err
	}

	return qtx.CreateAltsBaseStatsJunctionBulk(ctx, database.CreateAltsBaseStatsJunctionBulkParams{
		DataHash:         jParams.DataHashes,
		AltID: jParams.ParentIDs,
		BaseStatID:       jParams.ChildIDs,
	})
}

func (l *Lookup) getAltElementalResists(c Alt) ([]ElementalResist, error) {
	return c.ElemResists, nil
}

func (l *Lookup) seedJuncAltsElementalResists(qtx *database.Queries, ctx context.Context) error {
	const desc string = "alts + elemental resists"
	jParams, err := processJunctions(l, desc, l.getAlts(), l.getAltElementalResists)
	if err != nil {
		return err
	}

	return qtx.CreateAltsElemResistsJunctionBulk(ctx, database.CreateAltsElemResistsJunctionBulkParams{
		DataHash:         jParams.DataHashes,
		AltID: jParams.ParentIDs,
		ElemResistID:     jParams.ChildIDs,
	})
}

func (l *Lookup) getAltProperties(c Alt) ([]Property, error) {
	return getResources(c.Properties, l.Properties)
}

func (l *Lookup) seedJuncAltsProperties(qtx *database.Queries, ctx context.Context) error {
	const desc string = "alts + properties"
	jParams, err := processJunctions(l, desc, l.getAlts(), l.getAltProperties)
	if err != nil {
		return err
	}

	return qtx.CreateAltsPropertiesJunctionBulk(ctx, database.CreateAltsPropertiesJunctionBulkParams{
		DataHash:         jParams.DataHashes,
		AltID: jParams.ParentIDs,
		PropertyID:       jParams.ChildIDs,
	})
}

func (l *Lookup) getAltStatusImmunities(c Alt) ([]StatusCondition, error) {
	return getResources(c.StatusImmunities, l.StatusConditions)
}

func (l *Lookup) seedJuncAltsStatusImmunities(qtx *database.Queries, ctx context.Context) error {
	const desc string = "alts + status immunities"
	jParams, err := processJunctions(l, desc, l.getAlts(), l.getAltStatusImmunities)
	if err != nil {
		return err
	}

	return qtx.CreateAltsStatusImmunitiesJunctionBulk(ctx, database.CreateAltsStatusImmunitiesJunctionBulkParams{
		DataHash:          jParams.DataHashes,
		AltID:  jParams.ParentIDs,
		StatusConditionID: jParams.ChildIDs,
	})
}
