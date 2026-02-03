package main

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
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

func convertMonsterEquipment(cfg *Config, equipment seeding.MonsterEquipment) MonsterEquipment {
	monEquipment := MonsterEquipment{
		DropChance:        equipment.DropChance,
		Power:             equipment.Power,
		CriticalPlus:      equipment.CriticalPlus,
		AbilitySlots:      convertMonsterEquipmentSlots(cfg, equipment.AbilitySlots),
		AttachedAbilities: convertMonsterEquipmentSlots(cfg, equipment.AttachedAbilities),
		WeaponAbilities:   convertObjSlice(cfg, equipment.WeaponAbilities, convertEquipmentDrop),
		ArmorAbilities:    convertObjSlice(cfg, equipment.ArmorAbilities, convertEquipmentDrop),
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

func convertMonsterEquipmentSlots(cfg *Config, seedSlots seeding.MonsterEquipmentSlots) MonsterEquipmentSlots {
	equipmentSlots := MonsterEquipmentSlots{
		MinAmount: seedSlots.MinAmount,
		MaxAmount: seedSlots.MaxAmount,
		Chances:   convertObjSlice(cfg, seedSlots.Chances, convertMonsterEquipmentSlotsChance),
	}

	return equipmentSlots
}

type EquipmentSlotsChance struct {
	Amount int32 `json:"amount"`
	Chance int32 `json:"chance"`
}

func convertMonsterEquipmentSlotsChance(_ *Config, chance seeding.EquipmentSlotsChance) EquipmentSlotsChance {
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

func convertEquipmentDrop(cfg *Config, drop seeding.EquipmentDrop) EquipmentDrop {
	return EquipmentDrop{
		AutoAbility: nameToNamedAPIResource(cfg, cfg.e.autoAbilities, drop.Ability, nil),
		ForcedChars: namesToNamedAPIResources(cfg, cfg.e.characters, drop.Characters),
		IsForced:    drop.IsForced,
		Probability: drop.Probability,
	}
}
