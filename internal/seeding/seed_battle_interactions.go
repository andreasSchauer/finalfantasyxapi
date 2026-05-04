package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type BattleInteraction struct {
	ID                        int32
	Target                    string            `json:"target"`
	BasedOnUserAttack         bool              `json:"based_on_user_attack"`
	Range                     *int32            `json:"range"`
	Damage                    *Damage           `json:"damage"`
	ShatterRate               int32             `json:"shatter_rate"`
	Accuracy                  Accuracy          `json:"accuracy"`
	AffectedBy                []string          `json:"affected_by"`
	HitAmount                 int32             `json:"hit_amount"`
	SpecialAction             *string           `json:"special_action"`
	InflictedDelay            *InflictedDelay   `json:"inflicted_delay"`
	InflictedStatusConditions []InflictedStatus `json:"inflicted_status_conditions"`
	RemovedStatusConditions   []string          `json:"removed_status_conditions"`
	CopiedStatusConditions    []InflictedStatus `json:"copied_status_conditions"`
	StatChanges               []StatChange      `json:"stat_changes"`
	ModifierChanges           []ModifierChange  `json:"modifier_changes"`
}

func (bi BattleInteraction) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", bi),
		bi.Target,
		bi.BasedOnUserAttack,
		h.DerefOrNil(bi.Range),
		bi.ShatterRate,
		bi.Accuracy.ID,
		bi.HitAmount,
		h.DerefOrNil(bi.SpecialAction),
		h.ObjPtrToID(bi.InflictedDelay),
	}
}

func (bi BattleInteraction) GetID() int32 {
	return bi.ID
}

func (bi *BattleInteraction) SetID(id int32) {
	bi.ID = id
}

func (bi BattleInteraction) Error() string {
	return fmt.Sprintf("battle interaction with target: %s, phys attack: %t, range: %v, damage id: %v, shatter rate: %v, accuracy id: %d, hit amount: %d, special action: %v", bi.Target, bi.BasedOnUserAttack, h.PtrToString(bi.Range), h.ObjPtrToID(bi.Damage), bi.ShatterRate, bi.Accuracy.ID, bi.HitAmount, h.PtrToString(bi.SpecialAction))
}

func (l *Lookup) seedBattleInteractions(qtx *database.Queries, ability Ability, battleInteractions []BattleInteraction) error {
	for _, battleInteraction := range battleInteractions {
		junction, err := createJunctionSeed(qtx, ability, battleInteraction, l.seedBattleInteraction)
		if err != nil {
			return err
		}
		battleInteraction.ID = junction.ChildID

		err = qtx.CreateAbilitiesBattleInteractionsJunction(context.Background(), database.CreateAbilitiesBattleInteractionsJunctionParams{
			DataHash:            generateDataHash(junction),
			AbilityID:           junction.ParentID,
			BattleInteractionID: junction.ChildID,
		})
		if err != nil {
			return h.NewErr(battleInteraction.Error(), err, "couldn't junction battle interaction")
		}

		l.currentBI = battleInteraction

		err = l.seedBattleInteractionRelationships(qtx, ability, battleInteraction)
		if err != nil {
			return h.NewErr(battleInteraction.Error(), err)
		}
	}

	return nil
}

func (l *Lookup) seedBattleInteraction(qtx *database.Queries, battleInteraction BattleInteraction) (BattleInteraction, error) {
	var err error

	battleInteraction.Accuracy, err = seedObjAssignID(qtx, battleInteraction.Accuracy, l.seedAccuracy)
	if err != nil {
		return BattleInteraction{}, h.NewErr(battleInteraction.Error(), err)
	}

	battleInteraction.InflictedDelay, err = seedObjPtrAssignFK(qtx, battleInteraction.InflictedDelay, l.seedInflictedDelay)
	if err != nil {
		return BattleInteraction{}, h.NewErr(battleInteraction.Error(), err)
	}

	dbBattleInteraction, err := qtx.CreateBattleInteraction(context.Background(), database.CreateBattleInteractionParams{
		DataHash:          generateDataHash(battleInteraction),
		Target:            database.TargetType(battleInteraction.Target),
		BasedOnUserAttack: battleInteraction.BasedOnUserAttack,
		Range:             h.GetNullInt32(battleInteraction.Range),
		ShatterRate:       battleInteraction.ShatterRate,
		AccuracyID:        battleInteraction.Accuracy.ID,
		InflictedDelayID:  h.ObjPtrToNullInt32ID(battleInteraction.InflictedDelay),
		HitAmount:         battleInteraction.HitAmount,
		SpecialAction:     database.ToNullSpecialActionType(battleInteraction.SpecialAction),
	})
	if err != nil {
		return BattleInteraction{}, h.NewErr(battleInteraction.Error(), err, "couldn't create battle interaction")
	}

	battleInteraction.ID = dbBattleInteraction.ID

	return battleInteraction, nil
}

func (l *Lookup) seedBattleInteractionRelationships(qtx *database.Queries, ability Ability, battleInteraction BattleInteraction) error {
	if battleInteraction.Damage != nil {
		threeWay, err := createThreeWayJunctionSeed(qtx, ability, battleInteraction, *battleInteraction.Damage, l.seedDamage)
		if err != nil {
			return err
		}

		err = qtx.CreateBattleIntDamageJunction(context.Background(), database.CreateBattleIntDamageJunctionParams{
			DataHash:            generateDataHash(threeWay),
			AbilityID:           threeWay.GrandparentID,
			BattleInteractionID: threeWay.ParentID,
			DamageID:            threeWay.ChildID,
		})
		if err != nil {
			return h.NewErr(battleInteraction.Damage.Error(), err, "couldn't junction damage")
		}
	}

	functions := []func(*database.Queries, Ability, BattleInteraction) error{
		l.seedBattleIntAffectedBy,
		l.seedBattleIntInflictedConditions,
		l.seedBattleIntRemovedConditions,
		l.seedBattleIntCopiedConditions,
		l.seedBattleIntStatChanges,
		l.seedBattleIntModifierChanges,
	}

	for _, function := range functions {
		err := function(qtx, ability, battleInteraction)
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *Lookup) seedBattleIntAffectedBy(qtx *database.Queries, ability Ability, battleInteraction BattleInteraction) error {
	for _, conditionString := range battleInteraction.AffectedBy {
		threeWay, err := createThreeWayJunction(ability, battleInteraction, conditionString, l.StatusConditions)
		if err != nil {
			return err
		}

		err = qtx.CreateBattleIntAffectedByJunction(context.Background(), database.CreateBattleIntAffectedByJunctionParams{
			DataHash:            generateDataHash(threeWay),
			AbilityID:           threeWay.GrandparentID,
			BattleInteractionID: threeWay.ParentID,
			StatusConditionID:   threeWay.ChildID,
		})
		if err != nil {
			return h.NewErr(conditionString, err, "couldn't junction affected by status")
		}
	}

	return nil
}

func (l *Lookup) seedBattleIntInflictedConditions(qtx *database.Queries, ability Ability, battleInteraction BattleInteraction) error {
	for _, condition := range battleInteraction.InflictedStatusConditions {
		threeWay, err := createThreeWayJunctionSeed(qtx, ability, battleInteraction, condition, l.seedInflictedStatus)
		if err != nil {
			return err
		}

		err = qtx.CreateBattleIntInflictedConditionsJunction(context.Background(), database.CreateBattleIntInflictedConditionsJunctionParams{
			DataHash:            generateDataHash(threeWay),
			AbilityID:           threeWay.GrandparentID,
			BattleInteractionID: threeWay.ParentID,
			InflictedStatusID:   threeWay.ChildID,
		})
		if err != nil {
			return h.NewErr(condition.Error(), err, "couldn't junction inflicted status condition")
		}
	}

	return nil
}

func (l *Lookup) seedBattleIntRemovedConditions(qtx *database.Queries, ability Ability, battleInteraction BattleInteraction) error {
	for _, conditionString := range battleInteraction.RemovedStatusConditions {
		threeWay, err := createThreeWayJunction(ability, battleInteraction, conditionString, l.StatusConditions)
		if err != nil {
			return err
		}

		err = qtx.CreateBattleIntRemovedConditionsJunction(context.Background(), database.CreateBattleIntRemovedConditionsJunctionParams{
			DataHash:            generateDataHash(threeWay),
			AbilityID:           threeWay.GrandparentID,
			BattleInteractionID: threeWay.ParentID,
			StatusConditionID:   threeWay.ChildID,
		})
		if err != nil {
			return h.NewErr(conditionString, err, "couldn't junction removed status condition")
		}
	}

	return nil
}

func (l *Lookup) seedBattleIntCopiedConditions(qtx *database.Queries, ability Ability, battleInteraction BattleInteraction) error {
	for _, condition := range battleInteraction.CopiedStatusConditions {
		threeWay, err := createThreeWayJunctionSeed(qtx, ability, battleInteraction, condition, l.seedInflictedStatus)
		if err != nil {
			return err
		}

		err = qtx.CreateBattleIntCopiedConditionsJunction(context.Background(), database.CreateBattleIntCopiedConditionsJunctionParams{
			DataHash:            generateDataHash(threeWay),
			AbilityID:           threeWay.GrandparentID,
			BattleInteractionID: threeWay.ParentID,
			InflictedStatusID:   threeWay.ChildID,
		})
		if err != nil {
			return h.NewErr(condition.Error(), err, "couldn't junction copied status condition")
		}
	}

	return nil
}

func (l *Lookup) seedBattleIntStatChanges(qtx *database.Queries, ability Ability, battleInteraction BattleInteraction) error {
	for _, statChange := range battleInteraction.StatChanges {
		threeWay, err := createThreeWayJunctionSeed(qtx, ability, battleInteraction, statChange, l.seedStatChange)
		if err != nil {
			return err
		}

		err = qtx.CreateBattleIntStatChangesJunction(context.Background(), database.CreateBattleIntStatChangesJunctionParams{
			DataHash:            generateDataHash(threeWay),
			AbilityID:           threeWay.GrandparentID,
			BattleInteractionID: threeWay.ParentID,
			StatChangeID:        threeWay.ChildID,
		})
		if err != nil {
			return h.NewErr(statChange.Error(), err, "couldn't junction stat change")
		}
	}

	return nil
}

func (l *Lookup) seedBattleIntModifierChanges(qtx *database.Queries, ability Ability, battleInteraction BattleInteraction) error {
	for _, modifierChange := range battleInteraction.ModifierChanges {
		threeWay, err := createThreeWayJunctionSeed(qtx, ability, battleInteraction, modifierChange, l.seedModifierChange)
		if err != nil {
			return err
		}

		err = qtx.CreateBattleIntModifierChangesJunction(context.Background(), database.CreateBattleIntModifierChangesJunctionParams{
			DataHash:            generateDataHash(threeWay),
			AbilityID:           threeWay.GrandparentID,
			BattleInteractionID: threeWay.ParentID,
			ModifierChangeID:    threeWay.ChildID,
		})
		if err != nil {
			return h.NewErr(modifierChange.Error(), err, "couldn't junction modifier change")
		}
	}

	return nil
}

func (l *Lookup) loop3SeedBattleInteractions(qtx *database.Queries, ctx context.Context) error {
	bis, err := l.extractBattleInteractions()
	if err != nil {
		return err
	}

	params := database.CreateBattleInteractionBulkParams{
		DataHash:          make([]string, len(bis)),
		Target:            make([]database.TargetType, len(bis)),
		BasedOnUserAttack: make([]bool, len(bis)),
		Range:             make([]sql.NullInt32, len(bis)),
		ShatterRate:       make([]int32, len(bis)),
		AccuracyID:        make([]int32, len(bis)),
		InflictedDelayID:  make([]sql.NullInt32, len(bis)),
		HitAmount:         make([]int32, len(bis)),
		SpecialAction:     make([]database.NullSpecialActionType, len(bis)),
	}

	for i, bi := range bis {
		params.DataHash[i] = generateDataHash(bi)
		params.Target[i] = database.TargetType(bi.Target)
		params.BasedOnUserAttack[i] = bi.BasedOnUserAttack
		params.Range[i] = h.GetNullInt32(bi.Range)
		params.ShatterRate[i] = bi.ShatterRate
		params.AccuracyID[i] = bi.Accuracy.ID
		params.InflictedDelayID[i] = h.ObjPtrToNullInt32ID(bi.InflictedDelay)
		params.HitAmount[i] = bi.HitAmount
		params.SpecialAction[i] = database.ToNullSpecialActionType(bi.SpecialAction)
	}

	dbRows, err := qtx.CreateBattleInteractionBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create battle interactions: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractBattleInteractions() ([]BattleInteraction, error) {
	bis := []BattleInteraction{}

	for i := range l.json.playerAbilities {
		ability := &l.json.playerAbilities[i]

		bisNew, err := l.prepareBattleInteractions(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		bis = append(bis, bisNew...)
	}

	for i := range l.json.overdriveAbilities {
		ability := &l.json.overdriveAbilities[i]

		bisNew, err := l.prepareBattleInteractions(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		bis = append(bis, bisNew...)
	}

	for i := range l.json.items {
		item := &l.json.items[i]

		bisNew, err := l.prepareBattleInteractions(item.BattleInteractions)
		if err != nil {
			return nil, err
		}

		bis = append(bis, bisNew...)
	}

	for i := range l.json.triggerCommands {
		command := &l.json.triggerCommands[i]

		bisNew, err := l.prepareBattleInteractions(command.BattleInteractions)
		if err != nil {
			return nil, err
		}

		bis = append(bis, bisNew...)
	}

	for i := range l.json.unspecifiedAbilities {
		ability := &l.json.unspecifiedAbilities[i]

		bisNew, err := l.prepareBattleInteractions(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		bis = append(bis, bisNew...)
	}

	for i := range l.json.enemyAbilities {
		ability := &l.json.enemyAbilities[i]

		bisNew, err := l.prepareBattleInteractions(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		bis = append(bis, bisNew...)
	}

	return dedupeRows(bis, l.Hashes), nil
}

func (l *Lookup) prepareBattleInteractions(bis []BattleInteraction) ([]BattleInteraction, error) {
	bisNew := []BattleInteraction{}
	var err error

	for j := range bis {
		bi := &bis[j]

		bi.Accuracy.ID, err = l.getHashID(bi.Accuracy)
		if err != nil {
			return nil, err
		}

		if bi.InflictedDelay != nil {
			bi.InflictedDelay.ID, err = l.getHashID(bi.InflictedDelay)
		}

		bisNew = append(bisNew, *bi)
	}

	return bisNew, nil
}

func (l *Lookup) completeBattleInteractions(bis []BattleInteraction) error {
	for i := range bis {
		bi := &bis[i]

		err := l.assignID(bi)
		if err != nil {
			return err
		}

		err = l.completeDamage(bi.Damage)
		if err != nil {
			return err
		}

		err = assignIDs(l, bi.InflictedStatusConditions)
		if err != nil {
			return err
		}

		err = assignIDs(l, bi.CopiedStatusConditions)
		if err != nil {
			return err
		}

		err = assignIDs(l, bi.StatChanges)
		if err != nil {
			return err
		}

		err = assignIDs(l, bi.ModifierChanges)
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *Lookup) getBattleInteractionAffectedBy(bi BattleInteraction) ([]StatusCondition, error) {
	return getResources(bi.AffectedBy, l.StatusConditions)
}

func (l *Lookup) getBattleInteractionCopiedStatusConditions(bi BattleInteraction) ([]InflictedStatus, error) {
	return bi.CopiedStatusConditions, nil
}

func (l *Lookup) getBattleInteractionDamages(bi BattleInteraction) ([]Damage, error) {
	damages := []Damage{}
	if bi.Damage != nil {
		damages = append(damages, *bi.Damage)
	}
	return damages, nil
}

func (l *Lookup) getBattleInteractionInflictedStatusConditions(bi BattleInteraction) ([]InflictedStatus, error) {
	return bi.InflictedStatusConditions, nil
}

func (l *Lookup) getBattleInteractionModifierChanges(bi BattleInteraction) ([]ModifierChange, error) {
	return bi.ModifierChanges, nil
}

func (l *Lookup) getBattleInteractionRemovedStatusConditions(bi BattleInteraction) ([]StatusCondition, error) {
	return getResources(bi.RemovedStatusConditions, l.StatusConditions)
}

func (l *Lookup) getBattleInteractionStatChanges(bi BattleInteraction) ([]StatChange, error) {
	return bi.StatChanges, nil
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
