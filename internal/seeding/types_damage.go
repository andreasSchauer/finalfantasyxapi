package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Damage struct {
	ID              int32        `json:"damage_id,omitempty"`
	AttackType      string       `json:"attack_type"`
	DamageType      string       `json:"damage_type"`
	HP              *ClassDamage `json:"hp"`
	MP              *ClassDamage `json:"mp"`
	CTB             *ClassDamage `json:"ctb"`
	Critical        *string      `json:"critical"`
	CriticalPlusVal *int32       `json:"critical_plus_val"`
	IsPiercing      bool         `json:"is_piercing"`
	BreakDmgLimit   *string      `json:"break_dmg_lmt"`
	ElementID       *int32       `json:"element_id,omitempty"`
	Element         *string      `json:"element"`
}

func (d Damage) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", d),
		d.AttackType,
		d.DamageType,
		h.ObjPtrToID(d.HP),
		h.ObjPtrToID(d.MP),
		h.ObjPtrToID(d.CTB),
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
	return fmt.Sprintf("damage with attack type: %s, damage type %s, critical: %v, crit plus: %v, piercing: %t, bdl: %v, element: %v", d.AttackType, d.DamageType, h.PtrToString(d.Critical), h.PtrToString(d.CriticalPlusVal), d.IsPiercing, h.PtrToString(d.BreakDmgLimit), h.PtrToString(d.Element))
}

type ClassDamage struct {
	ID             int32   `json:"class_damage_id,omitempty"`
	Condition      *string `json:"condition"`
	DamageFormula  string  `json:"damage_formula"`
	DamageConstant int32   `json:"damage_constant"`
}

func (cd ClassDamage) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", cd),
		h.DerefOrNil(cd.Condition),
		cd.DamageFormula,
		cd.DamageConstant,
	}
}

func (cd ClassDamage) GetID() int32 {
	return cd.ID
}

func (cd *ClassDamage) SetID(id int32) {
	cd.ID = id
}

func (cd ClassDamage) Error() string {
	return fmt.Sprintf("class damage with formula: %s, damage constant %d, condition: %v", cd.DamageFormula, cd.DamageConstant, h.PtrToString(cd.Condition))
}
