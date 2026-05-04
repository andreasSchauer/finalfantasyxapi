package seeding

import (
	"fmt"

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
		fmt.Sprintf("%T", d),
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

func (d *Damage) SetID(id int32) {
	d.ID = id
}

func (d Damage) Error() string {
	return fmt.Sprintf("damage with critical: %v, crit plus: %v, piercing: %t, bdl: %v, element: %v", h.PtrToString(d.Critical), h.PtrToString(d.CriticalPlusVal), d.IsPiercing, h.PtrToString(d.BreakDmgLimit), h.PtrToString(d.Element))
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
		fmt.Sprintf("%T", ad),
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

func (ad *AbilityDamage) SetID(id int32) {
	ad.ID = id
}

func (ad AbilityDamage) Error() string {
	return fmt.Sprintf("ability damage with attack type: %s, target stat: %s, damage type: %s, formula: %s, damage constant %d, condition: %v", ad.AttackType, ad.TargetStat, ad.DamageType, ad.DamageFormula, ad.DamageConstant, h.PtrToString(ad.Condition))
}
