package api

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api_helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
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

func convertMonsterEquipment(equipment seedtypes.MonsterEquipment) MonsterEquipment {
	monEquipment := MonsterEquipment{
		DropChance:        equipment.DropChance,
		Power:             equipment.Power,
		CriticalPlus:      equipment.CriticalPlus,
		AbilitySlots:      convertMonsterEquipmentSlots(equipment.AbilitySlots),
		AttachedAbilities: convertMonsterEquipmentSlots(equipment.AttachedAbilities),
		WeaponAbilities:   ah.ConvertObjSlice(equipment.WeaponAbilities, convertEquipmentDrop),
		ArmorAbilities:    ah.ConvertObjSlice(equipment.ArmorAbilities, convertEquipmentDrop),
	}

	return monEquipment
}

func (me MonsterEquipment) IsZero() bool {
	return me.DropChance == 0
}

type MonsterEquipmentSlots struct {
	MinAmount int32                  `json:"min_amount"`
	MaxAmount int32                  `json:"max_amount"`
	Chances   []EquipmentSlotsChance `json:"chances"`
}

func convertMonsterEquipmentSlots(seedSlots seedtypes.MonsterEquipmentSlots) MonsterEquipmentSlots {
	equipmentSlots := MonsterEquipmentSlots{
		MinAmount: seedSlots.MinAmount,
		MaxAmount: seedSlots.MaxAmount,
		Chances:   ah.ConvertObjSlice(seedSlots.Chances, convertMonsterEquipmentSlotsChance),
	}

	return equipmentSlots
}

type EquipmentSlotsChance struct {
	Amount int32 `json:"amount"`
	Chance int32 `json:"chance"`
}

func convertMonsterEquipmentSlotsChance(chance seedtypes.EquipmentSlotsChance) EquipmentSlotsChance {
	return EquipmentSlotsChance{
		Amount: chance.Amount,
		Chance: chance.Chance,
	}
}

type EquipmentDrop struct {
	AutoAbility NamedAPIResource   `json:"auto_ability"`
	IsForced    bool               `json:"is_forced"`
	ForcedChars []NamedAPIResource `json:"forced_characters"`
	Probability *int32             `json:"probability,omitempty"`
}

func (ed EquipmentDrop) GetAPIResource() APIResource {
	return ed.AutoAbility.GetAPIResource()
}

func convertEquipmentDrop(drop seedtypes.EquipmentDrop) EquipmentDrop {
	return EquipmentDrop{
		AutoAbility: nameToNamedAPIResource(E.AutoAbilities, drop.Ability, nil),
		ForcedChars: namesToNamedAPIResources(E.Characters, drop.Characters),
		IsForced:    drop.IsForced,
		Probability: drop.Probability,
	}
}
