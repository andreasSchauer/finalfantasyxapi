package seeding

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func (l *Lookup) getStatusConditionModifierChanges(sc StatusCondition) ([]ModifierChange, error) {
	return sc.ModifierChanges, nil
}

func (l *Lookup) seedJuncStatusConditionModifierChanges(qtx *database.Queries, ctx context.Context) error {
	const desc string = "status conditions + modifier changes"
	jParams, err := processJunctions(l, desc, l.json.statusConditions, l.getStatusConditionModifierChanges)
	if err != nil {
		return err
	}

	return qtx.CreateStatusConditionsModifierChangesJunctionBulk(ctx, database.CreateStatusConditionsModifierChangesJunctionBulkParams{
		DataHash:          jParams.DataHashes,
		StatusConditionID: jParams.ParentIDs,
		ModifierChangeID:  jParams.ChildIDs,
	})
}

func (l *Lookup) getStatusConditionRelatedStats(sc StatusCondition) ([]Stat, error) {
	return getResources(sc.RelatedStats, l.Stats)
}

func (l *Lookup) seedJuncStatusConditionRelatedStats(qtx *database.Queries, ctx context.Context) error {
	const desc string = "status conditions + related stats"
	jParams, err := processJunctions(l, desc, l.json.statusConditions, l.getStatusConditionRelatedStats)
	if err != nil {
		return err
	}

	return qtx.CreateStatusConditionsRelatedStatsJunctionBulk(ctx, database.CreateStatusConditionsRelatedStatsJunctionBulkParams{
		DataHash:          jParams.DataHashes,
		StatusConditionID: jParams.ParentIDs,
		StatID:            jParams.ChildIDs,
	})
}

func (l *Lookup) getStatusConditionRemovedConditions(sc StatusCondition) ([]StatusCondition, error) {
	return getResources(sc.RemovedStatusConditions, l.StatusConditions)
}

func (l *Lookup) seedJuncStatusConditionRemovedConditions(qtx *database.Queries, ctx context.Context) error {
	const desc string = "status conditions + removed conditions"
	jParams, err := processJunctions(l, desc, l.json.statusConditions, l.getStatusConditionRemovedConditions)
	if err != nil {
		return err
	}

	return qtx.CreateStatusConditionsRemovedStatusConditionsJunctionBulk(ctx, database.CreateStatusConditionsRemovedStatusConditionsJunctionBulkParams{
		DataHash:          jParams.DataHashes,
		ParentConditionID: jParams.ParentIDs,
		ChildConditionID:  jParams.ChildIDs,
	})
}

func (l *Lookup) getStatusConditionStatChanges(sc StatusCondition) ([]StatChange, error) {
	return sc.StatChanges, nil
}

func (l *Lookup) seedJuncStatusConditionStatChanges(qtx *database.Queries, ctx context.Context) error {
	const desc string = "status conditions + stat changes"
	jParams, err := processJunctions(l, desc, l.json.statusConditions, l.getStatusConditionStatChanges)
	if err != nil {
		return err
	}

	return qtx.CreateStatusConditionsStatChangesJunctionBulk(ctx, database.CreateStatusConditionsStatChangesJunctionBulkParams{
		DataHash:          jParams.DataHashes,
		StatusConditionID: jParams.ParentIDs,
		StatChangeID:      jParams.ChildIDs,
	})
}
