package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

// might need a three-way junction to combine AbilityDamage with Damage and a specific ability
// maybe even four-way: AbilityDamage -> Damage -> BattleInteraction -> Ability
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

func (l *lookup) seedDamage(qtx *database.Queries, damage Damage) (Damage, error) {
	var err error

	damage.ElementID, err = assignFKPtr(damage.Element, l.getElement)
	if err != nil {
		return Damage{}, err
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
	ability := l.currentAbility
	battleInteraction := l.currentBI

	for _, abilityDamage := range damage.DamageCalc {
		fourWay, err := createFourWayJunctionSeed(qtx, ability, battleInteraction, damage, abilityDamage, l.seedAbilityDamage)
		if err != nil {
			return err
		}

		err = qtx.CreateDamagesDamageCalcJunction(context.Background(), database.CreateDamagesDamageCalcJunctionParams{
			DataHash:        		generateDataHash(fourWay),
			AbilityID:       		fourWay.GreatGrandparentID,
			BattleInteractionID: 	fourWay.GrandparentID,
			DamageID:        		fourWay.ParentID,
			AbilityDamageID: 		fourWay.ChildID,
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
		DataHash:       generateDataHash(abilityDamage),
		Condition:      getNullString(abilityDamage.Condition),
		AttackType:     database.AttackType(abilityDamage.AttackType),
		StatID:         abilityDamage.StatID,
		DamageType:     database.DamageType(abilityDamage.DamageType),
		DamageFormula:  database.DamageFormula(abilityDamage.DamageFormula),
		DamageConstant: abilityDamage.DamageConstant,
	})
	if err != nil {
		return AbilityDamage{}, fmt.Errorf("couldn't create ability damage: %v", err)
	}

	abilityDamage.ID = dbAbilityDamage.ID

	return abilityDamage, nil
}
