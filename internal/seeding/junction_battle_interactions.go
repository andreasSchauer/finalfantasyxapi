package seeding

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func (l *Lookup) getBattleInteractionAffectedBy(bi BattleInteraction) ([]StatusCondition, error) {
	return getResources(bi.AffectedBy, l.StatusConditions)
}

func (l *Lookup) seedJuncBattleInteractionsAffectedBy(qtx *database.Queries, ctx context.Context) error {
	const desc string = "battle interactions + affected by"
	jParams, err := processThreewayJunctions(l, desc, l.getAbilities(), l.getAbilityBattleInteractions, l.getBattleInteractionAffectedBy)
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

func (l *Lookup) getBattleInteractionCopiedStatusConditions(bi BattleInteraction) ([]InflictedStatus, error) {
	return bi.CopiedStatusConditions, nil
}

func (l *Lookup) seedJuncBattleInteractionsCopiedStatusConditions(qtx *database.Queries, ctx context.Context) error {
	const desc string = "battle interactions + copied status conditions"
	jParams, err := processThreewayJunctions(l, desc, l.getAbilities(), l.getAbilityBattleInteractions, l.getBattleInteractionCopiedStatusConditions)
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

func (l *Lookup) getBattleInteractionDamages(bi BattleInteraction) ([]Damage, error) {
	damages := []Damage{}
	if bi.Damage != nil {
		damages = append(damages, *bi.Damage)
	}
	return damages, nil
}

func (l *Lookup) seedJuncBattleInteractionsDamages(qtx *database.Queries, ctx context.Context) error {
	const desc string = "battle interactions + damages"
	jParams, err := processThreewayJunctions(l, desc, l.getAbilities(), l.getAbilityBattleInteractions, l.getBattleInteractionDamages)
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

func (l *Lookup) getBattleInteractionInflictedStatusConditions(bi BattleInteraction) ([]InflictedStatus, error) {
	return bi.InflictedStatusConditions, nil
}

func (l *Lookup) seedJuncBattleInteractionsInflictedStatusConditions(qtx *database.Queries, ctx context.Context) error {
	const desc string = "battle interactions + inflicted status conditions"
	jParams, err := processThreewayJunctions(l, desc, l.getAbilities(), l.getAbilityBattleInteractions, l.getBattleInteractionInflictedStatusConditions)
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

func (l *Lookup) getBattleInteractionModifierChanges(bi BattleInteraction) ([]ModifierChange, error) {
	return bi.ModifierChanges, nil
}

func (l *Lookup) seedJuncBattleInteractionsModifierChanges(qtx *database.Queries, ctx context.Context) error {
	const desc string = "battle interactions + modifier changes"
	jParams, err := processThreewayJunctions(l, desc, l.getAbilities(), l.getAbilityBattleInteractions, l.getBattleInteractionModifierChanges)
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

func (l *Lookup) getBattleInteractionRemovedStatusConditions(bi BattleInteraction) ([]StatusCondition, error) {
	return getResources(bi.RemovedStatusConditions, l.StatusConditions)
}

func (l *Lookup) seedJuncBattleInteractionsRemovedStatusConditions(qtx *database.Queries, ctx context.Context) error {
	const desc string = "battle interactions + removed status conditions"
	jParams, err := processThreewayJunctions(l, desc, l.getAbilities(), l.getAbilityBattleInteractions, l.getBattleInteractionRemovedStatusConditions)
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

func (l *Lookup) getBattleInteractionStatChanges(bi BattleInteraction) ([]StatChange, error) {
	return bi.StatChanges, nil
}

func (l *Lookup) seedJuncBattleInteractionsStatChanges(qtx *database.Queries, ctx context.Context) error {
	const desc string = "battle interactions + stat changes"
	jParams, err := processThreewayJunctions(l, desc, l.getAbilities(), l.getAbilityBattleInteractions, l.getBattleInteractionStatChanges)
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
