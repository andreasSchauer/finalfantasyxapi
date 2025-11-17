package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Damage struct {
	ID              int32
	DamageCalc      []AbilityDamage `json:"damage_calc"`
	Critical        *string         `json:"critical"`
	CriticalPlusVal *int32          `json:"critical_plus_val"`
	IsPiercing      bool            `json:"is_piercing"`
	BreakDmgLimit   *string         `json:"break_dmg_lmt"`
	ElementID       *int32
	Element         *string `json:"element"`
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

func (d Damage) Error() string {
	return fmt.Sprintf("damage with critical: %v, crit plus: %v, piercing: %t, bdl: %v, element: %v", derefOrNil(d.Critical), derefOrNil(d.CriticalPlusVal), d.IsPiercing, derefOrNil(d.BreakDmgLimit), derefOrNil(d.Element))
}

type AbilityDamage struct {
	ID             int32
	Condition      *string `json:"condition"`
	AttackType     string  `json:"attack_type"`
	StatID         int32
	TargetStat     string `json:"target_stat"`
	DamageType     string `json:"damage_type"`
	DamageFormula  string `json:"damage_formula"`
	DamageConstant int32  `json:"damage_constant"`
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

func (ad AbilityDamage) Error() string {
	return fmt.Sprintf("ability damage with attack type: %s, target stat: %s, damage type: %s, formula: %s, damage constant %d, condition: %v", ad.AttackType, ad.TargetStat, ad.DamageType, ad.DamageFormula, ad.DamageConstant, derefOrNil(ad.Condition))
}

func (l *Lookup) seedDamage(qtx *database.Queries, damage Damage) (Damage, error) {
	var err error

	damage.ElementID, err = assignFKPtr(damage.Element, l.getElement)
	if err != nil {
		return Damage{}, getErr(damage.Error(), err)
	}

	dbDamage, err := qtx.CreateDamage(context.Background(), database.CreateDamageParams{
		DataHash:        generateDataHash(damage),
		Critical:        nullCriticalType(damage.Critical),
		CriticalPlusVal: getNullInt32(damage.CriticalPlusVal),
		IsPiercing:      damage.IsPiercing,
		BreakDmgLimit:   nullBreakDmgLmtType(damage.BreakDmgLimit),
		ElementID:       getNullInt32(damage.ElementID),
	})
	if err != nil {
		return Damage{}, getErr(damage.Error(), err, "couldn't create damage")
	}

	damage.ID = dbDamage.ID

	err = l.seedAbilityDamages(qtx, damage)
	if err != nil {
		return Damage{}, getErr(damage.Error(), err)
	}

	return damage, nil
}

func (l *Lookup) seedAbilityDamages(qtx *database.Queries, damage Damage) error {
	ability := l.currentAbility
	battleInteraction := l.currentBI

	for _, abilityDamage := range damage.DamageCalc {
		fourWay, err := createFourWayJunctionSeed(qtx, ability, battleInteraction, damage, abilityDamage, l.seedAbilityDamage)
		if err != nil {
			return err
		}

		err = qtx.CreateDamagesDamageCalcJunction(context.Background(), database.CreateDamagesDamageCalcJunctionParams{
			DataHash:            generateDataHash(fourWay),
			AbilityID:           fourWay.GreatGrandparentID,
			BattleInteractionID: fourWay.GrandparentID,
			DamageID:            fourWay.ParentID,
			AbilityDamageID:     fourWay.ChildID,
		})
		if err != nil {
			return getErr(abilityDamage.Error(), err, "couldn't junction ability damage")
		}
	}

	return nil
}

func (l *Lookup) seedAbilityDamage(qtx *database.Queries, abilityDamage AbilityDamage) (AbilityDamage, error) {
	var err error

	abilityDamage.StatID, err = assignFK(abilityDamage.TargetStat, l.getStat)
	if err != nil {
		return AbilityDamage{}, getErr(abilityDamage.Error(), err)
	}

	dbAbilityDamage, err := qtx.CreateAbilityDamage(context.Background(), database.CreateAbilityDamageParams{
		DataHash:       generateDataHash(abilityDamage),
		Condition:      getNullString(abilityDamage.Condition),
		AttackType:     database.AttackType(abilityDamage.AttackType),
		StatID:         abilityDamage.StatID,
		DamageType:     database.DamageType(abilityDamage.DamageType),
		DamageFormula:  database.DamageFormula(abilityDamage.DamageFormula),
		DamageConstant: abilityDamage.DamageConstant,
	})
	if err != nil {
		return AbilityDamage{}, getErr(abilityDamage.Error(), err, "couldn't create ability damage")
	}

	abilityDamage.ID = dbAbilityDamage.ID

	return abilityDamage, nil
}
