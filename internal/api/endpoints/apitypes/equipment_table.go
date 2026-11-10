package apitypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
)

type EquipmentTable struct {
	ID                      int32               `json:"id"`
	Type                    string              `json:"type"`
	Classification          string              `json:"classification"`
	Priority                *int32              `json:"priority"`
	CelestialWeapon         *apiresource.Named  `json:"celestial_weapon,omitempty"`
	SpecificCharacter       *apiresource.Named  `json:"specific_character,omitempty"`
	RequiredAutoAbilities   []apiresource.Named `json:"required_auto_abilities"`
	SelectableAutoAbilities []AbilityPool       `json:"selectable_auto_abilities"`
	RequiredSlots           *int32              `json:"required_slots"`
	Equipment               []apiresource.Named `json:"equipment_names"`
}

type AbilityPool struct {
	AutoAbilities []apiresource.Named `json:"auto_abilities"`
	ReqAmount     int32               `json:"required_amount"`
}
