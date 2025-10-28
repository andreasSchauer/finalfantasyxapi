package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

// every slice might make use of a three way junction to be uniquely identified:
// slice item -> BattleInteraction -> Ability
type BattleInteraction struct {
	ID							int32
	Target						string				`json:"target"`
	BasedOnPhysAttack			bool				`json:"based_on_phys_attack"`
	Range						*int32				`json:"range"`
	Damage						*Damage				`json:"damage"`
	ShatterRate					*int32				`json:"shatter_rate"`
	Accuracy					Accuracy			`json:"accuracy"`
	AffectedBy					[]string			`json:"affected_by"`
	HitAmount					int32				`json:"hit_amount"`
	SpecialAction				*string				`json:"special_action"`
	InflictedDelay				[]InflictedDelay	`json:"inflicted_delay"`
	InflictedStatusConditions	[]InflictedStatus	`json:"inflicted_status_conditions"`
	RemovedStatusConditions		[]string			`json:"removed_status_conditions"`
	CopiedStatusConditions		[]InflictedStatus	`json:"copied_status_conditions"`
	StatChanges					[]StatChange		`json:"stat_changes"`
	ModifierChanges				[]ModifierChange	`json:"modifier_changes"`
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

// might need a three-way junction to combine AbilityDamage with Damage and a specific ability
// maybe even four-way: AbilityDamage -> Damage -> BattleInteraction -> Ability
type Damage struct {
	ID					int32
	DamageCalc			[]AbilityDamage		`json:"damage_calc"`
	Critical			*string				`json:"critical"`
	CriticalPlusVal		*int32				`json:"critical_plus_val"`
	IsPiercing			bool				`json:"is_piercing"`
	BreakDmgLimit		*string				`json:"break_dmg_lmt"`
	ElementID			*int32
	Element				*string				`json:"element"`
}

func (d Damage) ToHashFields() []any {
	return []any{
		derefOrNil(d.Critical),
		derefOrNil(d.CriticalPlusVal),
		d.IsPiercing,
		derefOrNil(d.BreakDmgLimit),
		derefOrNil(d.ElementID),
	}
}

func (d Damage) GetID() int32 {
	return d.ID
}


type AbilityDamage struct {
	ID					int32
	Condition			*string		`json:"condition"`
	AttackType			string		`json:"attack_type"`
	StatID				int32
	TargetStat			string		`json:"target_stat"`
	DamageType			string		`json:"damage_type"`
	DamageFormula		string		`json:"damage_formula"`
	DamageConstant		int32		`json:"damage_constant"`
}

func (ad AbilityDamage) ToHashFields() []any {
	return []any{
		derefOrNil(ad.Condition),
		ad.AttackType,
		ad.StatID,
		ad.DamageType,
		ad.DamageFormula,
		ad.DamageConstant,
	}
}

func (ad AbilityDamage) GetID() int32 {
	return ad.ID
}


type Accuracy struct {
	ID					int32
	AccSource			string		`json:"acc_source"`
	HitChance			int32		`json:"hit_chance"`
	AccModifier			*float32	`json:"acc_modifier"`
}

func (a Accuracy) ToHashFields() []any {
	return []any{
		a.AccSource,
		a.HitChance,
		derefOrNil(a.AccModifier),
	}
}

func (a Accuracy) GetID() int32 {
	return a.ID
}


type InflictedDelay struct {
	ID					int32
	Condition			*string		`json:"condition"`
	CTBAttackType		string		`json:"ctb_attack_type"`
	DelayType			string		`json:"delay_type"`
	DamageConstant		int32		`json:"damage_constant"`
}

func (id InflictedDelay) ToHashFields() []any {
	return []any{
		derefOrNil(id.Condition),
		id.CTBAttackType,
		id.DelayType,
		id.DamageConstant,
	}
}


func (id InflictedDelay) GetID() int32 {
	return id.ID
}


type ThreeWayJunction struct {
	GrandparentID	int32
	Junction
}


func (j ThreeWayJunction) ToHashFields() []any {
	return []any{
		j.GrandparentID,
		j.ParentID,
		j.ChildID,
	}
}


func (l *lookup) seedBattleInteractions(qtx *database.Queries, ability Ability, battleInteractions []BattleInteraction) error {
	for _, battleInteraction := range battleInteractions {
		junction, err := createJunctionSeed(qtx, ability, battleInteraction, l.seedBattleInteraction)
		if err != nil {
			return fmt.Errorf("ability %s: %v", createLookupKey(ability), err)
		}
		battleInteraction.ID = junction.ChildID
		
		err = qtx.CreateAbilitiesBattleInteractionsJunction(context.Background(), database.CreateAbilitiesBattleInteractionsJunctionParams{
			DataHash: 				generateDataHash(junction),
			AbilityID: 				junction.ParentID,
			BattleInteractionID: 	junction.ChildID,
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
		DataHash: generateDataHash(battleInteraction),
		Target: database.TargetType(battleInteraction.Target),
		BasedOnPhysAttack: battleInteraction.BasedOnPhysAttack,
		Range: getNullInt32(battleInteraction.Range),
		DamageID: ObjPtrToNullInt32ID(battleInteraction.Damage),
		ShatterRate: getNullInt32(battleInteraction.ShatterRate),
		AccuracyID: battleInteraction.Accuracy.ID,
		HitAmount: battleInteraction.HitAmount,
		SpecialAction: nullSpecialActionType(battleInteraction.SpecialAction),
	})
	if err != nil {
		return BattleInteraction{}, fmt.Errorf("couldn't create battle interaction: %v", err)
	}
	
	battleInteraction.ID = dbBattleInteraction.ID

	return battleInteraction, nil
}


func (l *lookup) createBattleInteractionRelationships(qtx *database.Queries, ability Ability, battleInteraction BattleInteraction) error {
	functions := []func(*database.Queries, Ability, BattleInteraction) error{
		l.createBattleInteractionAffectedBy,
		l.createBattleInteractionInflictedDelay,
		l.createBattleInteractionInflictedStatusConditions,
		l.createBattleInteractionRemovedStatusConditions,
		l.createBattleInteractionCopiedStatusConditions,
		l.createBattleInteractionStatChanges,
		l.createBattleInteractionModifierChanges,
	}

	for _, function := range functions {
		err := function(qtx, ability, battleInteraction)
		if err != nil {
			return fmt.Errorf("ability %s: %v", createLookupKey(ability), err)
		}
	}

	return nil
}


func (l *lookup) createBattleInteractionAffectedBy(qtx *database.Queries, ability Ability, battleInteraction BattleInteraction) error {
	for _, condition := range battleInteraction.AffectedBy {
		var err error
		threeWay := ThreeWayJunction{}
		threeWay.GrandparentID = ability.ID

		threeWay.Junction, err = createJunction(battleInteraction, condition, l.getStatusCondition)
		if err != nil {
			return err
		}

		err = qtx.CreateBattleInteractionsAffectedByJunction(context.Background(), database.CreateBattleInteractionsAffectedByJunctionParams{
			DataHash: 				generateDataHash(threeWay),
			AbilityID: 				threeWay.GrandparentID,
			BattleInteractionID: 	threeWay.ParentID,
			StatusConditionID: 		threeWay.ChildID,
		})
		if err != nil {
			return fmt.Errorf("couldn't create affected by: %v", err)
		}
	}

	return nil
}


func (l *lookup) createBattleInteractionInflictedDelay(qtx *database.Queries, ability Ability, battleInteraction BattleInteraction) error {
	for _, delay := range battleInteraction.InflictedDelay {
		var err error
		threeWay := ThreeWayJunction{}
		threeWay.GrandparentID = ability.ID

		threeWay.Junction, err = createJunctionSeed(qtx, battleInteraction, delay, l.seedInflictedDelay)
		if err != nil {
			return err
		}

		err = qtx.CreateBattleInteractionsInflictedDelayJunction(context.Background(), database.CreateBattleInteractionsInflictedDelayJunctionParams{
			DataHash: 				generateDataHash(threeWay),
			AbilityID: 				threeWay.GrandparentID,
			BattleInteractionID: 	threeWay.ParentID,
			InflictedDelayID: 		threeWay.ChildID,
		})
		if err != nil {
			return fmt.Errorf("couldn't create inflicted delay: %v", err)
		}
	}

	return nil
}


func (l *lookup) createBattleInteractionInflictedStatusConditions(qtx *database.Queries, ability Ability, battleInteraction BattleInteraction) error {
	for _, condition := range battleInteraction.InflictedStatusConditions {
		var err error
		threeWay := ThreeWayJunction{}
		threeWay.GrandparentID = ability.ID

		threeWay.Junction, err = createJunctionSeed(qtx, battleInteraction, condition, l.seedInflictedStatus)
		if err != nil {
			return err
		}

		err = qtx.CreateBattleInteractionsInflictedStatusConditionsJunction(context.Background(), database.CreateBattleInteractionsInflictedStatusConditionsJunctionParams{
			DataHash: 				generateDataHash(threeWay),
			AbilityID: 				threeWay.GrandparentID,
			BattleInteractionID: 	threeWay.ParentID,
			InflictedStatusID: 		threeWay.ChildID,
		})
		if err != nil {
			return fmt.Errorf("couldn't create inflicted status conditions: %v", err)
		}
	}

	return nil
}


func (l *lookup) createBattleInteractionRemovedStatusConditions(qtx *database.Queries, ability Ability, battleInteraction BattleInteraction) error {
	for _, condition := range battleInteraction.RemovedStatusConditions {
		var err error
		threeWay := ThreeWayJunction{}
		threeWay.GrandparentID = ability.ID

		threeWay.Junction, err = createJunction(battleInteraction, condition, l.getStatusCondition)
		if err != nil {
			return err
		}

		err = qtx.CreateBattleInteractionsRemovedStatusConditionsJunction(context.Background(), database.CreateBattleInteractionsRemovedStatusConditionsJunctionParams{
			DataHash: 				generateDataHash(threeWay),
			AbilityID: 				threeWay.GrandparentID,
			BattleInteractionID: 	threeWay.ParentID,
			StatusConditionID: 		threeWay.ChildID,
		})
		if err != nil {
			return fmt.Errorf("couldn't create removed status conditions: %v", err)
		}
	}

	return nil
}


func (l *lookup) createBattleInteractionCopiedStatusConditions(qtx *database.Queries, ability Ability, battleInteraction BattleInteraction) error {
	for _, condition := range battleInteraction.CopiedStatusConditions {
		var err error
		threeWay := ThreeWayJunction{}
		threeWay.GrandparentID = ability.ID

		threeWay.Junction, err = createJunctionSeed(qtx, battleInteraction, condition, l.seedInflictedStatus)
		if err != nil {
			return err
		}

		err = qtx.CreateBattleInteractionsCopiedStatusConditionsJunction(context.Background(), database.CreateBattleInteractionsCopiedStatusConditionsJunctionParams{
			DataHash: 				generateDataHash(threeWay),
			AbilityID: 				threeWay.GrandparentID,
			BattleInteractionID: 	threeWay.ParentID,
			InflictedStatusID: 		threeWay.ChildID,
		})
		if err != nil {
			return fmt.Errorf("couldn't create copied status conditions: %v", err)
		}
	}

	return nil
}


func (l *lookup) createBattleInteractionStatChanges(qtx *database.Queries, ability Ability, battleInteraction BattleInteraction) error {
	for _, statChange := range battleInteraction.StatChanges {
		var err error
		threeWay := ThreeWayJunction{}
		threeWay.GrandparentID = ability.ID

		threeWay.Junction, err = createJunctionSeed(qtx, battleInteraction, statChange, l.seedStatChange)
		if err != nil {
			return err
		}

		err = qtx.CreateBattleInteractionsStatChangesJunction(context.Background(), database.CreateBattleInteractionsStatChangesJunctionParams{
			DataHash: 				generateDataHash(threeWay),
			AbilityID: 				threeWay.GrandparentID,
			BattleInteractionID: 	threeWay.ParentID,
			StatChangeID: 			threeWay.ChildID,
		})
		if err != nil {
			return fmt.Errorf("couldn't create stat changes: %v", err)
		}
	}

	return nil
}


func (l *lookup) createBattleInteractionModifierChanges(qtx *database.Queries, ability Ability, battleInteraction BattleInteraction) error {
	for _, modifierChange := range battleInteraction.ModifierChanges {
		var err error
		threeWay := ThreeWayJunction{}
		threeWay.GrandparentID = ability.ID

		threeWay.Junction, err = createJunctionSeed(qtx, battleInteraction, modifierChange, l.seedModifierChange)
		if err != nil {
			return err
		}

		err = qtx.CreateBattleInteractionsModifierChangesJunction(context.Background(), database.CreateBattleInteractionsModifierChangesJunctionParams{
			DataHash: 				generateDataHash(threeWay),
			AbilityID: 				threeWay.GrandparentID,
			BattleInteractionID: 	threeWay.ParentID,
			ModifierChangeID: 		threeWay.ChildID,
		})
		if err != nil {
			return fmt.Errorf("couldn't create modifier changes: %v", err)
		}
	}

	return nil
}


func (l *lookup) seedDamage(qtx *database.Queries, damage Damage) (Damage, error) {
	var err error

	damage.ElementID, err = assignFKPtr(damage.Element, l.getElement)
	if err != nil {
		return Damage{}, err
	}

	dbDamage, err := qtx.CreateDamage(context.Background(), database.CreateDamageParams{
		DataHash: 			generateDataHash(damage),
		Critical: 			nullCriticalType(damage.Critical),
		CriticalPlusVal: 	getNullInt32(damage.CriticalPlusVal),
		IsPiercing: 		damage.IsPiercing,
		BreakDmgLimit: 		nullBreakDmgLmtType(damage.BreakDmgLimit),
		ElementID: 			getNullInt32(damage.ElementID),
	})
	if err != nil {
		return Damage{}, fmt.Errorf("couldn't create damage: %v", err)
	}

	damage.ID = dbDamage.ID

	err = l.seedAbilityDamages(qtx, damage)
	if err != nil {
		return Damage{}, err
	}

	return damage, nil
}


func (l *lookup) seedAbilityDamages(qtx *database.Queries, damage Damage) error {
	for _, abilityDamage := range damage.DamageCalc {
		junction, err := createJunctionSeed(qtx, damage, abilityDamage, l.seedAbilityDamage)
		if err != nil {
			return err
		}

		err = qtx.CreateDamagesDamageCalcJunction(context.Background(), database.CreateDamagesDamageCalcJunctionParams{
			DataHash: 			generateDataHash(junction),
			DamageID: 			junction.ParentID,
			AbilityDamageID: 	junction.ChildID,
		})
		if err != nil {
			return fmt.Errorf("couldn't create damage junction: %v", err)
		}
	}

	return nil
}


func (l *lookup) seedAbilityDamage(qtx *database.Queries, abilityDamage AbilityDamage) (AbilityDamage, error) {
	var err error

	abilityDamage.StatID, err = assignFK(abilityDamage.TargetStat, l.getStat)
	if err != nil {
		return AbilityDamage{}, err
	}

	dbAbilityDamage, err := qtx.CreateAbilityDamage(context.Background(), database.CreateAbilityDamageParams{
		DataHash: 		generateDataHash(abilityDamage),
		Condition: 		getNullString(abilityDamage.Condition),
		AttackType: 	database.AttackType(abilityDamage.AttackType),
		StatID: 		abilityDamage.StatID,
		DamageType: 	database.DamageType(abilityDamage.DamageType),
		DamageFormula: 	database.DamageFormula(abilityDamage.DamageFormula),
		DamageConstant: abilityDamage.DamageConstant,
	})
	if err != nil {
		return AbilityDamage{}, fmt.Errorf("couldn't create ability damage: %v", err)
	}

	abilityDamage.ID = dbAbilityDamage.ID

	return abilityDamage, nil
}


func (l *lookup) seedAccuracy(qtx *database.Queries, accuracy Accuracy) (Accuracy, error) {
	dbAccuracy, err := qtx.CreateAbilityAccuracy(context.Background(), database.CreateAbilityAccuracyParams{
		DataHash: 		generateDataHash(accuracy),
		AccSource: 		database.AccSourceType(accuracy.AccSource),
		HitChance: 		accuracy.HitChance,
		AccModifier: 	getNullFloat64(accuracy.AccModifier),
	})
	if err != nil {
		return Accuracy{}, fmt.Errorf("couldn't create accuracy: %v", err)
	}

	accuracy.ID = dbAccuracy.ID

	return accuracy, nil
}


func (l *lookup) seedInflictedDelay(qtx *database.Queries, delay InflictedDelay) (InflictedDelay, error) {
	dbDelay, err := qtx.CreateInflictedDelay(context.Background(), database.CreateInflictedDelayParams{
		DataHash: 		generateDataHash(delay),
		Condition: 		getNullString(delay.Condition),
		CtbAttackType: 	database.CtbAttackType(delay.CTBAttackType),
		DelayType: 		database.DelayType(delay.DelayType),
		DamageConstant: delay.DamageConstant,
	})
	if err != nil {
		return InflictedDelay{}, fmt.Errorf("couldn't create inflicted delay: %v", err)
	}

	delay.ID = dbDelay.ID

	return delay, nil
}