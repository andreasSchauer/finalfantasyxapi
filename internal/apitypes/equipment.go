package apitypes

import "github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"

type EquipmentName struct {
	ID                      int32                            `json:"id"`
	Name                    string                           `json:"name"`
	Character               apiresource.NamedAPIResource     `json:"character"`
	EquipmentTable          apiresource.UnnamedAPIResource   `json:"equipment_table"`
	Type                    string                           `json:"type"`
	Classification          string                           `json:"classification"`
	Priority                *int32                           `json:"priority"`
	CelestialWeapon         *apiresource.NamedAPIResource    `json:"celestial_weapon,omitempty"`
	RequiredAutoAbilities   []apiresource.NamedAPIResource   `json:"required_auto_abilities"`
	SelectableAutoAbilities []AbilityPool                    `json:"selectable_auto_abilities"`
	RequiredSlots           *int32                           `json:"required_slots"`
	Treasures               []apiresource.UnnamedAPIResource `json:"treasures"`
	Shops                   []apiresource.UnnamedAPIResource `json:"shops"`
}
