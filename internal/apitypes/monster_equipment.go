package apitypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
)

type MonsterEquipment struct {
	DropChance        int32                 `json:"drop_chance"`
	Power             int32                 `json:"power"`
	CriticalPlus      int32                 `json:"critical_plus"`
	AbilitySlots      MonsterEquipmentSlots `json:"ability_slots"`
	AttachedAbilities MonsterEquipmentSlots `json:"attached_abilities"`
	WeaponAbilities   []EquipmentDrop       `json:"weapon_abilities"`
	ArmorAbilities    []EquipmentDrop       `json:"armor_abilities"`
}

func (me MonsterEquipment) IsZero() bool {
	return me.DropChance == 0
}

type MonsterEquipmentSlots struct {
	MinAmount int32                  `json:"min_amount"`
	MaxAmount int32                  `json:"max_amount"`
	Chances   []EquipmentSlotsChance `json:"chances"`
}

type EquipmentSlotsChance struct {
	Amount int32 `json:"amount"`
	Chance int32 `json:"chance"`
}

type EquipmentDrop struct {
	AutoAbility apiresource.NamedAPIResource   `json:"auto_ability"`
	IsForced    bool                           `json:"is_forced"`
	ForcedChars []apiresource.NamedAPIResource `json:"forced_characters"`
	Probability *int32                         `json:"probability,omitempty"`
}

func (ed EquipmentDrop) GetAPIResource() apiresource.APIResource {
	return ed.AutoAbility.GetAPIResource()
}