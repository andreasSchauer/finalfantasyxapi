package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"

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

func convertAbilityPool(cfg *Config, p seedtypes.AbilityPool) AbilityPool {
	return AbilityPool{
		AutoAbilities: namesToNamedAPIResources(cfg, E.AutoAbilities, p.AutoAbilities),
		ReqAmount:     p.ReqAmount,
	}
}

func convertEquipmentName(cfg *Config, n seedtypes.EquipmentName) NamedAPIResource {
	return nameToNamedAPIResource(cfg, E.Equipment, n.Name, nil)
}
