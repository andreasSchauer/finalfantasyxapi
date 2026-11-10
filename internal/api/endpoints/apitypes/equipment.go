package apitypes

import "github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"

type EquipmentName struct {
	ID                      int32                 `json:"id"`
	Name                    string                `json:"name"`
	Character               apiresource.Named     `json:"character"`
	EquipmentTable          apiresource.Unnamed   `json:"equipment_table"`
	Type                    string                `json:"type"`
	Classification          string                `json:"classification"`
	Priority                *int32                `json:"priority"`
	CelestialWeapon         *apiresource.Named    `json:"celestial_weapon,omitempty"`
	RequiredAutoAbilities   []apiresource.Named   `json:"required_auto_abilities"`
	SelectableAutoAbilities []AbilityPool         `json:"selectable_auto_abilities"`
	RequiredSlots           *int32                `json:"required_slots"`
	Treasures               []apiresource.Unnamed `json:"treasures"`
	Shops                   []apiresource.Unnamed `json:"shops"`
}
