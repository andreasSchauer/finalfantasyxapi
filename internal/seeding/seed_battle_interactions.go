package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type BattleInteraction struct {
	ID                        int32
	Target                    string            `json:"target"`
	BasedOnPhysAttack         bool              `json:"based_on_phys_attack"`
	Range                     *int32            `json:"range"`
	Damage                    *Damage           `json:"damage"`
	ShatterRate               *int32            `json:"shatter_rate"`
	Accuracy                  Accuracy          `json:"accuracy"`
	AffectedBy                []string          `json:"affected_by"`
	HitAmount                 int32             `json:"hit_amount"`
	SpecialAction             *string           `json:"special_action"`
	InflictedDelay            []InflictedDelay  `json:"inflicted_delay"`
	InflictedStatusConditions []InflictedStatus `json:"inflicted_status_conditions"`
	RemovedStatusConditions   []string          `json:"removed_status_conditions"`
	CopiedStatusConditions    []InflictedStatus `json:"copied_status_conditions"`
	StatChanges               []StatChange      `json:"stat_changes"`
	ModifierChanges           []ModifierChange  `json:"modifier_changes"`
}

func (bi BattleInteraction) ToHashFields() []any {
	return []any{
		bi.Target,
		bi.BasedOnPhysAttack,
		h.DerefOrNil(bi.Range),
		h.ObjPtrToID(bi.Damage),
		h.DerefOrNil(bi.ShatterRate),
		bi.Accuracy.ID,
		bi.HitAmount,
		h.DerefOrNil(bi.SpecialAction),
	}
}

func (bi BattleInteraction) GetID() int32 {
	return bi.ID
}

func (bi BattleInteraction) Error() string {
	return fmt.Sprintf("battle interaction with target: %s, phys attack: %t, range: %v, damage id: %v, shatter rate: %v, accuracy id: %d, hit amount: %d, special action: %v", bi.Target, bi.BasedOnPhysAttack, h.DerefOrNil(bi.Range), h.ObjPtrToID(bi.Damage), h.DerefOrNil(bi.ShatterRate), bi.Accuracy.ID, bi.HitAmount, h.DerefOrNil(bi.SpecialAction))
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
			return h.GetErr(battleInteraction.Error(), err, "couldn't junction battle interaction")
		}

		l.currentBI = battleInteraction

		err = l.seedBattleInteractionRelationships(qtx, ability, battleInteraction)
		if err != nil {
			return h.GetErr(battleInteraction.Error(), err)
		}
	}

	return nil
}

func (l *Lookup) seedBattleInteraction(qtx *database.Queries, battleInteraction BattleInteraction) (BattleInteraction, error) {
	var err error

	battleInteraction.Accuracy, err = seedObjAssignID(qtx, battleInteraction.Accuracy, l.seedAccuracy)
	if err != nil {
		return BattleInteraction{}, h.GetErr(battleInteraction.Error(), err)
	}

	dbBattleInteraction, err := qtx.CreateBattleInteraction(context.Background(), database.CreateBattleInteractionParams{
		DataHash:          generateDataHash(battleInteraction),
		Target:            database.TargetType(battleInteraction.Target),
		BasedOnPhysAttack: battleInteraction.BasedOnPhysAttack,
		Range:             h.GetNullInt32(battleInteraction.Range),
		ShatterRate:       h.GetNullInt32(battleInteraction.ShatterRate),
		AccuracyID:        battleInteraction.Accuracy.ID,
		HitAmount:         battleInteraction.HitAmount,
		SpecialAction:     h.NullSpecialActionType(battleInteraction.SpecialAction),
	})
	if err != nil {
		return BattleInteraction{}, h.GetErr(battleInteraction.Error(), err, "couldn't create battle interaction")
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
			return h.GetErr(battleInteraction.Damage.Error(), err, "couldn't junction damage")
		}
	}

	functions := []func(*database.Queries, Ability, BattleInteraction) error{
		l.seedBattleIntAffectedBy,
		l.seedBattleIntInflictedDelay,
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
		threeWay, err := createThreeWayJunction(ability, battleInteraction, conditionString, l.statusConditions)
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
			return h.GetErr(conditionString, err, "couldn't junction affected by status")
		}
	}

	return nil
}

func (l *Lookup) seedBattleIntInflictedDelay(qtx *database.Queries, ability Ability, battleInteraction BattleInteraction) error {
	for _, delay := range battleInteraction.InflictedDelay {
		threeWay, err := createThreeWayJunctionSeed(qtx, ability, battleInteraction, delay, l.seedInflictedDelay)
		if err != nil {
			return err
		}

		err = qtx.CreateBattleIntInflictedDelayJunction(context.Background(), database.CreateBattleIntInflictedDelayJunctionParams{
			DataHash:            generateDataHash(threeWay),
			AbilityID:           threeWay.GrandparentID,
			BattleInteractionID: threeWay.ParentID,
			InflictedDelayID:    threeWay.ChildID,
		})
		if err != nil {
			return h.GetErr(delay.Error(), err, "couldn't junction inflicted delay")
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
			return h.GetErr(condition.Error(), err, "couldn't junction inflicted status condition")
		}
	}

	return nil
}

func (l *Lookup) seedBattleIntRemovedConditions(qtx *database.Queries, ability Ability, battleInteraction BattleInteraction) error {
	for _, conditionString := range battleInteraction.RemovedStatusConditions {
		threeWay, err := createThreeWayJunction(ability, battleInteraction, conditionString, l.statusConditions)
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
			return h.GetErr(conditionString, err, "couldn't junction removed status condition")
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
			return h.GetErr(condition.Error(), err, "couldn't junction copied status condition")
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
			return h.GetErr(statChange.Error(), err, "couldn't junction stat change")
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
			return h.GetErr(modifierChange.Error(), err, "couldn't junction modifier change")
		}
	}

	return nil
}
