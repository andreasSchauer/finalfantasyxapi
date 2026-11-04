package seeding

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func getStatusConditionModifierChanges(sc seedtypes.StatusCondition) ([]seedtypes.ModifierChange, error) {
	return sc.ModifierChanges, nil
}

func seedJuncStatusConditionModifierChanges(qtx *database.Queries, ctx context.Context) error {
	const desc string = "status conditions + modifier changes"
	jParams, err := processJunctions(desc, rawJSON.statusConditions, getStatusConditionModifierChanges)
	if err != nil {
		return err
	}

	return qtx.CreateStatusConditionsModifierChangesJunctionBulk(ctx, database.CreateStatusConditionsModifierChangesJunctionBulkParams{
		DataHash:          jParams.DataHashes,
		StatusConditionID: jParams.ParentIDs,
		ModifierChangeID:  jParams.ChildIDs,
	})
}

func getStatusConditionRelatedStats(sc seedtypes.StatusCondition) ([]seedtypes.Stat, error) {
	return lookup.GetResources(sc.RelatedStats, lookup.L.Stats)
}

func seedJuncStatusConditionRelatedStats(qtx *database.Queries, ctx context.Context) error {
	const desc string = "status conditions + related stats"
	jParams, err := processJunctions(desc, rawJSON.statusConditions, getStatusConditionRelatedStats)
	if err != nil {
		return err
	}

	return qtx.CreateStatusConditionsRelatedStatsJunctionBulk(ctx, database.CreateStatusConditionsRelatedStatsJunctionBulkParams{
		DataHash:          jParams.DataHashes,
		StatusConditionID: jParams.ParentIDs,
		StatID:            jParams.ChildIDs,
	})
}

func getStatusConditionRemovedConditions(sc seedtypes.StatusCondition) ([]seedtypes.StatusCondition, error) {
	return lookup.GetResources(sc.RemovedStatusConditions, lookup.L.StatusConditions)
}

func seedJuncStatusConditionRemovedConditions(qtx *database.Queries, ctx context.Context) error {
	const desc string = "status conditions + removed conditions"
	jParams, err := processJunctions(desc, rawJSON.statusConditions, getStatusConditionRemovedConditions)
	if err != nil {
		return err
	}

	return qtx.CreateStatusConditionsRemovedStatusConditionsJunctionBulk(ctx, database.CreateStatusConditionsRemovedStatusConditionsJunctionBulkParams{
		DataHash:          jParams.DataHashes,
		ParentConditionID: jParams.ParentIDs,
		ChildConditionID:  jParams.ChildIDs,
	})
}

func getStatusConditionStatChanges(sc seedtypes.StatusCondition) ([]seedtypes.StatChange, error) {
	return sc.StatChanges, nil
}

func seedJuncStatusConditionStatChanges(qtx *database.Queries, ctx context.Context) error {
	const desc string = "status conditions + stat changes"
	jParams, err := processJunctions(desc, rawJSON.statusConditions, getStatusConditionStatChanges)
	if err != nil {
		return err
	}

	return qtx.CreateStatusConditionsStatChangesJunctionBulk(ctx, database.CreateStatusConditionsStatChangesJunctionBulkParams{
		DataHash:          jParams.DataHashes,
		StatusConditionID: jParams.ParentIDs,
		StatChangeID:      jParams.ChildIDs,
	})
}
