package apitypes

import "github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"

type Treasure struct {
	ID              int32                                           `json:"id"`
	Area            apiresource.Area                                `json:"area"`
	Availability    string                                          `json:"availability"`
	IsAnimaTreasure bool                                            `json:"is_anima_treasure"`
	Notes           *string                                         `json:"notes,omitempty"`
	TreasureType    string                                          `json:"treasure_type"`
	LootType        string                                          `json:"loot_type"`
	GilAmount       *int32                                          `json:"gil_amount,omitempty"`
	Items           []apiresource.ResourceAmount[apiresource.Typed] `json:"items,omitempty"`
	Equipment       *FoundEquipment                                 `json:"equipment,omitempty"`
}
