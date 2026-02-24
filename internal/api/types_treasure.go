package api


type Treasure struct {
	ID              int32
	Area            AreaAPIResource  `json:"area"`
	IsPostAirship   bool             `json:"is_post_airship"`
	IsAnimaTreasure bool             `json:"is_anima_treasure"`
	Notes           *string          `json:"notes,omitempty"`
	TreasureType    string 			 `json:"treasure_type"`
	LootType        NamedAPIResource `json:"loot_type"`
	GilAmount       *int32           `json:"gil_amount,omitempty"`
	Items           []ItemAmount     `json:"items,omitempty"`
	Equipment       *FoundEquipment  `json:"equipment,omitempty"`
}