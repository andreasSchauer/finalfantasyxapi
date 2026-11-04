package api

import (
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

func convertMonsterEquipment(server *Server, equipment seedtypes.MonsterEquipment) MonsterEquipment {
	monEquipment := MonsterEquipment{
		DropChance:        equipment.DropChance,
		Power:             equipment.Power,
		CriticalPlus:      equipment.CriticalPlus,
		AbilitySlots:      convertMonsterEquipmentSlots(server, equipment.AbilitySlots),
		AttachedAbilities: convertMonsterEquipmentSlots(server, equipment.AttachedAbilities),
		WeaponAbilities:   convertObjSlice(server, equipment.WeaponAbilities, convertEquipmentDrop),
		ArmorAbilities:    convertObjSlice(server, equipment.ArmorAbilities, convertEquipmentDrop),
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

func convertMonsterEquipmentSlots(server *Server, seedSlots seedtypes.MonsterEquipmentSlots) MonsterEquipmentSlots {
	equipmentSlots := MonsterEquipmentSlots{
		MinAmount: seedSlots.MinAmount,
		MaxAmount: seedSlots.MaxAmount,
		Chances:   convertObjSlice(server, seedSlots.Chances, convertMonsterEquipmentSlotsChance),
	}

	return equipmentSlots
}

type EquipmentSlotsChance struct {
	Amount int32 `json:"amount"`
	Chance int32 `json:"chance"`
}

func convertMonsterEquipmentSlotsChance(_ *Server, chance seedtypes.EquipmentSlotsChance) EquipmentSlotsChance {
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

func convertEquipmentDrop(server *Server, drop seedtypes.EquipmentDrop) EquipmentDrop {
	return EquipmentDrop{
		AutoAbility: nameToNamedAPIResource(server, E.AutoAbilities, drop.Ability, nil),
		ForcedChars: namesToNamedAPIResources(server, E.Characters, drop.Characters),
		IsForced:    drop.IsForced,
		Probability: drop.Probability,
	}
}
