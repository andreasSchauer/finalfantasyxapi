package junctions

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func getStatusConditionModifierChanges(sc seedtypes.StatusCondition) ([]seedtypes.ModifierChange, error) {
	return sc.ModifierChanges, nil
}

func StatusConditionModifierChanges(qtx *database.Queries, ctx context.Context) error {
	const desc string = "status conditions + modifier changes"
	jParams, err := processJunctions(desc, sh.RawJSON.StatusConditions, getStatusConditionModifierChanges)
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

func StatusConditionRelatedStats(qtx *database.Queries, ctx context.Context) error {
	const desc string = "status conditions + related stats"
	jParams, err := processJunctions(desc, sh.RawJSON.StatusConditions, getStatusConditionRelatedStats)
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

func StatusConditionRemovedConditions(qtx *database.Queries, ctx context.Context) error {
	const desc string = "status conditions + removed conditions"
	jParams, err := processJunctions(desc, sh.RawJSON.StatusConditions, getStatusConditionRemovedConditions)
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

func StatusConditionStatChanges(qtx *database.Queries, ctx context.Context) error {
	const desc string = "status conditions + stat changes"
	jParams, err := processJunctions(desc, sh.RawJSON.StatusConditions, getStatusConditionStatChanges)
	if err != nil {
		return err
	}

	return qtx.CreateStatusConditionsStatChangesJunctionBulk(ctx, database.CreateStatusConditionsStatChangesJunctionBulkParams{
		DataHash:          jParams.DataHashes,
		StatusConditionID: jParams.ParentIDs,
		StatChangeID:      jParams.ChildIDs,
	})
}
