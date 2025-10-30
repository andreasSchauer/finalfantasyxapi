package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

// every slice might make use of a three way junction to be uniquely identified:
// slice item -> BattleInteraction -> Ability
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
		derefOrNil(bi.Range),
		ObjPtrToHashID(bi.Damage),
		derefOrNil(bi.ShatterRate),
		bi.Accuracy.ID,
		bi.HitAmount,
		derefOrNil(bi.SpecialAction),
	}
}

func (bi BattleInteraction) GetID() int32 {
	return bi.ID
}

func (l *lookup) seedBattleInteractions(qtx *database.Queries, ability Ability, battleInteractions []BattleInteraction) error {
	for _, battleInteraction := range battleInteractions {
		junction, err := createJunctionSeed(qtx, ability, battleInteraction, l.seedBattleInteraction)
		if err != nil {
			return fmt.Errorf("ability %s: %v", createLookupKey(ability), err)
		}
		battleInteraction.ID = junction.ChildID

		err = qtx.CreateAbilitiesBattleInteractionsJunction(context.Background(), database.CreateAbilitiesBattleInteractionsJunctionParams{
			DataHash:            generateDataHash(junction),
			AbilityID:           junction.ParentID,
			BattleInteractionID: junction.ChildID,
		})
		if err != nil {
			return fmt.Errorf("ability %s: couldn't create battle interactions: %v", createLookupKey(ability), err)
		}

		err = l.createBattleInteractionRelationships(qtx, ability, battleInteraction)
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *lookup) seedBattleInteraction(qtx *database.Queries, battleInteraction BattleInteraction) (BattleInteraction, error) {
	var err error

	battleInteraction.Damage, err = seedObjPtrAssignFK(qtx, battleInteraction.Damage, l.seedDamage)
	if err != nil {
		return BattleInteraction{}, err
	}

	battleInteraction.Accuracy, err = seedObjAssignID(qtx, battleInteraction.Accuracy, l.seedAccuracy)
	if err != nil {
		return BattleInteraction{}, err
	}

	dbBattleInteraction, err := qtx.CreateBattleInteraction(context.Background(), database.CreateBattleInteractionParams{
		DataHash:          generateDataHash(battleInteraction),
		Target:            database.TargetType(battleInteraction.Target),
		BasedOnPhysAttack: battleInteraction.BasedOnPhysAttack,
		Range:             getNullInt32(battleInteraction.Range),
		DamageID:          ObjPtrToNullInt32ID(battleInteraction.Damage),
		ShatterRate:       getNullInt32(battleInteraction.ShatterRate),
		AccuracyID:        battleInteraction.Accuracy.ID,
		HitAmount:         battleInteraction.HitAmount,
		SpecialAction:     nullSpecialActionType(battleInteraction.SpecialAction),
	})
	if err != nil {
		return BattleInteraction{}, fmt.Errorf("couldn't create battle interaction: %v", err)
	}

	battleInteraction.ID = dbBattleInteraction.ID

	return battleInteraction, nil
}

func (l *lookup) createBattleInteractionRelationships(qtx *database.Queries, ability Ability, battleInteraction BattleInteraction) error {
	functions := []func(*database.Queries, Ability, BattleInteraction) error{
		l.createBattleIntAffectedBy,
		l.createBattleIntInflictedDelay,
		l.createBattleIntInflictedConditions,
		l.createBattleIntRemovedConditions,
		l.createBattleIntCopiedConditions,
		l.createBattleIntStatChanges,
		l.createBattleIntModifierChanges,
	}

	for _, function := range functions {
		err := function(qtx, ability, battleInteraction)
		if err != nil {
			return fmt.Errorf("ability %s: %v", createLookupKey(ability), err)
		}
	}

	return nil
}

func (l *lookup) createBattleIntAffectedBy(qtx *database.Queries, ability Ability, battleInteraction BattleInteraction) error {
	for _, condition := range battleInteraction.AffectedBy {
		threeWay, err := createThreeWayJunction(ability, battleInteraction, condition, l.getStatusCondition)
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
			return fmt.Errorf("couldn't create affected by: %v", err)
		}
	}

	return nil
}

func (l *lookup) createBattleIntInflictedDelay(qtx *database.Queries, ability Ability, battleInteraction BattleInteraction) error {
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
			return fmt.Errorf("couldn't create inflicted delay: %v", err)
		}
	}

	return nil
}

func (l *lookup) createBattleIntInflictedConditions(qtx *database.Queries, ability Ability, battleInteraction BattleInteraction) error {
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
			return fmt.Errorf("couldn't create inflicted status conditions: %v", err)
		}
	}

	return nil
}

func (l *lookup) createBattleIntRemovedConditions(qtx *database.Queries, ability Ability, battleInteraction BattleInteraction) error {
	for _, condition := range battleInteraction.RemovedStatusConditions {
		threeWay, err := createThreeWayJunction(ability, battleInteraction, condition, l.getStatusCondition)
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
			return fmt.Errorf("couldn't create removed status conditions: %v", err)
		}
	}

	return nil
}

func (l *lookup) createBattleIntCopiedConditions(qtx *database.Queries, ability Ability, battleInteraction BattleInteraction) error {
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
			return fmt.Errorf("couldn't create copied status conditions: %v", err)
		}
	}

	return nil
}

func (l *lookup) createBattleIntStatChanges(qtx *database.Queries, ability Ability, battleInteraction BattleInteraction) error {
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
			return fmt.Errorf("couldn't create stat changes: %v", err)
		}
	}

	return nil
}

func (l *lookup) createBattleIntModifierChanges(qtx *database.Queries, ability Ability, battleInteraction BattleInteraction) error {
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
			return fmt.Errorf("couldn't create modifier changes: %v", err)
		}
	}

	return nil
}
