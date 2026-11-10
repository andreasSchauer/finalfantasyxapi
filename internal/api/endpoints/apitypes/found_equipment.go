package apitypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
)

type FoundEquipment struct {
	EquipmentName    apiresource.Named   `json:"name"`
	Abilities        []apiresource.Named `json:"abilities"`
	EmptySlotsAmount int32               `json:"empty_slots_amount"`
}
