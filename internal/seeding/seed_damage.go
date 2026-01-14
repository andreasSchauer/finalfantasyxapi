package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
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
		h.DerefOrNil(d.Critical),
		h.DerefOrNil(d.CriticalPlusVal),
		d.IsPiercing,
		h.DerefOrNil(d.BreakDmgLimit),
		h.DerefOrNil(d.ElementID),
	}
}

func (d Damage) GetID() int32 {
	return d.ID
}

func (d Damage) Error() string {
	return fmt.Sprintf("damage with critical: %v, crit plus: %v, piercing: %t, bdl: %v, element: %v", h.DerefOrNil(d.Critical), h.DerefOrNil(d.CriticalPlusVal), d.IsPiercing, h.DerefOrNil(d.BreakDmgLimit), h.DerefOrNil(d.Element))
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
		h.DerefOrNil(ad.Condition),
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
	return fmt.Sprintf("ability damage with attack type: %s, target stat: %s, damage type: %s, formula: %s, damage constant %d, condition: %v", ad.AttackType, ad.TargetStat, ad.DamageType, ad.DamageFormula, ad.DamageConstant, h.DerefOrNil(ad.Condition))
}

func (l *Lookup) seedDamage(qtx *database.Queries, damage Damage) (Damage, error) {
	var err error

	damage.ElementID, err = assignFKPtr(damage.Element, l.Elements)
	if err != nil {
		return Damage{}, h.NewErr(damage.Error(), err)
	}

	dbDamage, err := qtx.CreateDamage(context.Background(), database.CreateDamageParams{
		DataHash:        generateDataHash(damage),
		Critical:        h.NullCriticalType(damage.Critical),
		CriticalPlusVal: h.GetNullInt32(damage.CriticalPlusVal),
		IsPiercing:      damage.IsPiercing,
		BreakDmgLimit:   h.NullBreakDmgLmtType(damage.BreakDmgLimit),
		ElementID:       h.GetNullInt32(damage.ElementID),
	})
	if err != nil {
		return Damage{}, h.NewErr(damage.Error(), err, "couldn't create damage")
	}

	damage.ID = dbDamage.ID

	err = l.seedAbilityDamages(qtx, damage)
	if err != nil {
		return Damage{}, h.NewErr(damage.Error(), err)
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
			return h.NewErr(abilityDamage.Error(), err, "couldn't junction ability damage")
		}
	}

	return nil
}

func (l *Lookup) seedAbilityDamage(qtx *database.Queries, abilityDamage AbilityDamage) (AbilityDamage, error) {
	var err error

	abilityDamage.StatID, err = assignFK(abilityDamage.TargetStat, l.Stats)
	if err != nil {
		return AbilityDamage{}, h.NewErr(abilityDamage.Error(), err)
	}

	dbAbilityDamage, err := qtx.CreateAbilityDamage(context.Background(), database.CreateAbilityDamageParams{
		DataHash:       generateDataHash(abilityDamage),
		Condition:      h.GetNullString(abilityDamage.Condition),
		AttackType:     database.AttackType(abilityDamage.AttackType),
		StatID:         abilityDamage.StatID,
		DamageType:     database.DamageType(abilityDamage.DamageType),
		DamageFormula:  database.DamageFormula(abilityDamage.DamageFormula),
		DamageConstant: abilityDamage.DamageConstant,
	})
	if err != nil {
		return AbilityDamage{}, h.NewErr(abilityDamage.Error(), err, "couldn't create ability damage")
	}

	abilityDamage.ID = dbAbilityDamage.ID

	return abilityDamage, nil
}
