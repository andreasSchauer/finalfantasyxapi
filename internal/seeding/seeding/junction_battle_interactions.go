package seeding

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func getBattleInteractionAffectedBy(bi seedtypes.BattleInteraction) ([]seedtypes.StatusCondition, error) {
	return lookup.GetResources(bi.AffectedBy, lookup.L.StatusConditions)
}

func seedJuncBattleInteractionsAffectedBy(qtx *database.Queries, ctx context.Context) error {
	const desc string = "battle interactions + affected by"
	jParams, err := processThreewayJunctions(desc, getAbilities(), getAbilityBattleInteractions, getBattleInteractionAffectedBy)
	if err != nil {
		return err
	}

	return qtx.CreateBattleIntAffectedByJunctionBulk(ctx, database.CreateBattleIntAffectedByJunctionBulkParams{
		DataHash:            jParams.DataHashes,
		AbilityID:           jParams.GrandParentIDs,
		BattleInteractionID: jParams.ParentIDs,
		StatusConditionID:   jParams.ChildIDs,
	})
}

func getBattleInteractionCopiedStatusConditions(bi seedtypes.BattleInteraction) ([]seedtypes.InflictedStatus, error) {
	return bi.CopiedStatusConditions, nil
}

func seedJuncBattleInteractionsCopiedStatusConditions(qtx *database.Queries, ctx context.Context) error {
	const desc string = "battle interactions + copied status conditions"
	jParams, err := processThreewayJunctions(desc, getAbilities(), getAbilityBattleInteractions, getBattleInteractionCopiedStatusConditions)
	if err != nil {
		return err
	}

	return qtx.CreateBattleIntCopiedConditionsJunctionBulk(ctx, database.CreateBattleIntCopiedConditionsJunctionBulkParams{
		DataHash:            jParams.DataHashes,
		AbilityID:           jParams.GrandParentIDs,
		BattleInteractionID: jParams.ParentIDs,
		InflictedStatusID:   jParams.ChildIDs,
	})
}

func getBattleInteractionDamages(bi seedtypes.BattleInteraction) ([]seedtypes.Damage, error) {
	damages := []seedtypes.Damage{}
	if bi.Damage != nil {
		damages = append(damages, *bi.Damage)
	}
	return damages, nil
}

func seedJuncBattleInteractionsDamages(qtx *database.Queries, ctx context.Context) error {
	const desc string = "battle interactions + damages"
	jParams, err := processThreewayJunctions(desc, getAbilities(), getAbilityBattleInteractions, getBattleInteractionDamages)
	if err != nil {
		return err
	}

	return qtx.CreateBattleIntDamageJunctionBulk(ctx, database.CreateBattleIntDamageJunctionBulkParams{
		DataHash:            jParams.DataHashes,
		AbilityID:           jParams.GrandParentIDs,
		BattleInteractionID: jParams.ParentIDs,
		DamageID:            jParams.ChildIDs,
	})
}

func getBattleInteractionInflictedStatusConditions(bi seedtypes.BattleInteraction) ([]seedtypes.InflictedStatus, error) {
	return bi.InflictedStatusConditions, nil
}

func seedJuncBattleInteractionsInflictedStatusConditions(qtx *database.Queries, ctx context.Context) error {
	const desc string = "battle interactions + inflicted status conditions"
	jParams, err := processThreewayJunctions(desc, getAbilities(), getAbilityBattleInteractions, getBattleInteractionInflictedStatusConditions)
	if err != nil {
		return err
	}

	return qtx.CreateBattleIntInflictedConditionsJunctionBulk(ctx, database.CreateBattleIntInflictedConditionsJunctionBulkParams{
		DataHash:            jParams.DataHashes,
		AbilityID:           jParams.GrandParentIDs,
		BattleInteractionID: jParams.ParentIDs,
		InflictedStatusID:   jParams.ChildIDs,
	})
}

func getBattleInteractionModifierChanges(bi seedtypes.BattleInteraction) ([]seedtypes.ModifierChange, error) {
	return bi.ModifierChanges, nil
}

func seedJuncBattleInteractionsModifierChanges(qtx *database.Queries, ctx context.Context) error {
	const desc string = "battle interactions + modifier changes"
	jParams, err := processThreewayJunctions(desc, getAbilities(), getAbilityBattleInteractions, getBattleInteractionModifierChanges)
	if err != nil {
		return err
	}

	return qtx.CreateBattleIntModifierChangesJunctionBulk(ctx, database.CreateBattleIntModifierChangesJunctionBulkParams{
		DataHash:            jParams.DataHashes,
		AbilityID:           jParams.GrandParentIDs,
		BattleInteractionID: jParams.ParentIDs,
		ModifierChangeID:    jParams.ChildIDs,
	})
}

func getBattleInteractionRemovedStatusConditions(bi seedtypes.BattleInteraction) ([]seedtypes.StatusCondition, error) {
	return lookup.GetResources(bi.RemovedStatusConditions, lookup.L.StatusConditions)
}

func seedJuncBattleInteractionsRemovedStatusConditions(qtx *database.Queries, ctx context.Context) error {
	const desc string = "battle interactions + removed status conditions"
	jParams, err := processThreewayJunctions(desc, getAbilities(), getAbilityBattleInteractions, getBattleInteractionRemovedStatusConditions)
	if err != nil {
		return err
	}

	return qtx.CreateBattleIntRemovedConditionsJunctionBulk(ctx, database.CreateBattleIntRemovedConditionsJunctionBulkParams{
		DataHash:            jParams.DataHashes,
		AbilityID:           jParams.GrandParentIDs,
		BattleInteractionID: jParams.ParentIDs,
		StatusConditionID:   jParams.ChildIDs,
	})
}

func getBattleInteractionStatChanges(bi seedtypes.BattleInteraction) ([]seedtypes.StatChange, error) {
	return bi.StatChanges, nil
}

func seedJuncBattleInteractionsStatChanges(qtx *database.Queries, ctx context.Context) error {
	const desc string = "battle interactions + stat changes"
	jParams, err := processThreewayJunctions(desc, getAbilities(), getAbilityBattleInteractions, getBattleInteractionStatChanges)
	if err != nil {
		return err
	}

	return qtx.CreateBattleIntStatChangesJunctionBulk(ctx, database.CreateBattleIntStatChangesJunctionBulkParams{
		DataHash:            jParams.DataHashes,
		AbilityID:           jParams.GrandParentIDs,
		BattleInteractionID: jParams.ParentIDs,
		StatChangeID:        jParams.ChildIDs,
	})
}
