package apitypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
)

type FoundEquipment struct {
	EquipmentName    apiresource.NamedAPIResource   `json:"name"`
	Abilities        []apiresource.NamedAPIResource `json:"abilities"`
	EmptySlotsAmount int32                          `json:"empty_slots_amount"`
}
