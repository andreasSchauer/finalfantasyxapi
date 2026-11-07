package sectiontypes

type TreasureSimple struct {
	ID           int32            `json:"id"`
	URL          string           `json:"url"`
	Area         string           `json:"area"`
	Availability string           `json:"availability"`
	Notes        *string          `json:"notes,omitempty"`
	TreasureType string           `json:"treasure_type"`
	LootType     string           `json:"loot_type"`
	GilAmount    *int32           `json:"gil_amount,omitempty"`
	Items        []string         `json:"items,omitempty"`
	Equipment    *EquipmentSimple `json:"equipment,omitempty"`
}

func (t TreasureSimple) GetURL() string {
	return t.URL
}
