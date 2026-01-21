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
	AutoAbility NamedAPIResource   `json:"auto_ability"`
	IsForced    bool               `json:"is_forced"`
	ForcedChars []NamedAPIResource `json:"forced_characters"`
	Probability *int32             `json:"probability,omitempty"`
}

func (ed EquipmentDrop) GetAPIResource() APIResource {
	return ed.AutoAbility.GetAPIResource()
}

func getMonsterEquipment(cfg *Config, equipment *seeding.MonsterEquipment) *MonsterEquipment {
	if equipment == nil {
		return nil
	}

	monEquipment := MonsterEquipment{
		DropChance:        equipment.DropChance,
		Power:             equipment.Power,
		CriticalPlus:      equipment.CriticalPlus,
		AbilitySlots:      getMonsterEquipmentSlots(equipment.AbilitySlots),
		AttachedAbilities: getMonsterEquipmentSlots(equipment.AttachedAbilities),
		WeaponAbilities:   getEquipmentDrops(cfg, equipment.WeaponAbilities),
		ArmorAbilities:    getEquipmentDrops(cfg, equipment.ArmorAbilities),
	}

	return &monEquipment
}

func getMonsterEquipmentSlots(seedSlots seeding.MonsterEquipmentSlots) MonsterEquipmentSlots {
	equipmentSlots := MonsterEquipmentSlots{
		MinAmount: seedSlots.MinAmount,
		MaxAmount: seedSlots.MaxAmount,
		Chances:   getMonsterEquipmentSlotsChances(seedSlots.Chances),
	}

	return equipmentSlots
}

func getMonsterEquipmentSlotsChances(seedChances []seeding.EquipmentSlotsChance) []EquipmentSlotsChance {
	chances := []EquipmentSlotsChance{}

	for _, seedChance := range seedChances {
		chance := EquipmentSlotsChance{
			Amount: seedChance.Amount,
			Chance: seedChance.Chance,
		}
		chances = append(chances, chance)
	}

	return chances
}

func getEquipmentDrops(cfg *Config, seedDrops []seeding.EquipmentDrop) []EquipmentDrop {
	drops := []EquipmentDrop{}

	for _, seedDrop := range seedDrops {
		drop := EquipmentDrop{
			AutoAbility: nameToNamedAPIResource(cfg, cfg.e.autoAbilities, seedDrop.Ability, nil),
			ForcedChars: namesToNamedAPIResources(cfg, cfg.e.characters, seedDrop.Characters),
			IsForced:    seedDrop.IsForced,
			Probability: seedDrop.Probability,
		}

		drops = append(drops, drop)
	}

	return drops
}
