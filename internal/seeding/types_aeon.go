package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Aeon struct {
	ID int32
	PlayerUnit
	UnlockCondition     string       `json:"unlock_condition"`
	LocationArea        LocationArea `json:"location_area"`
	AreaID              *int32
	Category            *string         `json:"category"`
	IsOptional          bool            `json:"is_optional"`
	BattlesToRegenerate int32           `json:"num_battles_to_regenerate"`
	Weapon              []AeonEquipment `json:"weapon"`
	Armor               []AeonEquipment `json:"armor"`
	PhysAtkDmgConstant  *int32          `json:"phys_atk_damage_constant"`
	PhysAtkRange        *int32          `json:"phys_atk_range"`
	PhysAtkShatterRate  *int32          `json:"phys_atk_shatter_rate"`
	PhysAtkAccuracy     *Accuracy       `json:"phys_atk_accuracy"`
	BaseStats           AeonStat
}

func (a Aeon) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", a),
		a.PlayerUnit.ID,
		a.UnlockCondition,
		h.DerefOrNil(a.AreaID),
		a.IsOptional,
		a.BattlesToRegenerate,
		h.DerefOrNil(a.PhysAtkDmgConstant),
		h.DerefOrNil(a.PhysAtkRange),
		h.DerefOrNil(a.PhysAtkShatterRate),
		h.ObjPtrToID(a.PhysAtkAccuracy),
	}
}

func (a Aeon) ToKeyFields() []any {
	return []any{
		a.Name,
	}
}

func (a Aeon) GetID() int32 {
	return a.ID
}

func (a Aeon) Error() string {
	return fmt.Sprintf("aeon %s", a.Name)
}

func (a Aeon) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   a.ID,
		Name: a.Name,
	}
}
