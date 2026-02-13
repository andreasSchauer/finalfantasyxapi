package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/seeding"


type FoundEquipment struct {
	EquipmentName    NamedAPIResource   `json:"name"`
	Abilities        []NamedAPIResource `json:"abilities"`
	EmptySlotsAmount int32              `json:"empty_slots_amount"`
}

func convertFoundEquipment(cfg *Config, fe seeding.FoundEquipment) FoundEquipment {
	return FoundEquipment{
		EquipmentName:    nameToNamedAPIResource(cfg, cfg.e.equipment, fe.Name, nil),
		Abilities:        namesToNamedAPIResources(cfg, cfg.e.autoAbilities, fe.Abilities),
		EmptySlotsAmount: fe.EmptySlotsAmount,
	}
}