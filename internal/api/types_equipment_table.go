package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/seeding"

type EquipmentTable struct {
	ID                      int32              `json:"id"`
	Type                    string             `json:"type"`
	Classification          string             `json:"classification"`
	Priority                *int32             `json:"priority"`
	CelestialWeapon         *NamedAPIResource  `json:"celestial_weapon,omitempty"`
	SpecificCharacter       *NamedAPIResource  `json:"specific_character,omitempty"`
	RequiredAutoAbilities   []NamedAPIResource `json:"required_auto_abilities"`
	SelectableAutoAbilities []AbilityPool      `json:"selectable_auto_abilities"`
	RequiredSlots           *int32             `json:"required_slots"`
	Equipment               []NamedAPIResource `json:"equipment_names"`
}

type AbilityPool struct {
	AutoAbilities []NamedAPIResource `json:"auto_abilities"`
	ReqAmount     int32              `json:"required_amount"`
}

func convertAbilityPool(cfg *Config, p seeding.AbilityPool) AbilityPool {
	return AbilityPool{
		AutoAbilities: namesToNamedAPIResources(cfg, cfg.e.autoAbilities, p.AutoAbilities),
		ReqAmount:     p.ReqAmount,
	}
}

func convertEquipmentName(cfg *Config, n seeding.EquipmentName) NamedAPIResource {
	return nameToNamedAPIResource(cfg, cfg.e.equipment, n.Name, nil)
}
