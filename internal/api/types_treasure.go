package api

type Treasure struct {
	ID              int32                              `json:"id"`
	Area            AreaAPIResource                    `json:"area"`
	IsPostAirship   bool                               `json:"is_post_airship"`
	IsStoryBased    bool                               `json:"is_story_based"`
	IsAnimaTreasure bool                               `json:"is_anima_treasure"`
	Notes           *string                            `json:"notes,omitempty"`
	TreasureType    string                             `json:"treasure_type"`
	LootType        NamedAPIResource                   `json:"loot_type"`
	GilAmount       *int32                             `json:"gil_amount,omitempty"`
	Items           []ResourceAmount[TypedAPIResource] `json:"items,omitempty"`
	Equipment       *FoundEquipment                    `json:"equipment,omitempty"`
}
