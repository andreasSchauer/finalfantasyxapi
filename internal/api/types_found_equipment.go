package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"

type FoundEquipment struct {
	EquipmentName    NamedAPIResource   `json:"name"`
	Abilities        []NamedAPIResource `json:"abilities"`
	EmptySlotsAmount int32              `json:"empty_slots_amount"`
}

func convertFoundEquipment(cfg *Config, fe seedtypes.TreasureEquipment) FoundEquipment {
	return FoundEquipment{
		EquipmentName:    nameToNamedAPIResource(cfg, cfg.E.Equipment, fe.Name, nil),
		Abilities:        namesToNamedAPIResources(cfg, cfg.E.AutoAbilities, fe.Abilities),
		EmptySlotsAmount: fe.EmptySlotsAmount,
	}
}
