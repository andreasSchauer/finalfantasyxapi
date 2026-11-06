package api

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func convertMonsterEquipment(equipment seedtypes.MonsterEquipment) apitypes.MonsterEquipment {
	monEquipment := apitypes.MonsterEquipment{
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

func convertMonsterEquipmentSlots(seedSlots seedtypes.MonsterEquipmentSlots) apitypes.MonsterEquipmentSlots {
	equipmentSlots := apitypes.MonsterEquipmentSlots{
		MinAmount: seedSlots.MinAmount,
		MaxAmount: seedSlots.MaxAmount,
		Chances:   ah.ConvertObjSlice(seedSlots.Chances, convertMonsterEquipmentSlotsChance),
	}

	return equipmentSlots
}

func convertMonsterEquipmentSlotsChance(chance seedtypes.EquipmentSlotsChance) apitypes.EquipmentSlotsChance {
	return apitypes.EquipmentSlotsChance{
		Amount: chance.Amount,
		Chance: chance.Chance,
	}
}

func convertEquipmentDrop(drop seedtypes.EquipmentDrop) apitypes.EquipmentDrop {
	return apitypes.EquipmentDrop{
		AutoAbility: nameToNamedAPIResource(E.AutoAbilities, drop.Ability, nil),
		ForcedChars: namesToNamedAPIResources(E.Characters, drop.Characters),
		IsForced:    drop.IsForced,
		Probability: drop.Probability,
	}
}
