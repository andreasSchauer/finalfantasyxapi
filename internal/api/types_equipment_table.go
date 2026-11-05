package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

type EquipmentTable struct {
	ID                      int32                          `json:"id"`
	Type                    string                         `json:"type"`
	Classification          string                         `json:"classification"`
	Priority                *int32                         `json:"priority"`
	CelestialWeapon         *apiresource.NamedAPIResource  `json:"celestial_weapon,omitempty"`
	SpecificCharacter       *apiresource.NamedAPIResource  `json:"specific_character,omitempty"`
	RequiredAutoAbilities   []apiresource.NamedAPIResource `json:"required_auto_abilities"`
	SelectableAutoAbilities []AbilityPool                  `json:"selectable_auto_abilities"`
	RequiredSlots           *int32                         `json:"required_slots"`
	Equipment               []apiresource.NamedAPIResource `json:"equipment_names"`
}

type AbilityPool struct {
	AutoAbilities []apiresource.NamedAPIResource `json:"auto_abilities"`
	ReqAmount     int32                          `json:"required_amount"`
}

func convertAbilityPool(p seedtypes.AbilityPool) AbilityPool {
	return AbilityPool{
		AutoAbilities: namesToNamedAPIResources(E.AutoAbilities, p.AutoAbilities),
		ReqAmount:     p.ReqAmount,
	}
}

func convertEquipmentName(n seedtypes.EquipmentName) apiresource.NamedAPIResource {
	return nameToNamedAPIResource(E.Equipment, n.Name, nil)
}
