package sectiontypes

type ShopSimple struct {
	ID          int32          `json:"id"`
	URL         string         `json:"url"`
	Area        string         `json:"area"`
	Category    string         `json:"category"`
	Notes       *string        `json:"notes,omitempty"`
	PreAirship  *SubShopSimple `json:"pre_airship"`
	PostAirship *SubShopSimple `json:"post_airship"`
}

func (s ShopSimple) GetURL() string {
	return s.URL
}

type SubShopSimple struct {
	Items     []string              `json:"items"`
	Equipment []ShopEquipmentSimple `json:"equipment"`
}

type ShopEquipmentSimple struct {
	EquipmentName string  `json:"equipment_name"`
	Abilities     *string `json:"abilities"`
}
