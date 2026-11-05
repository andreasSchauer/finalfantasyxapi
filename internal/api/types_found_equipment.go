package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"

type FoundEquipment struct {
	EquipmentName    NamedAPIResource   `json:"name"`
	Abilities        []NamedAPIResource `json:"abilities"`
	EmptySlotsAmount int32              `json:"empty_slots_amount"`
}

func convertFoundEquipment(fe seedtypes.TreasureEquipment) FoundEquipment {
	return FoundEquipment{
		EquipmentName:    nameToNamedAPIResource(E.Equipment, fe.Name, nil),
		Abilities:        namesToNamedAPIResources(E.AutoAbilities, fe.Abilities),
		EmptySlotsAmount: fe.EmptySlotsAmount,
	}
}
